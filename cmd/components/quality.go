package components

import (
	"encoding/json"
	"sonar-api/internal/sonar"
	"strconv"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var (
	output string
)

func ListQualityCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "quality",
		Short:   "Listar perfiles de calidad",
		Long:    `Lista los perfiles de calidad disponibles en SonarCloud.`,
		Example: `sonarcli get quality --org my-org --project-key my-project-key`,
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := sonar.LoadConfig()
			if err != nil {
				pterm.Error.Printf("Error al cargar la configuración: %v\n", err)
				return
			}

			finalOrg := org
			if finalOrg == "" {
				finalOrg = cfg.Organization
			}
			if finalOrg == "" {
				pterm.Error.Printf("Error: se requiere la organización\n")
				return
			}

			params := map[string]string{
				"organization": finalOrg,
			}

			client := sonar.NewClient()
			profiles, err := sonar.ListQualityProfiles(client, params)
			if err != nil {
				pterm.Error.Printf("Error al obtener los perfiles de calidad: %v\n", err)
				return
			}

			if len(profiles.QualityGates) == 0 {
				pterm.Info.Println("No se encontraron perfiles de calidad.")
				return
			}

			table := pterm.TableData{
				{"ID", "Nombre", "Descripción"},
			}

			if output == "json" {
				jsonBytes, err := json.MarshalIndent(profiles, "", "  ")
				if err != nil {
					pterm.Error.Printf("Error al serializar a JSON: %v\n", err)
					return
				}
				pterm.Println(string(jsonBytes))
				return
			}

			for _, profile := range profiles.QualityGates {
				table = append(table, []string{
					strconv.FormatInt(int64(profile.ID), 10),
					profile.Name,
					strconv.FormatBool(profile.IsDefault),
				})
			}

			pterm.DefaultTable.WithHasHeader(true).WithData(table).Render()
		},
	}

	cmd.Flags().StringVarP(&org, "org", "o", "", "Organización de SonarCloud (opcional, usa la configuración por defecto si no se especifica)")
	cmd.Flags().StringVar(&output, "output", "", "Formato de salida (json, table)")

	return cmd
}

func GetQualityCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "quality",
		Short:   "Obtener perfil de calidad por proyecto",
		Long:    `Obtiene el perfil de calidad asociado a un proyecto en SonarCloud.`,
		Example: `sonarcli get quality --org my-org --project-key my-project-key`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if projectKey == "" {
				pterm.Error.Println("Error: se requiere la clave del proyecto")
				return cmd.Help()
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := sonar.LoadConfig()
			if err != nil {
				pterm.Error.Printf("Error al cargar la configuración: %v\n", err)
				return
			}

			finalOrg := org
			if finalOrg == "" {
				finalOrg = cfg.Organization
			}
			if finalOrg == "" {
				pterm.Error.Printf("Error: se requiere la organización\n")
				return
			}

			params := map[string]string{
				"organization": finalOrg,
				"project":      projectKey,
			}

			client := sonar.NewClient()
			profile, err := sonar.GetQualityProfile(client, params)
			if err != nil {
				pterm.Error.Printf("Error al obtener el perfil de calidad: %v\n", err)
				return
			}

			if profile.QualityGate.ID == 0 {
				pterm.Error.Println("No se encontró un perfil de calidad para el proyecto especificado.")
				return
			}

			table := pterm.TableData{
				{"ID", "Nombre", "Predeterminado"},
				{strconv.FormatInt(int64(profile.QualityGate.ID), 10), profile.QualityGate.Name, strconv.FormatBool(profile.QualityGate.Default)},
			}

			pterm.DefaultTable.WithHasHeader(true).WithData(table).Render()

		},
	}

	cmd.Flags().StringVarP(&org, "org", "o", "", "Organización de SonarCloud (opcional, usa la configuración por defecto si no se especifica)")
	cmd.Flags().StringVarP(&projectKey, "project-key", "p", "", "Clave del proyecto de SonarCloud (requerido)")
	cmd.MarkFlagRequired("project-key")

	return cmd
}

func StatusQualityCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "quality",
		Short:   "Verificar el estado del perfil de calidad de un proyecto",
		Long:    `Verifica el estado del perfil de calidad asociado a un proyecto en SonarCloud.`,
		Example: `sonarcli status quality --project-key my-project-key --branch my-branch`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if projectKey == "" {
				pterm.Error.Println("Error: se requiere la clave del proyecto")
				return cmd.Help()
			}
			if branch == "" {
				branch = "main"
				return nil
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			params := map[string]string{
				"projectKey": projectKey,
			}

			if branch != "" {
				params["branch"] = branch
			}

			client := sonar.NewClient()
			response, err := sonar.StatusQualityGate(client, params)
			if err != nil {
				pterm.Error.Printf("Error al obtener el estado del perfil de calidad: %v\n", err)
				return
			}

			if len(response.ProjectStatus.Conditions) == 0 {
				pterm.Error.Println("No se encontraron condiciones para el perfil de calidad del proyecto especificado.")
				return
			}

			table := pterm.TableData{
				{"Métrica", "Estado", "Valor Actual", "Umbral de Error", "Comparador"},
			}

			for _, condition := range response.ProjectStatus.Conditions {
				table = append(table, []string{
					condition.MetricKey,
					condition.Status,
					condition.ActualValue,
					condition.ErrorThreshold,
					condition.Comparator,
				})
			}

			pterm.DefaultSection.Println("Estado del Perfil de Calidad del Proyecto")
			pterm.DefaultTable.WithHasHeader(true).WithData(table).Render()
		},
	}

	cmd.Flags().StringVarP(&projectKey, "project-key", "p", "", "Clave del proyecto de SonarCloud (requerido)")
	cmd.Flags().StringVarP(&branch, "branch", "b", "", "Nombre de la rama del proyecto (opcional, por defecto es 'main')")
	cmd.MarkFlagRequired("project-key")

	return cmd
}
