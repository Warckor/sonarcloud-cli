package sonar

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Token        string `yaml:"token"`
	URL          string `yaml:"url"`
	Organization string `yaml:"organization"`
}

func getConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, "sonarcli", "config.yaml")
}

func LoadConfig() (*Config, error) {
	path := getConfigPath()
	data, err := os.ReadFile(path)

	// Si el archivo no existe, intentamos cargar desde variables de entorno
	if err != nil {
		// Intentar cargar desde variables de entorno
		var cfg Config
		cfg.Token = os.Getenv("SONAR_TOKEN")
		cfg.URL = os.Getenv("SONAR_BASE_URL")
		cfg.Organization = os.Getenv("SONAR_ORGANIZATION")

		// Si tenemos todas las variables de entorno, usarlas
		if cfg.Token != "" && cfg.URL != "" && cfg.Organization != "" {
			return &cfg, nil
		}

		// Si el archivo no existe y no tenemos variables de entorno, sugerir configuración
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("no se encontró el archivo de configuración; ejecuta 'go run main.go config set --token <token> --url <url> --org <org>' para configurar")
		}

		return nil, fmt.Errorf("configuración incompleta; asegúrate de tener SONAR_TOKEN, SONAR_BASE_URL y SONAR_ORGANIZATION definidos en tu entorno o en el archivo de configuración")
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func SaveConfig(cfg *Config) error {
	path := getConfigPath()
	dirPath := filepath.Dir(path)

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return fmt.Errorf("error al crear directorio de configuración: %w", err)
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("error al serializar configuración: %w", err)
	}

	return os.WriteFile(path, data, 0644)
}
