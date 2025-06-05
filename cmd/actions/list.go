package actions

import (
	"sonar-api/cmd/components"

	"github.com/spf13/cobra"
)

func ListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Recupera recursos desde SonarCloud",
		Long:  "Este comando permite recuperar todos los recursos específicos desde SonarCloud, como proyectos, métricas, etc.",
	}

	cmd.AddCommand(components.ListProjectsCmd())
	cmd.AddCommand(components.ListQualityCmd())
	return cmd
}
