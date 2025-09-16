package user

import (
	"context"

	userModel "github.com/davidsugianto/cloud-cost-tracker/internal/model/user"
	userRepo "github.com/davidsugianto/cloud-cost-tracker/internal/repository/user"
)

type Usecase interface {
	SignUp(ctx context.Context, req *userModel.SignUpRequest) (data userModel.SignUpResponse, err error)
	SignIn(ctx context.Context, req *userModel.SignInRequest) (data userModel.SignInResponse, err error)
}

type usecase struct {
	userRepo userRepo.Repository
}

func New(userRepo userRepo.Repository) Usecase {
	return &usecase{
		userRepo: userRepo,
	}
}
