package cmd

import (
	"aliasvault/vault"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Search for command aliases in the vault",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.ToLower(args[0])
		results, err := vault.SearchAliases(query)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		if len(results) == 0 {
			fmt.Println("No matches found.")
			return
		}
		for _, a := range results {
			fmt.Printf("📎 %s\n  → %s\n  🏷️ %v\n\n", a.Alias, a.Command, a.Tags)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
