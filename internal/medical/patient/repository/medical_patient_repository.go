package repository

import (
	"ps-halo-suster/internal/medical/patient/dto"
	"ps-halo-suster/internal/medical/patient/model"
)

type MedicalPatientRepository interface {
	CreatePatient(request *dto.MedicalPatientReq) error
	GetPatients(params *dto.MedicalPatientReqParams) ([]model.MedicalPatient, error)
}
