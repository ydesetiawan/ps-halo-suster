package repository

import (
	"github.com/jmoiron/sqlx"
	"ps-halo-suster/internal/medical/dto"
)

type medicalRecordRepository struct {
	db *sqlx.DB
}

func NewMedicalRecordRepositoryImpl(db *sqlx.DB) MedicalRecordRepository {
	return &medicalRecordRepository{db: db}
}

func (m medicalRecordRepository) CreateRecord(request dto.MedicalRecordReq) error {
	//TODO implement me
	panic("implement me")
}

func (m medicalRecordRepository) GetRecords(params dto.MedicalRecordReqParams) ([]dto.MedicalRecordResp, error) {
	//TODO implement me
	panic("implement me")
}
