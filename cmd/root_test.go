package cmd

import (
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"testing"
)

func TestTokenFromSecureStorage(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("fake shell script is not supported on windows")
	}
	tests := []struct {
		name     string
		host     string
		user     string
		wantArgs []string
	}{
		{
			name:     "without user",
			host:     "github.com",
			user:     "",
			wantArgs: []string{"auth", "token", "--secure-storage", "--hostname", "github.com"},
		},
		{
			name:     "with user",
			host:     "github.com",
			user:     "alice",
			wantArgs: []string{"auth", "token", "--secure-storage", "--hostname", "github.com", "--user", "alice"},
		},
		{
			name:     "with enterprise host and user",
			host:     "enterprise.internal",
			user:     "bob",
			wantArgs: []string{"auth", "token", "--secure-storage", "--hostname", "enterprise.internal", "--user", "bob"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			argsFile := filepath.Join(dir, "args")
			scriptPath := filepath.Join(dir, "gh")
			script := "#!/bin/sh\nprintf '%s\\n' \"$@\" > " + argsFile + "\necho fake-token\n"
			if err := os.WriteFile(scriptPath, []byte(script), 0o755); err != nil { //nolint:gosec // executable permission is required for the fake shell script
				t.Fatal(err)
			}
			t.Setenv("GH_PATH", scriptPath)

			got, err := tokenFromSecureStorage(tt.host, tt.user)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != "fake-token" {
				t.Errorf("token = %q, want %q", got, "fake-token")
			}

			b, err := os.ReadFile(argsFile)
			if err != nil {
				t.Fatalf("failed to read args: %v", err)
			}
			gotArgs := strings.Split(strings.TrimRight(string(b), "\n"), "\n")
			if !slices.Equal(gotArgs, tt.wantArgs) {
				t.Errorf("args = %v, want %v", gotArgs, tt.wantArgs)
			}
		})
	}
}

func TestTokenFromSecureStorageError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("fake shell script is not supported on windows")
	}
	dir := t.TempDir()
	scriptPath := filepath.Join(dir, "gh")
	script := "#!/bin/sh\necho 'boom' >&2\nexit 1\n"
	if err := os.WriteFile(scriptPath, []byte(script), 0o755); err != nil { //nolint:gosec // executable permission is required for the fake shell script
		t.Fatal(err)
	}
	t.Setenv("GH_PATH", scriptPath)

	_, err := tokenFromSecureStorage("github.com", "")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Error() != "boom" {
		t.Errorf("error = %q, want %q", err.Error(), "boom")
	}
}
