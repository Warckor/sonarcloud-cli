package sonar

import (
	"fmt"

	"resty.dev/v3"
)

type Project struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type listResponse struct {
	Components []Project `json:"components"`
}

func ListProjects(client *resty.Client, org string) ([]Project, error) {
	var resp listResponse
	_, err := client.R().
		SetQueryParam("organization", org).
		SetResult(&resp).
		Get("/api/projects/search")

	if err != nil {
		return nil, fmt.Errorf("error en la llamada a la API de SonarCloud: %w", err)
	}

	return resp.Components, nil
}

func GetProject(client *resty.Client, params map[string]string) (Project, error) {
	var resp listResponse

	_, err := client.R().
		SetQueryParams(params).
		SetResult(&resp).
		Get("/api/projects/search")

	if err != nil {
		return Project{}, fmt.Errorf("error en la llamada a la API de SonarCloud: %w", err)
	}

	if len(resp.Components) == 0 || len(resp.Components) > 1 {
		return Project{}, fmt.Errorf("no se encontraron proyectos con los parámetros proporcionados %v", params)
	}

	return resp.Components[0], nil
}
