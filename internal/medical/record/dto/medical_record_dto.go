package dto

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type MedicalRecordReq struct {
	IdentityNumber string `json:"identityNumber" validate:"required,len=16,numeric"`
	Symptoms       string `json:"symptoms" validate:"required,min=1,max=2000"`
	Medications    string `json:"medications" validate:"required,min=1,max=2000"`
}

func ValidateMedicalRecordReq(req *MedicalRecordReq) error {
	validate := validator.New()
	return validate.Struct(req)
}

type IdentityDetail struct {
	IdentityNumber      int    `json:"identityNumber"`
	PhoneNumber         string `json:"phoneNumber"`
	Name                string `json:"name"`
	BirthDate           string `json:"birthDate"`
	Gender              string `json:"gender"`
	IdentityCardScanImg string `json:"identityCardScanImg"`
}

// CreatedBy represents the creator details of a record
type CreatedBy struct {
	NIP    int    `json:"nip"`
	Name   string `json:"name"`
	UserID string `json:"userId"`
}

type MedicalRecordResp struct {
	IdentityDetail IdentityDetail `json:"identityDetail"`
	Symptoms       string         `json:"symptoms"`
	Medications    string         `json:"medications"`
	CreatedAt      time.Time      `json:"createdAt"`
	CreatedBy      CreatedBy      `json:"createdBy"`
}

type IdentityDetailParams struct {
	IdentityNumber int `query:"identityNumber"`
}

type CreatedByParams struct {
	UserID string `query:"userId"`
	NIP    string `query:"nip"`
}

type MedicalRecordReqParams struct {
	IdentityDetail IdentityDetailParams `query:"identityDetail"`
	CreatedBy      CreatedByParams      `query:"createdBy"`
	Limit          int                  `query:"limit"`
	Offset         int                  `query:"offset"`
	CreatedAt      string               `query:"createdAt"`
}
