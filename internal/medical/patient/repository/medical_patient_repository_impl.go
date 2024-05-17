package repository

import (
	"github.com/jmoiron/sqlx"
	"ps-halo-suster/internal/medical/patient/dto"
	"ps-halo-suster/internal/medical/patient/model"
)

type medicalPatientRepository struct {
	db *sqlx.DB
}

func NewMedicalPatientRepositoryImpl(db *sqlx.DB) MedicalPatientRepository {
	return &medicalPatientRepository{db: db}
}

func (m *medicalPatientRepository) CreatePatient(request *dto.MedicalPatientReq) error {
	//TODO implement me
	panic("implement me")
}

func (m *medicalPatientRepository) GetPatients(params *dto.MedicalPatientReqParams) ([]model.MedicalPatient, error) {
	//TODO implement me
	panic("implement me")
}
