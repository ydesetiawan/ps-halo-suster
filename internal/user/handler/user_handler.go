package handler

import (
	"github.com/labstack/echo/v4"
	"ps-halo-suster/internal/user/dto"
	"ps-halo-suster/internal/user/service"
	"ps-halo-suster/pkg/helper"
	"ps-halo-suster/pkg/httphelper/response"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Register(ctx echo.Context) *response.WebResponse {
	var request = new(dto.RegisterReq)
	err := ctx.Bind(&request)

	err = dto.ValidateRegisterReq(request)
	helper.Panic400IfError(err)

	result, err := h.userService.RegisterUser(request)
	helper.PanicIfError(err, "register user failed")

	return &response.WebResponse{
		Status:  201,
		Message: "User registered successfully",
		Data:    result,
	}
}

func (h *UserHandler) Login(ctx echo.Context) *response.WebResponse {
	var request = new(dto.LoginReq)
	err := ctx.Bind(&request)

	err = dto.ValidateLoginReq(request)
	helper.Panic400IfError(err)

	result, err := h.userService.Login(request)
	helper.PanicIfError(err, "failed to login")

	return &response.WebResponse{
		Status:  200,
		Message: "User logged successfully",
		Data:    result,
	}
}
