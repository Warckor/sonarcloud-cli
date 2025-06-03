package cmd

import (
	"log"
	"sonar-api/internal/sonar"

	"github.com/spf13/cobra"
)

var (
	org        string
	projectKey string
)

var listCmd = &cobra.Command{
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

var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "Obtener el nombre de un proyecto específico",
	Long:    `Obtiene los detalles de un proyecto específico en SonarCloud utilizando su clave.`,
	Example: `sonarcli get --org my-org --project-key my-project-key`,
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

func init() {
	listCmd.Flags().StringVarP(&org, "org", "O", "", "Organización de SonarCloud (opcional, usa la configuración por defecto si no se especifica)")
	getCmd.Flags().StringVarP(&org, "org", "O", "", "Organización de SonarCloud (opcional, usa la configuración por defecto si no se especifica)")
	getCmd.Flags().StringVarP(&projectKey, "project-key", "P", "", "Clave del proyecto en SonarCloud (obligatorio)")
	getCmd.MarkFlagRequired("project-key")

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(getCmd)
}
