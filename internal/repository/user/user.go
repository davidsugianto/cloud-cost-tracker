package user

import (
	"context"
	"time"

	userModel "github.com/davidsugianto/cloud-cost-tracker/internal/model/user"
	"github.com/davidsugianto/cloud-cost-tracker/internal/pkg/auth"
)

func (r *repository) Create(ctx context.Context, row *userModel.User) error {
	err := r.db.WithContext(ctx).Table(userModel.UserTable).Create(&row).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetByUsername(ctx context.Context, username string) (*userModel.User, error) {
	var user userModel.User
	err := r.db.WithContext(ctx).First(&user, username).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetToken(ctx context.Context, id string) (data string, err error) {
	token, err := auth.GenerateToken(id, r.jwtSecret, 24*time.Hour)
	if err != nil {
		return data, err
	}
	return token, nil
}

// func (r *repository) Search(ctx context.Context, req *projectModel.SearchProjectRequest, limit, offset int) (data []projectModel.Project, total int64, err error) {
// 	baseQuery := r.db.WithContext(ctx).Table(projectModel.ProjectTable)

// 	if req.Provider != "" {
// 		baseQuery = baseQuery.Where("provider = ?", req.Provider)
// 	}

// 	if req.ProviderAccount != "" {
// 		baseQuery = baseQuery.Where("provider_account_id = ?", req.ProviderAccount)
// 	}

// 	baseQuery = baseQuery.Unscoped()
// 	countQuery := baseQuery.Count(&total)

// 	fetchList := baseQuery.
// 		Select("*").
// 		Limit(limit).
// 		Offset(offset).
// 		Find(&data)

// 	if fetchList.Error != nil {
// 		return data, total, fetchList.Error
// 	}

// 	if countQuery.Error != nil {
// 		return data, total, countQuery.Error
// 	}
// 	return data, total, nil
// }
