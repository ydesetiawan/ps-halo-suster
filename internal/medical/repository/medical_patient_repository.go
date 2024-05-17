package repository

import (
	"ps-halo-suster/internal/medical/dto"
	"ps-halo-suster/internal/medical/model"
)

type MedicalPatientRepository interface {
	CreatePatient(request *dto.MedicalPatientReq) error
	GetPatients(params *dto.MedicalPatientReqParams) ([]model.MedicalPatient, error)
}
