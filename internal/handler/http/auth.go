package http

import (
	"net/http"

	userModel "github.com/davidsugianto/cloud-cost-tracker/internal/model/user"
	"github.com/davidsugianto/cloud-cost-tracker/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var req userModel.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err)
		return
	}

	data, err := h.userUseCase.SignUp(c, &req)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err)
		return
	}
	response.OK(c, data)
}

func (h *Handler) Login(c *gin.Context) {
	var req userModel.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err)
		return
	}

	data, err := h.userUseCase.SignIn(c, &req)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err)
		return
	}
	response.OK(c, data)
}
