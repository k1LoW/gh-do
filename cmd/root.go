/*
Copyright © 2023 Ken'ichiro Oyama <k1lowxb@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cli/safeexec"
	"github.com/k1LoW/gh-do/version"
	"github.com/k1LoW/go-github-client/v52/factory"
	"github.com/spf13/cobra"
)

var (
	host     string
	insecure bool
	cenvs    []string
)

var rootCmd = &cobra.Command{
	Use:          "gh-do",
	Short:        "gh-do is a tool to do anything using GitHub credentials",
	Long:         `gh-do is a tool to do anything using GitHub credentials.`,
	Version:      version.Version,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if host == "" {
			host = os.Getenv("GH_HOST")
		}
		os.Setenv("GH_HOST", host)
		token, v3ep, _, v4ep, host, _, _ := factory.GetAllDetected()
		if !insecure {
			sectoken, err := tokenFromSecureStorage(host)
			if err != nil {
				return fmt.Errorf("failed to get credentials stored in secure storage for %s: %w", host, err)
			}
			token = sectoken
		}

		if token == "" {
			return fmt.Errorf("failed to get credentials for %s", host)
		}

		// Clear environment variables for GitHub
		os.Unsetenv("GH_HOST")
		os.Unsetenv("GH_TOKEN")
		os.Unsetenv("GH_ENTERPRISE_TOKEN")
		os.Unsetenv("GITHUB_ENTERPRISE_TOKEN")
		os.Unsetenv("GITHUB_TOKEN")
		os.Unsetenv("GITHUB_API_URL")
		os.Unsetenv("GITHUB_GRAPHQL_URL")

		var etoken string
		if !strings.Contains(host, "github.com") {
			etoken = token
		}

		if len(args) == 0 {
			cmd.Printf("export GH_HOST=%s\n", host)
			cmd.Printf("export GH_TOKEN=%s\n", token)
			cmd.Printf("export GH_ENTERPRISE_TOKEN=%s\n", etoken)
			cmd.Printf("export GITHUB_ENTERPRISE_TOKEN=%s\n", etoken)
			cmd.Printf("export GITHUB_TOKEN=%s\n", token)
			cmd.Printf("export GITHUB_API_URL=%s\n", v3ep)
			cmd.Printf("export GITHUB_GRAPHQL_URL=%s\n", v4ep)
			for _, e := range cenvs {
				cmd.Printf("export %s=%s\n", e, token)
			}
			return nil
		}

		envs := os.Environ()
		envs = append(envs, fmt.Sprintf("GH_HOST=%s", host))
		envs = append(envs, fmt.Sprintf("GH_TOKEN=%s", token))
		envs = append(envs, fmt.Sprintf("GH_ENTERPRISE_TOKEN=%s", etoken))
		envs = append(envs, fmt.Sprintf("GITHUB_ENTERPRISE_TOKEN=%s", etoken))
		envs = append(envs, fmt.Sprintf("GITHUB_TOKEN=%s", token))
		envs = append(envs, fmt.Sprintf("GITHUB_API_URL=%s", v3ep))
		envs = append(envs, fmt.Sprintf("GITHUB_GRAPHQL_URL=%s", v4ep))
		for _, e := range cenvs {
			envs = append(envs, fmt.Sprintf("%s=%s", e, token))
		}
		command := args[0]
		c := exec.Command(command, args[1:]...)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Env = envs
		cmd.SilenceErrors = true
		if err := c.Run(); err != nil {
			os.Exit(c.ProcessState.ExitCode())
		}

		return nil
	},
}

func Execute() {
	rootCmd.SetOut(os.Stdout)
	rootCmd.SetErr(os.Stderr)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("help", "", false, "") // disable -h for help
	rootCmd.Flags().StringVarP(&host, "hostname", "h", "", "The hostname of the GitHub instance to do")
	rootCmd.Flags().BoolVarP(&insecure, "insecure", "", false, "Use insecure credentials")
	rootCmd.Flags().StringSliceVarP(&cenvs, "credential-env-key", "e", []string{}, "Set credential to specified env key")
}

func tokenFromSecureStorage(host string) (string, error) {
	var err error
	gh := os.Getenv("GH_PATH")
	if gh == "" {
		gh, err = safeexec.LookPath("gh")
		if err != nil {
			return "", err
		}
	}
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	cmd := exec.Command(gh, "auth", "token", "--secure-storage", "--hostname", host)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	if err := cmd.Run(); err != nil {
		return "", errors.New(strings.TrimSpace(stderr.String()))
	}
	return strings.TrimSpace(stdout.String()), nil
}
