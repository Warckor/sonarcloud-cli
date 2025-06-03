package cmd

import (
	"fmt"
	"log"
	"sonar-api/internal/sonar"

	"github.com/spf13/cobra"
)

var (
	cfgToken string
	cfgUrl   string
	cfgOrg   string
)

var configCmd = &cobra.Command{
	Use:     "config",
	Short:   "Gestión de la configuración",
	Long:    `Permite configurar el token, URL y organización de SonarCloud para su uso en la CLI.`,
	Example: `sonarcli config set --token <your-token> --url https://sonarcloud.io --org <your-org>`,
}

var configSetCmd = &cobra.Command{
	Use:     "set",
	Short:   "Guarda la configuración en # ~/.sonarcli/config.yaml",
	Long:    `Guarda la configuración de SonarCloud (token, URL y organización) en un archivo de configuración local para su uso posterior.`,
	Example: `sonarcli config set --token <your-token> --url https://sonarcloud.io --org <your-org>`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := &sonar.Config{
			Token:        cfgToken,
			URL:          cfgUrl,
			Organization: cfgOrg,
		}

		if err := sonar.SaveConfig(cfg); err != nil {
			log.Fatal("Error guardando config: ", err)
		}

		fmt.Println("Configuración guardada correctamente.")
	},
}

var configShowCmd = &cobra.Command{
	Use:     "show",
	Short:   "Muestra la configuración actual",
	Long:    `Muestra la configuración actual de SonarCloud (token, URL y organización).`,
	Example: `sonarcli config show`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := sonar.LoadConfig()
		if err != nil {
			log.Fatal("Error cargando config: ", err)
		}

		fmt.Printf("🛠️ Configuración actual:\n")
		fmt.Printf("  🗝️ Token: %s\n", cfg.Token)
		fmt.Printf("  🔗 URL: %s\n", cfg.URL)
		fmt.Printf("  🏢 Organization: %s\n", cfg.Organization)
	},
}

func init() {
	configSetCmd.Flags().StringVar(&cfgToken, "token", "", "Token de acceso a SonarCLoud")
	configSetCmd.Flags().StringVar(&cfgUrl, "url", "https://sonarcloud.io", "URL de SonarCLoud (por defecto: https://sonarcloud.io)")
	configSetCmd.Flags().StringVar(&cfgOrg, "org", "", "Organización de SonarCLoud")

	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configShowCmd)
	rootCmd.AddCommand(configCmd)
}
