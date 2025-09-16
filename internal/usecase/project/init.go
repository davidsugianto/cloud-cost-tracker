package project

import (
	// projectModel "github.com/davidsugianto/cloud-cost-tracker/internal/model/project"
	projectRepo "github.com/davidsugianto/cloud-cost-tracker/internal/repository/project"
)

type Usecase interface {
	// CreateUser(user *userModel.User) error
	// GetUserByID(id uint) (*userModel.User, error)
	// GetAllUsers() ([]userModel.User, error)
	// UpdateUser(user *userModel.User) error
	// DeleteUser(id uint) error
}

type usecase struct {
	projectRepo projectRepo.Repository
}

func New(projectRepo projectRepo.Repository) Usecase {
	return &usecase{projectRepo: projectRepo}
}
