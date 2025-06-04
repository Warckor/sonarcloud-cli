package cmd

import (
	"log"
	"sonar-api/internal/sonar"

	"github.com/spf13/cobra"
)

func ListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Lista todos los proyectos de SonarCloud",
		Long:    `Lista todos los proyectos disponibles en SonarCloud para una organización específica.`,
		Example: `sonarcli list --org my-org`,
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

			client := sonar.NewClient()
			projects, err := sonar.ListProjects(client, finalOrg)
			if err != nil {
				log.Fatal("Error al listar proyectos:", err)
				return
			}

			if len(projects) == 0 {
				log.Println("No se encontraron proyectos.")
				return
			}

			log.Println("Proyectos encontrados:")
			for _, project := range projects {
				log.Printf(" - %s (%s)\n", project.Name, project.Key)
			}
		},
	}

	cmd.Flags().StringVarP(&org, "org", "O", "", "Organización de SonarCloud (opcional, usa la configuración por defecto si no se especifica)")
	return cmd
}
