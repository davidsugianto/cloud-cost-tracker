package cost

import (
	"context"

	costModel "github.com/davidsugianto/cloud-cost-tracker/internal/model/cost"
)

func (r *repository) UpsertDaily(ctx context.Context, row *costModel.Cost) error {
	err := r.db.WithContext(ctx).Table(costModel.CostTable).Create(&row).Error
	if err != nil {
		return err
	}
	return nil
}
