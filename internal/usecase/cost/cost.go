package cost

import (
	"context"

	costModel "github.com/davidsugianto/cloud-cost-tracker/internal/model/cost"
	pgn "github.com/davidsugianto/cloud-cost-tracker/internal/pkg/pagination"
)

func (u *usecase) SearchCost(ctx context.Context, req *costModel.SearchCostRequest, pgn pgn.Pagination) (data []costModel.Cost, total int64, err error) {
	result, count, err := u.costRepo.Search(ctx, req, pgn.Limit(), pgn.Offset())
	if err != nil {
		return data, total, err
	}
	return result, count, nil
}

func (u *usecase) Summary(ctx context.Context, req *costModel.SummaryRequest) (data []costModel.SummaryCost, err error) {
	if req.GroupBy == "" {
		req.GroupBy = "project"
	}
	return u.costRepo.Summary(ctx, req)
}
