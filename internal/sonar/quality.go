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

type QualityGateStatus struct {
	ProjectStatus struct {
		Status     string `json:"status"`
		Conditions []struct {
			Status         string `json:"status"`
			MetricKey      string `json:"metricKey"`
			Comparator     string `json:"comparator"`
			PeriodIndex    int    `json:"periodIndex"`
			ErrorThreshold string `json:"errorThreshold"`
			ActualValue    string `json:"actualValue"`
		} `json:"conditions"`
		Periods []struct {
			Index int    `json:"index"`
			Mode  string `json:"mode"`
			Date  string `json:"date"`
		} `json:"periods"`
		IgnoredConditions bool `json:"ignoredConditions"`
	} `json:"projectStatus"`
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

func StatusQualityGate(client *resty.Client, params map[string]string) (QualityGateStatus, error) {
	var resp QualityGateStatus

	_, err := client.R().
		SetQueryParams(params).
		SetResult(&resp).
		Get("/api/qualitygates/project_status")

	if err != nil {
		return QualityGateStatus{}, fmt.Errorf("error en la llamada a la API de SonarCloud: %w", err)
	}

	return resp, nil
}
