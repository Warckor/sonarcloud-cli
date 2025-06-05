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
