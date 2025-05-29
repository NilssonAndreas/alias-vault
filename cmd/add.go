package cmd

import (
	"aliasvault/vault"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var tags string

var addCmd = &cobra.Command{
	Use:   "add <alias> <command>",
	Short: "Add a new command alias to the vault",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var alias, command, tagStr string

		if len(args) < 2 {
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Alias: ")
			aliasInput, _ := reader.ReadString('\n')
			alias = strings.TrimSpace(aliasInput)

			fmt.Print("Command: ")
			cmdInput, _ := reader.ReadString('\n')
			command = strings.TrimSpace(cmdInput)

			fmt.Print("Tags (comma-separated): ")
			tagInput, _ := reader.ReadString('\n')
			tagStr = strings.TrimSpace(tagInput)
		} else {
			alias = args[0]
			command = args[1]
			tagStr = tags
		}

		tagList := strings.Split(tagStr, ",")
		for i := range tagList {
			tagList[i] = strings.TrimSpace(tagList[i])
		}

		err := vault.SaveAlias(alias, command, tagList)
		if err != nil {
			fmt.Println("Error while saving:", err)
		} else {
			fmt.Println("Alias saved:", alias)
		}
	},
}

func init() {
	addCmd.Flags().StringVarP(&tags, "tags", "t", "", "Comma-separated tags for the alias")
	rootCmd.AddCommand(addCmd)
}
