package dto

import (
	"github.com/go-playground/validator/v10"
	"strconv"
)

type MedicalRecordReq struct {
	IdentityNumber int64  `json:"identityNumber" validate:"required,identityNumberLength"`
	Symptoms       string `json:"symptoms" validate:"required,min=1,max=2000"`
	Medications    string `json:"medications" validate:"required,min=1,max=2000"`
	UserId         string `json:"-"`
}

func identityNumberLength(fl validator.FieldLevel) bool {
	identityNumber := strconv.FormatInt(fl.Field().Int(), 10)
	return len(identityNumber) == 16
}

func ValidateMedicalRecordReq(req *MedicalRecordReq) error {
	validate := validator.New()
	validate.RegisterValidation("identityNumberLength", identityNumberLength)
	return validate.Struct(req)
}

type IdentityDetail struct {
	IdentityNumber      int64  `json:"identityNumber"`
	PhoneNumber         string `json:"phoneNumber"`
	Name                string `json:"name"`
	BirthDate           string `json:"birthDate"`
	Gender              string `json:"gender"`
	IdentityCardScanImg string `json:"identityCardScanImg"`
}

type CreatedBy struct {
	NIP    string `json:"nip"`
	Name   string `json:"name"`
	UserID string `json:"userId"`
}

type MedicalRecordResp struct {
	IdentityDetail IdentityDetail `json:"identityDetail"`
	Symptoms       string         `json:"symptoms"`
	Medications    string         `json:"medications"`
	CreatedAt      string         `json:"createdAt"`
	CreatedBy      CreatedBy      `json:"createdBy"`
}

type MedicalRecordReqParams struct {
	IdentityNumber int    `query:"identityDetail.identityNumber"`
	UserID         string `query:"createdBy.userId"`
	NIP            string `query:"createdBy.nip"`
	Limit          int    `query:"limit"`
	Offset         int    `query:"offset"`
	CreatedAt      string `query:"createdAt"`
}
