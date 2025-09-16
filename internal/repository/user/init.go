package user

import (
	"context"

	userModel "github.com/davidsugianto/cloud-cost-tracker/internal/model/user"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, row *userModel.User) error
	GetByUsername(ctx context.Context, username string) (*userModel.User, error)
	GetToken(ctx context.Context, id string) (data string, err error)
}

type repository struct {
	db        *gorm.DB
	jwtSecret string
}

func New(db *gorm.DB, jwtSecret string) Repository {
	return &repository{
		db:        db,
		jwtSecret: jwtSecret,
	}
}
