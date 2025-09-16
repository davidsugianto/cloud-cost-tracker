package sync

import (
	"context"

	billingRepo "github.com/davidsugianto/cloud-cost-tracker/internal/repository/billing"
	budgetRepo "github.com/davidsugianto/cloud-cost-tracker/internal/repository/budget"
	costRepo "github.com/davidsugianto/cloud-cost-tracker/internal/repository/cost"
	projectRepo "github.com/davidsugianto/cloud-cost-tracker/internal/repository/project"
)

type Usecase interface {
	SyncCosts(ctx context.Context, projectID string) error
}

type usecase struct {
	projectRepo projectRepo.Repository
	budgetRepo  budgetRepo.Repository
	costRepo    costRepo.Repository
	billingRepo billingRepo.Repository
}

func New(projectRepo projectRepo.Repository, budgetRepo budgetRepo.Repository, costRepo costRepo.Repository, billingRepo billingRepo.Repository) Usecase {
	return &usecase{
		projectRepo: projectRepo,
		budgetRepo:  budgetRepo,
		costRepo:    costRepo,
		billingRepo: billingRepo,
	}
}
