package actions

import (
	"sonar-api/cmd/components"

	"github.com/spf13/cobra"
)

func GetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Recupera recursos desde SonarCloud",
		Long:  "Este comando permite recuperar recursos específicos desde SonarCloud, como proyectos, métricas, etc.",
	}

	cmd.AddCommand(components.GetProjectCmd())
	cmd.AddCommand(components.GetQualityCmd())
	return cmd
}
