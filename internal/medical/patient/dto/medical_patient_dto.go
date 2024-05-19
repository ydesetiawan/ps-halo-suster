package dto

import (
	"ps-halo-suster/internal/medical/patient/model"
	"ps-halo-suster/pkg/helper"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type MedicalPatientReq struct {
	IdentityNumber      int    `json:"identityNumber" validate:"required,numeric,identityNumber"`
	PhoneNumber         string `json:"phoneNumber" validate:"required,min=10,max=15,startswith=+62"`
	Name                string `json:"name" validate:"required,min=3,max=30"`
	BirthDate           string `json:"birthDate" validate:"required,validateISO8601"`
	Gender              string `json:"gender" validate:"required,oneof=male female"`
	IdentityCardScanImg string `json:"identityCardScanImg" validate:"required,validUrl"`
}

func ValidateMedicalPatientReq(req *MedicalPatientReq) error {
	validate := validator.New()
	//TODO add validation phone number

	validate.RegisterValidation("identityNumber", func(fl validator.FieldLevel) bool {
		return len(strconv.Itoa(fl.Field().Interface().(int))) <= 16
	})

	validate.RegisterValidation("validUrl", helper.ValidateURL)
	validate.RegisterValidation("validateISO8601", helper.ValidateISO8601)

	return validate.Struct(req)
}

type MedicalPatientReqParams struct {
	IdentityNumber int    `query:"identityNumber"`
	Limit          int    `query:"limit"`
	Offset         int    `query:"offset"`
	Name           string `query:"name"`
	PhoneNumber    string `query:"phoneNumber"`
	CreatedAt      string `query:"createdAt"`
}

func NewMedicalPatient(req MedicalPatientReq) *model.MedicalPatient {
	return &model.MedicalPatient{
		IdentityNumber:      req.IdentityNumber,
		Name:                req.Name,
		PhoneNumber:         req.PhoneNumber,
		BirthDate:           req.BirthDate,
		Gender:              req.Gender,
		IdentityCardScanImg: req.IdentityCardScanImg,
	}
}

type MedicalPatientResp struct {
	IdentityNumber      int    `json:"identityNumber"`
	Name                string `json:"name"`
	PhoneNumber         string `json:"phoneNumber"`
	BirthDate           string `json:"birthDate"`
	Gender              string `json:"gender"`
	IdentityCardScanImg string `json:"identityCardScanImg"`
	CreatedAt           string `json:"createdAt"`
}
