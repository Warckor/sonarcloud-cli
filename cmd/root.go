package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sonarcli",
	Short: "CLI para trabajar con SonarCloud",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(ListCmd())
	rootCmd.AddCommand(GetCmd())
}
