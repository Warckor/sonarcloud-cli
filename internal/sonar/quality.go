package sonar

import (
	"fmt"

	"resty.dev/v3"
)

type QualityProfile struct {
	QualityGate struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Default bool   `json:"default"`
	} `json:"qualityGate"`
}

func ListQualityProfiles(client *resty.Client, params map[string]string) ([]QualityProfile, error) {
	var resp []QualityProfile

	_, err := client.R().
		SetQueryParams(params).
		SetResult(&resp).
		Get("/api/qualityprofiles/search")

	if err != nil {
		return nil, fmt.Errorf("error en la llamada a la API de SonarCloud: %w", err)
	}

	return resp, nil
}
