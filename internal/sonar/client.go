package sonar

import (
	"log"

	"github.com/joho/godotenv"
	"resty.dev/v3"
)

var (
	token   string
	baseURL string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env no encontrado, se usar'a el entorno actual")
	}

	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal("Error cargando la configuración: ", err)
	}

	token = cfg.Token
	baseURL = cfg.URL

	if token == "" {
		log.Fatal("SONAR_TOKEN no encontrado en el entorno o .env")
	}
	if baseURL == "" {
		log.Fatal("SONAR_BASE_URL no encontrado en el entorno o .env")
	}
}

func NewClient() *resty.Client {
	return resty.New().
		SetBaseURL(baseURL).
		SetAuthToken(token).
		SetHeader("Accept", "application/json")
}
