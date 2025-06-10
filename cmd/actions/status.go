package actions

import (
	"sonar-api/cmd/components"

	"github.com/spf13/cobra"
)

func StatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Verifica el estado de un recurso específico en SonarCloud",
		Long:  "Este comando permite verificar el estado de recursos específicos en SonarCloud, como proyectos, métricas, etc.",
	}

	cmd.AddCommand(components.StatusQualityCmd())
	return cmd
}
