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
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/k1LoW/gh-do/version"
	"github.com/k1LoW/go-github-client/v52/factory"
	"github.com/spf13/cobra"
)

var host string

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
			return nil
		}

		envs := os.Environ()
		envs = append(envs, fmt.Sprintf("export GH_HOST=%s\n", host))
		envs = append(envs, fmt.Sprintf("export GH_TOKEN=%s\n", token))
		envs = append(envs, fmt.Sprintf("export GH_ENTERPRISE_TOKEN=%s\n", etoken))
		envs = append(envs, fmt.Sprintf("export GITHUB_ENTERPRISE_TOKEN=%s\n", etoken))
		envs = append(envs, fmt.Sprintf("export GITHUB_TOKEN=%s\n", token))
		envs = append(envs, fmt.Sprintf("export GITHUB_API_URL=%s\n", v3ep))
		envs = append(envs, fmt.Sprintf("export GITHUB_GRAPHQL_URL=%s\n", v4ep))
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
	rootCmd.Flags().StringVarP(&host, "host", "", "", "host")
}
