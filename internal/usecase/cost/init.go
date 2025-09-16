package cost

import (
	"context"

	costModel "github.com/davidsugianto/cloud-cost-tracker/internal/model/cost"
	pgn "github.com/davidsugianto/cloud-cost-tracker/internal/pkg/pagination"
	costRepo "github.com/davidsugianto/cloud-cost-tracker/internal/repository/cost"
)

type Usecase interface {
	SearchCost(ctx context.Context, req *costModel.SearchCostRequest, pgn pgn.Pagination) (data []costModel.Cost, total int64, err error)
	Summary(ctx context.Context, req *costModel.SummaryRequest) (data []costModel.SummaryCost, err error)
}

type usecase struct {
	costRepo costRepo.Repository
}

func New(costRepo costRepo.Repository) Usecase {
	return &usecase{costRepo: costRepo}
}
