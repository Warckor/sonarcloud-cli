package cmd

import (
	"os"

	"sonar-api/cmd/actions"

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
	rootCmd.AddCommand(actions.ListCmd())
	rootCmd.AddCommand(actions.GetCmd())
}
