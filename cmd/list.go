package cmd

import (
	"aliasvault/vault"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved command aliases",
	Run: func(cmd *cobra.Command, args []string) {
		aliases, err := vault.GetAllAliases()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		if len(aliases) == 0 {
			fmt.Println("No aliases found.")
			return
		}

		writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(writer, "ALIAS\tCOMMAND\tTAGS\t")
		for _, a := range aliases {
			fmt.Fprintf(writer, "%s\t%s\t%v\n", a.Alias, a.Command, a.Tags)
		}
		writer.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
