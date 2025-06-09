package sonar

import (
	"fmt"

	"resty.dev/v3"
)

type QualityProfile struct {
	QualityGates []struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		IsDefault  bool   `json:"isdefault"`
		Conditions []struct {
			ID     int    `json:"id"`
			Metric string `json:"metric"`
			Op     string `json:"op"`
			Error  string `json:"error"`
		} `json:"conditions"`
	} `json:"qualityGates"`
}

type QualityProfileProject struct {
	QualityGate struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Default bool   `json:"default"`
	}
}

func ListQualityProfiles(client *resty.Client, params map[string]string) (QualityProfile, error) {
	var resp QualityProfile

	_, err := client.R().
		SetQueryParams(params).
		SetResult(&resp).
		Get("/api/qualitygates/list")

	if err != nil {
		return QualityProfile{}, fmt.Errorf("error en la llamada a la API de SonarCloud: %w", err)
	}

	return resp, nil
}

func GetQualityProfile(client *resty.Client, params map[string]string) (QualityProfileProject, error) {
	var resp QualityProfileProject

	_, err := client.R().
		SetQueryParams(params).
		SetResult(&resp).
		Get("/api/qualitygates/get_by_project")

	if err != nil {
		return QualityProfileProject{}, fmt.Errorf("error en la llamada a la API de SonarCloud: %w", err)
	}

	return resp, nil
}
