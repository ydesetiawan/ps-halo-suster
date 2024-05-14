package dto

import "github.com/go-playground/validator/v10"

type LoginReq struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Password    string `json:"password" validate:"required,min=5,max=15"`
}

func ValidateLoginReq(loginReq LoginReq) error {
	validate := validator.New()
	return validate.Struct(loginReq)
}

type RegisterReq struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,max=30,e164"`
	Password    string `json:"password" validate:"required,min=5,max=15"`
	Name        string `json:"name" validate:"required,min=5,max=50"`
}

func ValidateRegisterReq(req RegisterReq) error {
	validate := validator.New()
	return validate.Struct(req)
}

type RegisterResp struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	AccessToken string `json:"accessToken"`
}
