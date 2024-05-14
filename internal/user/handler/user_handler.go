package handler

import (
	"github.com/labstack/echo/v4"
	"ps-halo-suster/internal/user/dto"
	"ps-halo-suster/internal/user/model"
	"ps-halo-suster/internal/user/service"
	"ps-halo-suster/pkg/base/handler"
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

func hasAuthorize(ctx echo.Context, h *UserHandler) error {
	userInfo, err := handler.ExtractUserInfo(ctx)
	if err != nil {
		return err
	}

	_, err = h.userService.GetUserByIdAndRole(userInfo.UserId, userInfo.Role)
	if err != nil {
		return err
	}

	return nil
}

func (h *UserHandler) RegisterIT(ctx echo.Context) *response.WebResponse {
	var request = new(dto.RegisterITReq)
	err := ctx.Bind(&request)

	err = dto.ValidateRegisterITReq(request)
	helper.Panic400IfError(err)

	result, err := h.userService.RegisterUser(dto.NewUserIT(*request))
	helper.PanicIfError(err, "register user failed")

	return &response.WebResponse{
		Status:  201,
		Message: "User registered successfully",
		Data:    result,
	}
}

func (h *UserHandler) RegisterNurse(ctx echo.Context) *response.WebResponse {
	var request = new(dto.RegisterNurseReq)
	err := ctx.Bind(&request)

	err = dto.ValidateRegisterNurseReq(request)
	helper.Panic400IfError(err)

	err = hasAuthorize(ctx, h)
	helper.PanicIfError(err, "user unauthorized")

	result, err := h.userService.RegisterUser(dto.NewUserNurse(*request))
	helper.PanicIfError(err, "register user failed")

	return &response.WebResponse{
		Status:  201,
		Message: "Nurse registered successfully",
		Data: dto.RegisterNurseResp{
			UserId: result.UserId,
			NIP:    result.NIP,
			Name:   result.Name,
		},
	}
}

func (h *UserHandler) LoginIT(ctx echo.Context) *response.WebResponse {
	var request = new(dto.LoginReq)
	err := ctx.Bind(&request)
	request.Role = model.IT
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

func (h *UserHandler) LoginNurse(ctx echo.Context) *response.WebResponse {
	var request = new(dto.LoginReq)
	err := ctx.Bind(&request)
	request.Role = model.NURSE
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
