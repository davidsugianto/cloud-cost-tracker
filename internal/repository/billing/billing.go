package billing

import (
	"context"
	"time"

	costModel "github.com/davidsugianto/cloud-cost-tracker/internal/model/cost"
	"github.com/davidsugianto/cloud-cost-tracker/internal/model/provider"
)

func (r *repository) FetchCosts(ctx context.Context, provider provider.CloudProvider, projectID string) (data []costModel.Cost, err error) {
	return []costModel.Cost{
		{
			ProjectID: projectID,
			Provider:  "gcp",
			Service:   "compute-engine",
			Region:    "us-central1",
			Currency:  "USD",
			Cost:      23.50,
			Usage:     120.0,
			Date:      time.Now().AddDate(0, 0, -1),
		},
		{
			ProjectID: projectID,
			Provider:  "gcp",
			Service:   "cloud-storage",
			Region:    "us-central1",
			Currency:  "USD",
			Cost:      5.75,
			Usage:     45.0,
			Date:      time.Now().AddDate(0, 0, -1),
		},
	}, nil
}
