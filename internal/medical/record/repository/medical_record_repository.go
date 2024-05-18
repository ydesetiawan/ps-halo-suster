package repository

import (
	"ps-halo-suster/internal/medical/record/dto"
	"ps-halo-suster/internal/medical/record/model"
)

type MedicalRecordRepository interface {
	CreateRecord(mRecord *model.MedicalRecord) error
	GetRecords(params *dto.MedicalRecordReqParams) ([]dto.MedicalRecordResp, error)
}
