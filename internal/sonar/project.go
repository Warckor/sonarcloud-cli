package sonar

import (
	"fmt"

	"resty.dev/v3"
)

type Project struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type NewProjectParams struct {
	Organization           string `json:"organization"`
	Name                   string `json:"name"`
	Project                string `json:"project"`
	Visibility             string `json:"visibility"`
	NewCodeDefinitionType  string `json:"newCodeDefinitionType"`
	NewCodeDefinitionValue string `json:"newCodeDefinitionValue"`
}

type NewProjectResponse struct {
	Project struct {
		Key        string `json:"key"`
		Name       string `json:"name"`
		Qualifier  string `json:"qualifier"`
		Visibility string `json:"visibility"`
		UUID       string `json:"uuid"`
	} `json:"project"`
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

func GetProject(client *resty.Client, params map[string]string) ([]Project, error) {
	var resp listResponse

	_, err := client.R().
		SetQueryParams(params).
		SetResult(&resp).
		Get("/api/projects/search")

	if err != nil {
		return []Project{}, fmt.Errorf("error en la llamada a la API de SonarCloud: %w", err)
	}

	if len(resp.Components) == 0 {
		return []Project{}, fmt.Errorf("no se encontraron proyectos con los parámetros proporcionados %v", params)
	}

	return resp.Components, nil
}

func CreateProject(client *resty.Client, params NewProjectParams) (NewProjectResponse, error) {
	var resp NewProjectResponse

	_, err := client.R().
		SetFormData(map[string]string{
			"organization":           params.Organization,
			"name":                   params.Name,
			"project":                params.Project,
			"visibility":             params.Visibility,
			"newCodeDefinitionType":  params.NewCodeDefinitionType,
			"newCodeDefinitionValue": params.NewCodeDefinitionValue,
		}).
		SetResult(&resp).
		Post("/api/projects/create")

	if err != nil {
		return NewProjectResponse{}, fmt.Errorf("error en la llamada a la API de SonarCloud: %w", err)
	}

	return resp, nil
}
