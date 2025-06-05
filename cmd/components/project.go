package components

import (
	"sonar-api/internal/sonar"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var (
	org        string
	projectKey string
	name       string
)

func GetProjectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "project",
		Short:   "Obtener el nombre de un proyecto específico",
		Long:    `Obtiene los detalles de un proyecto específico en SonarCloud utilizando su clave.`,
		Example: `sonarcli get project --org my-org --project-key my-project-key`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if projectKey == "" && name == "" {
				pterm.Error.Printf("Error: Al menos uno de los siguientes parámetros es requerido: --project-key, --name\n")
				return cmd.Help()
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := sonar.LoadConfig()
			if err != nil {
				pterm.Error.Printf("Error loading config: %v\n", err)
				return
			}

			finalOrg := org
			if finalOrg == "" {
				finalOrg = cfg.Organization
			}

			if finalOrg == "" {
				pterm.Error.Printf("Error: organization is required\n")
				return
			}

			params := map[string]string{
				"organization": finalOrg,
			}

			if projectKey != "" && name != "" {
				params["projects"] = projectKey
				params["name"] = name
			} else if projectKey != "" && name == "" {
				params["projects"] = projectKey
			} else if projectKey == "" && name != "" {
				params["q"] = name
			}

			client := sonar.NewClient()
			project, err := sonar.GetProject(client, params)
			if err != nil {
				pterm.Fatal.Printf("Error al obtener el proyecto: %v\n", err)
				return
			}

			if len(project.Key) == 0 {
				pterm.Warning.Println("No se encontró el proyecto.")
				return
			}

			pterm.DefaultSection.Println("Detalles del Proyecto")
			pterm.DefaultTable.WithHasHeader(true).WithData(pterm.TableData{
				{"Clave", "Nombre"},
				{project.Key, project.Name},
			}).Render()
		},
	}

	cmd.Flags().StringVarP(&org, "org", "o", "", "Organización de SonarCloud (opcional, usa la configuración por defecto si no se especifica)")
	cmd.Flags().StringVarP(&projectKey, "projectKey", "p", "", "Clave del proyecto de SonarCloud")
	cmd.Flags().StringVarP(&name, "name", "n", "", "Nombre del proyecto de SonarCloud")

	return cmd
}

func ListProjectsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "project",
		Short:   "Lista todos los proyectos de SonarCloud",
		Long:    `Lista todos los proyectos disponibles en SonarCloud para una organización específica.`,
		Example: `sonarcli list --org my-org`,
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := sonar.LoadConfig()
			if err != nil {
				pterm.Error.Printf("Error loading config: %v\n", err)
				return
			}

			finalOrg := org
			if finalOrg == "" {
				finalOrg = cfg.Organization
			}

			client := sonar.NewClient()
			projects, err := sonar.ListProjects(client, finalOrg)
			if err != nil {
				pterm.Fatal.Printf("Error al listar proyectos: %v\n", err)
				return
			}

			if len(projects) == 0 {
				pterm.Warning.Println("No se encontraron proyectos.")
				return
			}

			table := pterm.TableData{
				{"Clave", "Nombre"},
			}

			for _, project := range projects {
				table = append(table, []string{project.Key, project.Name})
			}

			pterm.DefaultTable.WithHasHeader(true).WithData(table).Render()
		},
	}

	cmd.Flags().StringVarP(&org, "org", "o", "", "Organización de SonarCloud (opcional, usa la configuración por defecto si no se especifica)")
	return cmd
}
