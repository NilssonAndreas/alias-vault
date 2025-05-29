package cmd

import (
	"aliasvault/vault"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var tags string

var addCmd = &cobra.Command{
	Use:   "add <alias> <command>",
	Short: "Add a new command alias to the vault",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		alias := args[0]
		command := args[1]
		tagList := strings.Split(tags, ",")

		err := vault.SaveAlias(alias, command, tagList)
		if err != nil {
			fmt.Println("Error while savaing:", err)
		} else {
			fmt.Println("Alias saved:", alias)
		}
	},
}

func init() {
	addCmd.Flags().StringVarP(&tags, "tags", "t", "", "Comma-separated tags for the alias")
	rootCmd.AddCommand(addCmd)
}
