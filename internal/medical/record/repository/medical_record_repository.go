package repository

import (
	"ps-halo-suster/internal/medical/record/dto"
)

type MedicalRecordRepository interface {
	CreateRecord(request *dto.MedicalRecordReq) error
	GetRecords(params *dto.MedicalRecordReqParams) ([]dto.MedicalRecordResp, error)
}
