package project

import (
	"context"

	projectModel "github.com/davidsugianto/cloud-cost-tracker/internal/model/project"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, row *projectModel.Project) error
	GetByID(ctx context.Context, id string) (*projectModel.Project, error)
	Search(ctx context.Context, req *projectModel.SearchProjectRequest, limit, offset int) (data []projectModel.Project, total int64, err error)
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}
