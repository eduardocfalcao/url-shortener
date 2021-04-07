package model

type HealthcheckStatus struct {
	Health string `json:"health"`
	Host   string `json:"host"`
}
