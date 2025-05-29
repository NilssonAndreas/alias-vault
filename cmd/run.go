package cmd

import (
	"aliasvault/vault"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

func getShellCommand(cmd string) *exec.Cmd {
	if runtime.GOOS == "windows" {
		return exec.Command("cmd", "/C", cmd)
	}
	return exec.Command("sh", "-c", cmd)
}

var runCmd = &cobra.Command{
	Use:   "run <alias>",
	Short: "Run a saved command alias from the vault",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		alias := args[0]
		entry, err := vault.GetAlias(alias)
		if err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(1)
		}

		fmt.Println("Run:", entry.Command)
		execCmd := getShellCommand(entry.Command)
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr
		execCmd.Stdin = os.Stdin

		if err := execCmd.Run(); err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
