package handler

import (
	"github.com/labstack/echo/v4"
	"ps-halo-suster/internal/user/dto"
	"ps-halo-suster/internal/user/model"
	"ps-halo-suster/internal/user/service"
	"ps-halo-suster/pkg/base/handler"
	"ps-halo-suster/pkg/errs"
	"ps-halo-suster/pkg/helper"
	"ps-halo-suster/pkg/httphelper/response"
	"strconv"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func hasAuthorizeRoleIT(ctx echo.Context, h *UserHandler) error {
	userInfo, err := handler.ExtractUserInfo(ctx)
	if err != nil {
		return err
	}

	_, err = h.userService.GetUserByIdAndRole(userInfo.UserId, string(model.IT))
	if err != nil {
		return err
	}

	return nil
}

func (h *UserHandler) RegisterNurse(ctx echo.Context) *response.WebResponse {
	var request = new(dto.RegisterNurseReq)
	err := ctx.Bind(&request)

	err = dto.ValidateRegisterNurseReq(request)
	helper.Panic400IfError(err)

	err = hasAuthorizeRoleIT(ctx, h)
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

func (h *UserHandler) LoginIT(ctx echo.Context) *response.WebResponse {
	return h.LoginUser(ctx, model.IT)
}

func (h *UserHandler) LoginNurse(ctx echo.Context) *response.WebResponse {
	return h.LoginUser(ctx, model.NURSE)
}

func (h *UserHandler) LoginUser(ctx echo.Context, role model.Role) *response.WebResponse {
	var request = new(dto.LoginReq)
	err := ctx.Bind(&request)
	request.Role = role
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

func (h *UserHandler) GetNurses(ctx echo.Context) *response.WebResponse {
	err := hasAuthorizeRoleIT(ctx, h)
	helper.PanicIfError(err, "user unauthorized")

	var params = new(dto.GetNurseParams)
	err = ctx.Bind(params)

	results, err := h.userService.GetNurses(params)
	helper.PanicIfError(err, "failed to GetNurses")
	if len(results) == 0 {
		results = []dto.GetNurseResp{}
	}

	return &response.WebResponse{
		Status:  200,
		Message: "get nurses successfully",
		Data:    results,
	}
}

func (h *UserHandler) UpdateNurse(ctx echo.Context) *response.WebResponse {
	err := hasAuthorizeRoleIT(ctx, h)
	helper.PanicIfError(err, "user unauthorized")

	var request = new(dto.UpdateUserReq)
	err = ctx.Bind(&request)
	request.ID = ctx.Param("userId")

	err = dto.ValidateUpdateUserReq(request)
	helper.Panic400IfError(err)

	if !helper.ValidatePrefixNurse(strconv.Itoa(request.NIP)) {
		helper.Panic404IfError(errs.NewErrDataNotFound1(" user is not a nurse (nip not starts with 303)"))
	}

	err = h.userService.UpdateNurse(request)
	helper.PanicIfError(err, "failed to UpdateNurse")

	return &response.WebResponse{
		Status:  200,
		Message: "User update successfully",
	}
}

func (h *UserHandler) DeleteNurse(ctx echo.Context) *response.WebResponse {
	err := hasAuthorizeRoleIT(ctx, h)
	helper.PanicIfError(err, "user unauthorized")

	nurseId := ctx.Param("userId")
	err = h.userService.DeleteNurse(nurseId)
	helper.PanicIfError(err, "failed to DeleteNurse")

	return &response.WebResponse{
		Status:  200,
		Message: "User delete successfully",
	}
}

func (h *UserHandler) GrantAccessNurse(ctx echo.Context) *response.WebResponse {
	err := hasAuthorizeRoleIT(ctx, h)
	helper.PanicIfError(err, "user unauthorized")

	var request = new(dto.GrantAccessReq)
	err = ctx.Bind(&request)
	request.ID = ctx.Param("userId")

	err = dto.ValidateGrantAccessReq(request)
	helper.Panic400IfError(err)

	err = h.userService.GrantAccessNurse(request)
	helper.PanicIfError(err, "failed to UpdateNurse")

	return &response.WebResponse{
		Status:  200,
		Message: "User update successfully",
	}
}
