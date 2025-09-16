package cost

import (
	"context"

	costModel "github.com/davidsugianto/cloud-cost-tracker/internal/model/cost"
	"gorm.io/gorm"
)

type Repository interface {
	UpsertDaily(ctx context.Context, row *costModel.Cost) error
	Search(ctx context.Context, req *costModel.SearchCostRequest, limit, offset int) (data []costModel.Cost, total int64, err error)
	Summary(ctx context.Context, req *costModel.SummaryRequest) (data []costModel.SummaryCost, err error)
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}
