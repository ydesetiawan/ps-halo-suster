package dto

import (
	"github.com/go-playground/validator/v10"
	"ps-halo-suster/internal/user/model"
	"ps-halo-suster/pkg/helper"
	"time"
)

type LoginReq struct {
	NIP      int        `json:"nip" validate:"required,validateNip"`
	Password string     `json:"password" validate:"required,min=5,max=15"`
	Role     model.Role `json:"-"`
}

func ValidateLoginReq(loginReq *LoginReq) error {
	validate := validator.New()
	validate.RegisterValidation("validateNip", helper.ValidateNIP)
	return validate.Struct(loginReq)
}

type RegisterITReq struct {
	NIP      int    `json:"nip" validate:"required,validateNipForIT,validateNip"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=5,max=15"`
	Role     string `json:"-"`
}

func ValidateRegisterITReq(req *RegisterITReq) error {
	validate := validator.New()
	validate.RegisterValidation("validateNipForIT", helper.ValidateNIPForIT)
	validate.RegisterValidation("validateNip", helper.ValidateNIP)
	return validate.Struct(req)
}

type RegisterResp struct {
	UserId      string `json:"userId"`
	NIP         int    `json:"nip"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

func NewUserIT(req RegisterITReq) *model.User {
	return &model.User{
		ID:                  helper.GenerateULID(),
		Name:                req.Name,
		NIP:                 req.NIP,
		Password:            req.Password,
		Role:                string(model.IT),
		IdentityCardScanImg: nil,
	}
}

type RegisterNurseReq struct {
	NIP                 int     `json:"nip" validate:"required,validateNipForNurse,validateNip"`
	Name                string  `json:"name" validate:"required,min=5,max=50"`
	IdentityCardScanImg *string `json:"identityCardScanImg" validate:"required,validateUrl"`
	Role                string  `json:"-"`
}

func ValidateRegisterNurseReq(req *RegisterNurseReq) error {
	validate := validator.New()
	validate.RegisterValidation("validateNipForNurse", helper.ValidateNIPForNurse)
	validate.RegisterValidation("validateNip", helper.ValidateNIP)
	validate.RegisterValidation("validateUrl", helper.ValidateURL)
	return validate.Struct(req)
}

type RegisterNurseResp struct {
	UserId string `json:"userId"`
	NIP    int    `json:"nip"`
	Name   string `json:"name"`
}

func NewUserNurse(req RegisterNurseReq) *model.User {
	return &model.User{
		ID:                  helper.GenerateULID(),
		Name:                req.Name,
		NIP:                 req.NIP,
		IdentityCardScanImg: req.IdentityCardScanImg,
		Role:                string(model.NURSE),
	}
}

type UpdateUserReq struct {
	NIP  int    `json:"nip" validate:"required,validateNip"`
	Name string `json:"name" validate:"required,min=5,max=50"`
	ID   string `json:"userId" validate:"required"`
}

func ValidateUpdateUserReq(req *UpdateUserReq) error {
	validate := validator.New()
	validate.RegisterValidation("validateNip", helper.ValidateNIP)
	return validate.Struct(req)
}

type GrantAccessReq struct {
	Password string `json:"password" validate:"required,min=5,max=33"`
	ID       string `json:"-"`
}

func ValidateGrantAccessReq(req *GrantAccessReq) error {
	validate := validator.New()
	return validate.Struct(req)
}

type GetNurseParams struct {
	UserId    string `query:"userId"`
	Limit     int    `query:"limit"`
	Offset    int    `query:"offset"`
	Name      string `query:"name"`
	NIP       int    `query:"nip"`
	Role      string `query:"role"`
	CreatedAt string `query:"createdAt"`
}

type GetNurseResp struct {
	UserId        string    `json:"userId"`
	NIP           int       `json:"nip"`
	Name          string    `json:"name"`
	CreatedAt     string    `json:"createdAt"`
	CreatedAtTime time.Time `json:"-"`
}
