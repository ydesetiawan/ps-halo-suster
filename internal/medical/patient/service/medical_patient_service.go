package service

import (
	"ps-halo-suster/internal/medical/patient/dto"
	"ps-halo-suster/internal/medical/patient/model"
	medicalRepository "ps-halo-suster/internal/medical/patient/repository"
)

type MedicalPatientService interface {
	CreatePatient(request *dto.MedicalPatientReq) error
	GetPatients(params *dto.MedicalPatientReqParams) ([]model.MedicalPatient, error)
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

func (m *medicalPatientService) CreatePatient(request *dto.MedicalPatientReq) error {
	//TODO implement me
	panic("implement me")
}

func (m *medicalPatientService) GetPatients(params *dto.MedicalPatientReqParams) ([]model.MedicalPatient, error) {
	return m.medicalPatientRepository.GetPatients(params)
}
