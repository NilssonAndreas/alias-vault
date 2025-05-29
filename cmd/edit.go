package cmd

import (
	"aliasvault/vault"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit <alias>",
	Short: "Edit an existing alias in the vault",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		aliasName := args[0]

		// Get existing alias
		alias, err := vault.GetAliasByName(aliasName)
		if err != nil {
			fmt.Println("Alias not found:", aliasName)
			return
		}

		reader := bufio.NewReader(os.Stdin)

		// Prompt for new command
		fmt.Printf("Current Command: %s\nNew Command (press Enter to keep): ", alias.Command)
		newCmd, _ := reader.ReadString('\n')
		newCmd = strings.TrimSpace(newCmd)
		if newCmd != "" {
			alias.Command = newCmd
		}

		// Prompt for new tags
		fmt.Printf("Current Tags: %s\nNew Tags (comma-separated, Enter to keep): ", strings.Join(alias.Tags, ", "))
		newTagsRaw, _ := reader.ReadString('\n')
		newTagsRaw = strings.TrimSpace(newTagsRaw)
		if newTagsRaw != "" {
			newTags := strings.Split(newTagsRaw, ",")
			for i := range newTags {
				newTags[i] = strings.TrimSpace(newTags[i])
			}
			alias.Tags = newTags
		}

		// Save updated alias
		err = vault.SaveAlias(alias.Alias, alias.Command, alias.Tags)
		if err != nil {
			fmt.Println("Failed to update alias:", err)
		} else {
			fmt.Println("Alias updated:", alias.Alias)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
