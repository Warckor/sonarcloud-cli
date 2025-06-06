package actions

import (
	"sonar-api/cmd/components"

	"github.com/spf13/cobra"
)

func CreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Crear un nuevo recurso",
		Long:  `Crea un nuevo recurso en SonarCloud.`,
	}

	cmd.AddCommand(components.CreateProjectCmd())
	return cmd
}
