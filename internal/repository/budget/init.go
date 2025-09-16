package budget

import (
	// "context"

	// costModel "github.com/davidsugianto/cloud-cost-tracker/internal/model/cost"
	"gorm.io/gorm"
)

type Repository interface {
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}
