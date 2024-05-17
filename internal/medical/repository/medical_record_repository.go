package repository

import "ps-halo-suster/internal/medical/dto"

type MedicalRecordRepository interface {
	CreateRecord(request dto.MedicalRecordReq) error
	GetRecords(params dto.MedicalRecordReqParams) ([]dto.MedicalRecordResp, error)
}
