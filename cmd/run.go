package cmd

import (
	"aliasvault/vault"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

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
		parts := []string{"cmd", "/C", entry.Command}
		if execCmd := exec.Command(parts[0], parts[1:]...); execCmd != nil {
			execCmd.Stdout = os.Stdout
			execCmd.Stderr = os.Stderr
			execCmd.Stdin = os.Stdin
			err := execCmd.Run()
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
