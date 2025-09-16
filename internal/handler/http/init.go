package http

import (
	user "github.com/davidsugianto/cloud-cost-tracker/internal/usecase/user"
)

type Handler struct {
	userUseCase user.Usecase
	// costUseCase cost.Usecase
	// syncUseCase sync.Usecase
}

type Dependencies struct {
	UserUseCase user.Usecase
	// costUseCase cost.Usecase
	// syncUseCase sync.Usecase
}

func New(deps Dependencies) *Handler {
	return &Handler{
		userUseCase: deps.UserUseCase,
		// costUseCase: deps.costUseCase,
		// syncUseCase: deps.syncUseCase,
	}
}
