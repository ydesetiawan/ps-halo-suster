package dto

type MedicalPatientReq struct {
	IdentityNumber      string `json:"identityNumber" validate:"required,len=16,numeric"`
	PhoneNumber         string `json:"phoneNumber" validate:"required,min=10,max=15,startswith=+62"`
	Name                string `json:"name" validate:"required,min=3,max=30"`
	BirthDate           string `json:"birthDate" validate:"required,datetime=2006-01-02"`
	Gender              string `json:"gender" validate:"required,oneof=male female"`
	IdentityCardScanImg string `json:"identityCardScanImg" validate:"required,url"`
}

type MedicalPatientReqParams struct {
	IdentityNumber *int    `query:"identityNumber"`
	Limit          int     `query:"limit"`
	Offset         int     `query:"offset"`
	Name           *string `query:"name"`
	PhoneNumber    *string `query:"phoneNumber"`
	CreatedAt      *string `query:"createdAt"`
}
