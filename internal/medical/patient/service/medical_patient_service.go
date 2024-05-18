package service

import (
	"ps-halo-suster/internal/medical/patient/dto"
	"ps-halo-suster/internal/medical/patient/model"
	medicalRepository "ps-halo-suster/internal/medical/patient/repository"
)

type MedicalPatientService interface {
	CreatePatient(request *dto.MedicalPatientReq) (model.MedicalPatient, error)
	GetPatients(params *dto.MedicalPatientReqParams) ([]dto.MedicalPatientResp, error)
}

type medicalPatientService struct {
	medicalPatientRepository medicalRepository.MedicalPatientRepository
}

func NewMedicalPatientServiceImpl(
	medicalPatientRepository medicalRepository.MedicalPatientRepository) MedicalPatientService {
	return &medicalPatientService{
		medicalPatientRepository: medicalPatientRepository,
	}
}

func (m *medicalPatientService) CreatePatient(request *dto.MedicalPatientReq) (model.MedicalPatient, error) {
  medicalPatientReq := dto.NewMedicalPatient(*request)

	medicalPatient, err := m.medicalPatientRepository.CreatePatient(*medicalPatientReq)

  if err != nil {
    return medicalPatient, err;
  }

  return medicalPatient, nil
}

func (m *medicalPatientService) GetPatients(params *dto.MedicalPatientReqParams) ([]dto.MedicalPatientResp, error) {
	return m.medicalPatientRepository.GetPatients(params)
}
