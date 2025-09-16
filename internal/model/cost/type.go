package cost

import (
	"time"

	providerModel "github.com/davidsugianto/cloud-cost-tracker/internal/model/provider"
	pgn "github.com/davidsugianto/cloud-cost-tracker/internal/pkg/pagination"
)

const (
	CostTable = "costs"
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

type SearchCostRequest struct {
	ProjectID string                       `json:"project_id"`
	Provider  *providerModel.CloudProvider `json:"provider"`
	Service   *string                      `json:"service"`
	Region    *string                      `json:"region"`
	From      time.Time                    `json:"from"`
	To        time.Time                    `json:"to"`
}

type SearchCostResponse struct {
	Metadata   []Cost          `json:"metadata,omitempty"`
	Pagination *pgn.Pagination `json:"pagination,omitempty"`
}

type SummaryRequest struct {
	GroupBy   string    `json:"group_by"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
	ProjectID *string   `json:"project_id"`
}

type SummaryCost struct {
	Key       string  `json:"key"`
	TotalCost float64 `json:"total_cost"`
	Currency  string  `json:"currency"`
}
