package cost

import (
	"context"

	costModel "github.com/davidsugianto/cloud-cost-tracker/internal/model/cost"
)

func (r *repository) Search(ctx context.Context, req *costModel.SearchCostRequest, limit, offset int) (data []costModel.Cost, total int64, err error) {
	// baseQuery := r.db.WithContext(ctx).Table(costModel.CostTable)

	// if req.ProjectID != "" {
	// 	baseQuery = baseQuery.Where("project_id = ?", req.ProjectID)
	// }

	// baseQuery = baseQuery.Unscoped()
	// countQuery := baseQuery.Count(&total)

	// fetchList := baseQuery.
	// 	Select("*").
	// 	Limit(limit).
	// 	Offset(offset).
	// 	Find(&data)

	// if fetchList.Error != nil {
	// 	return data, total, fetchList.Error
	// }

	// if countQuery.Error != nil {
	// 	return data, total, countQuery.Error
	// }
	return data, total, nil
}

func (r *repository) Summary(ctx context.Context, req *costModel.SummaryRequest) (data []costModel.SummaryCost, err error) {
	return data, nil
}
