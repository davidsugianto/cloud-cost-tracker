package project

import (
	"time"

	providerModel "github.com/davidsugianto/cloud-cost-tracker/internal/model/provider"
)

type Project struct {
	ID              string                      `json:"id"`
	Name            string                      `json:"name"`
	Provider        providerModel.CloudProvider `json:"provider"`
	ProviderAccount string                      `json:"provider_account_id"` // FK
	ExternalID      string                      `json:"external_id"`         // e.g., GCP project ID
	CreatedAt       time.Time                   `json:"created_at"`
}
