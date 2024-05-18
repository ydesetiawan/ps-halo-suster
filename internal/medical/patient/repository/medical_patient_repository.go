package repository

import (
	"ps-halo-suster/internal/medical/patient/dto"
	"ps-halo-suster/internal/medical/patient/model"
)

type MedicalPatientRepository interface {
	CreatePatient(medicalPatient model.MedicalPatient) (model.MedicalPatient, error)
	GetPatients(params *dto.MedicalPatientReqParams) ([]dto.MedicalPatientResp, error)
}
