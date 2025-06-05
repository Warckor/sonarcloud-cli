package components

import (
	"sonar-api/internal/sonar"
	"strconv"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var (
	org        string
	projectKey string
)

func ListQualityCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "quality",
		Short:   "Listar perfiles de calidad",
		Long:    `Lista los perfiles de calidad disponibles en SonarCloud.`,
		Example: `sonarcli get quality --org my-org`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if projectKey == "" {
				pterm.Error.Printf("Error: se requiere la clave del proyecto\n")
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

			if projectKey == "" {
				pterm.Error.Printf("Error: se requiere la clave del proyecto\n")
				return
			}

			params := map[string]string{
				"organization": finalOrg,
				"project":      projectKey,
			}

			client := sonar.NewClient()
			profiles, err := sonar.ListQualityProfiles(client, params)
			if err != nil {
				pterm.Error.Printf("Error al obtener los perfiles de calidad: %v\n", err)
				return
			}

			if len(profiles) == 0 {
				pterm.Info.Println("No se encontraron perfiles de calidad.")
				return
			}

			table := pterm.TableData{
				{"ID", "Nombre", "Descripción"},
			}
			for _, profile := range profiles {
				table = append(table, []string{
					profile.QualityGate.ID,
					profile.QualityGate.Name,
					strconv.FormatBool(profile.QualityGate.Default),
				})
			}

			pterm.DefaultTable.WithHasHeader(true).WithData(table).Render()
		},
	}

	cmd.Flags().StringVarP(&org, "org", "o", "", "Organización de SonarCloud (opcional, usa la configuración por defecto si no se especifica)")
	cmd.Flags().StringVarP(&projectKey, "projectKey", "p", "", "Clave del proyecto de SonarCloud")
	cmd.MarkFlagsOneRequired("projectKey")

	return cmd
}
