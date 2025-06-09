package components

import (
	"sonar-api/internal/sonar"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
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
			projects, err := sonar.GetProject(client, params)
			if err != nil {
				pterm.Error.Printf("Error al obtener el proyecto: %v\n", err)
				return
			}

			if len(projects) == 0 {
				pterm.Warning.Println("No se encontró el proyecto.")
				return
			}

			table := pterm.TableData{
				{"Clave", "Nombre"},
			}

			for _, project := range projects {
				table = append(table, []string{project.Key, project.Name})
			}

			pterm.DefaultSection.Println("Detalles del Proyecto")
			pterm.DefaultTable.WithHasHeader(true).WithData(table).Render()
		},
	}

	cmd.Flags().StringVarP(&org, "org", "o", "", "Organización de SonarCloud (opcional, usa la configuración por defecto si no se especifica)")
	cmd.Flags().StringVarP(&projectKey, "project-key", "p", "", "Clave del proyecto de SonarCloud")
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
				pterm.Error.Printf("Error al listar proyectos: %v\n", err)
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

			pterm.DefaultSection.Println("Lista de Proyectos")
			pterm.DefaultTable.WithHasHeader(true).WithData(table).Render()
		},
	}

	cmd.Flags().StringVarP(&org, "org", "o", "", "Organización de SonarCloud (opcional, usa la configuración por defecto si no se especifica)")
	return cmd
}

func CreateProjectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "project",
		Short:   "Crea un nuevo proyecto en SonarCloud",
		Long:    `Crea un nuevo proyecto en SonarCloud utilizando la clave y el nombre proporcionados.`,
		Example: `sonarcli create project --org my-org --project-key my-new-project --name "My New Project" --visibility private --code-definition previous_version`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if projectKey == "" || name == "" || visibility == "" || codeDefinition == "" {
				pterm.Error.Printf("Error: Los parámetros --project-key, --name, --visibility y --code-definition son requeridos\n")
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

			params := sonar.NewProjectParams{
				Organization:           finalOrg,
				Name:                   name,
				Project:                projectKey,
				Visibility:             visibility,
				NewCodeDefinitionType:  codeDefinition,
				NewCodeDefinitionValue: codeDefinition,
			}

			client := sonar.NewClient()
			response, err := sonar.CreateProject(client, params)
			if err != nil {
				pterm.Error.Printf("Error al crear el proyecto: %v\n", err)
				return
			}

			table := pterm.TableData{
				{"Clave", "Nombre", "Visibilidad", "UUID"},
				{response.Project.Key, response.Project.Name, response.Project.Visibility, response.Project.UUID},
			}

			pterm.DefaultSection.Println("Proyecto Creado")
			pterm.DefaultTable.WithHasHeader(true).WithData(table).Render()
		},
	}

	cmd.Flags().StringVarP(&org, "org", "o", "", "Organización de SonarCloud (opcional, usa la configuración por defecto si no se especifica)")
	cmd.Flags().StringVarP(&projectKey, "project-key", "p", "", "Clave del nuevo proyecto de SonarCloud")
	cmd.Flags().StringVarP(&name, "name", "n", "", "Nombre del nuevo proyecto de SonarCloud")
	cmd.Flags().StringVarP(&visibility, "visibility", "V", "", "Visibilidad del proyecto (public, private)")
	cmd.Flags().StringVarP(&codeDefinition, "code-definition", "C", "", "Tipo de definición del nuevo código (previous_version, main_branch, specific_version)")
	cmd.MarkFlagsRequiredTogether("project-key", "name", "visibility", "code-definition")

	return cmd
}
