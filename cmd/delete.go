package cmd

import (
	"aliasvault/vault"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <alias>",
	Short: "Remove a saved command alias from the vault",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		alias := args[0]
		err := vault.DeleteAlias(alias)
		if err != nil {
			fmt.Println("Error", err)
			os.Exit(1)
		}
		fmt.Println("Removed", alias)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
