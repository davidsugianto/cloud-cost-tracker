package project

import (
	"time"

	providerModel "github.com/davidsugianto/cloud-cost-tracker/internal/model/provider"
)

const (
	ProjectTable = "projects"
)

type Project struct {
	ID              string                      `json:"id" gorm:"primaryKey"`
	Name            string                      `json:"name" gorm:"not null"`
	Provider        providerModel.CloudProvider `json:"provider"`
	ProviderAccount string                      `json:"provider_account_id"` // FK
	ExternalID      string                      `json:"external_id"`         // e.g., GCP project ID
	CreatedAt       time.Time                   `json:"created_at"`
}

type SearchProjectRequest struct {
	Provider        string `json:"provider"`
	ProviderAccount string `json:"provider_account_id"`
}
