package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vault",
	Short: "AliasVault - Save, tag, and reuse commands",
	Long:  `AliasVault is a CLI tool for saving, tagging, and reusing commands.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
