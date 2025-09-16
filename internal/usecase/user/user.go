package user

import (
	"context"

	userModel "github.com/davidsugianto/cloud-cost-tracker/internal/model/user"
	"github.com/davidsugianto/cloud-cost-tracker/internal/pkg/auth"
	"github.com/google/uuid"
)

func (u *usecase) SignUp(ctx context.Context, req *userModel.SignUpRequest) (data userModel.SignUpResponse, err error) {
	hash, err := auth.HashPassword(req.Password)
	if err != nil {
		return data, err
	}

	user := newUser(req.Username, hash)

	if err := u.userRepo.Create(ctx, user); err != nil {
		return data, err
	}

	result := userModel.SignUpResponse{
		ID:       user.ID,
		Username: user.Username,
	}

	return result, nil
}

func (u *usecase) SignIn(ctx context.Context, req *userModel.SignInRequest) (data userModel.SignInResponse, err error) {
	user, err := u.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return data, err
	}

	if err := auth.CheckPassword(user.Password, req.Password); err != nil {
		return data, err
	}

	token, err := u.userRepo.GetToken(ctx, user.ID)
	if err != nil {
		return data, nil
	}

	result := userModel.SignInResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    token,
	}

	return result, nil
}

func newUser(username, passwordHash string) *userModel.User {
	return &userModel.User{
		ID:       uuid.New().String(),
		Username: username,
		Password: passwordHash,
	}
}
