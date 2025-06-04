package components

import (
	"log"
	"sonar-api/internal/sonar"

	"github.com/spf13/cobra"
)

var (
	org        string
	projectKey string
)

func GetProjectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "project",
		Short:   "Obtener el nombre de un proyecto específico",
		Long:    `Obtiene los detalles de un proyecto específico en SonarCloud utilizando su clave.`,
		Example: `sonarcli get project --org my-org --project-key my-project-key`,
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := sonar.LoadConfig()
			if err != nil {
				cmd.PrintErrf("Error loading config: %v\n", err)
				return
			}

			finalOrg := org
			if finalOrg == "" {
				finalOrg = cfg.Organization
			}

			if finalOrg == "" {
				cmd.PrintErrf("Error: organization is required\n")
				return
			} else if projectKey == "" {
				cmd.PrintErrf("Error: project-key is required\n")
				return
			}

			params := map[string]string{
				"organization": finalOrg,
				"projects":     projectKey,
			}

			client := sonar.NewClient()
			project, err := sonar.GetProject(client, params)
			if err != nil {
				log.Fatal("Error al obtener el proyecto:", err)
				return
			}
			if len(project.Key) == 0 {
				log.Println("No se encontró el proyecto.")
				return
			}
			log.Printf("Detalles del proyecto %s:\n", project.Name)
			log.Printf(" - Clave: %s\n", project.Key)
			log.Printf(" - Nombre: %s\n", project.Name)
		},
	}

	cmd.Flags().StringVarP(&org, "org", "O", "", "Organización de SonarCloud (opcional, usa la configuración por defecto si no se especifica)")
	cmd.Flags().StringVarP(&projectKey, "project-key", "p", "", "Clave del proyecto de SonarCloud (requerido)")
	cmd.MarkFlagRequired("project-key")

	return cmd
}
