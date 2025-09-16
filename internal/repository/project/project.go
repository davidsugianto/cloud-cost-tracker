package project

import (
	"context"

	projectModel "github.com/davidsugianto/cloud-cost-tracker/internal/model/project"
)

func (r *repository) Create(ctx context.Context, row *projectModel.Project) error {
	err := r.db.WithContext(ctx).Table(projectModel.ProjectTable).Create(&row).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetByID(ctx context.Context, id string) (*projectModel.Project, error) {
	var project projectModel.Project
	err := r.db.WithContext(ctx).First(&project, id).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *repository) Search(ctx context.Context, req *projectModel.SearchProjectRequest, limit, offset int) (data []projectModel.Project, total int64, err error) {
	baseQuery := r.db.WithContext(ctx).Table(projectModel.ProjectTable)

	if req.Provider != "" {
		baseQuery = baseQuery.Where("provider = ?", req.Provider)
	}

	if req.ProviderAccount != "" {
		baseQuery = baseQuery.Where("provider_account_id = ?", req.ProviderAccount)
	}

	baseQuery = baseQuery.Unscoped()
	countQuery := baseQuery.Count(&total)

	fetchList := baseQuery.
		Select("*").
		Limit(limit).
		Offset(offset).
		Find(&data)

	if fetchList.Error != nil {
		return data, total, fetchList.Error
	}

	if countQuery.Error != nil {
		return data, total, countQuery.Error
	}
	return data, total, nil
}
