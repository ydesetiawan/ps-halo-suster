package dto

import (
	"github.com/go-playground/validator/v10"
	"ps-halo-suster/pkg/helper"
)

type LoginReq struct {
	NIP      int    `json:"nip" validate:"required,validateNip"`
	Password string `json:"password" validate:"required,min=5,max=15"`
}

func ValidateLoginReq(loginReq *LoginReq) error {
	validate := validator.New()
	validate.RegisterValidation("validateNip", helper.ValidateNIP)
	return validate.Struct(loginReq)
}

type RegisterReq struct {
	NIP      int    `json:"nip" validate:"required,validateNip"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=5,max=15"`
	Role     string `json:"-"`
}

func ValidateRegisterReq(req *RegisterReq) error {
	validate := validator.New()
	validate.RegisterValidation("validateNip", helper.ValidateNIP)
	return validate.Struct(req)
}

type RegisterResp struct {
	UserId      string `json:"userId"`
	NIP         int    `json:"nip"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}
