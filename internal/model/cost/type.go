package cost

import (
	"time"

	providerModel "github.com/davidsugianto/cloud-cost-tracker/internal/model/provider"
)

type Cost struct {
	ID        string                      `json:"id"`
	ProjectID string                      `json:"project_id"`
	Provider  providerModel.CloudProvider `json:"provider"`
	Service   string                      `json:"service"` // e.g., Compute Engine
	Region    string                      `json:"region"`
	Currency  string                      `json:"currency"`
	Cost      float64                     `json:"cost"`
	Usage     float64                     `json:"usage"` // optional, unit varies by service
	Date      time.Time                   `json:"date"`
}
