package billing

import (
	"context"

	costModel "github.com/davidsugianto/cloud-cost-tracker/internal/model/cost"
	"github.com/davidsugianto/cloud-cost-tracker/internal/model/provider"
	gcpPkg "github.com/davidsugianto/cloud-cost-tracker/internal/pkg/gcp"
)

type Repository interface {
	FetchCosts(ctx context.Context, provider provider.CloudProvider, projectID string) (data []costModel.Cost, err error)
}

type repository struct {
	gcp *gcpPkg.Service
}

func New(gcp *gcpPkg.Service) Repository {
	return &repository{gcp: gcp}
}
