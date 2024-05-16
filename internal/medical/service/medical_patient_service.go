package service

import (
	medicalRepository "ps-halo-suster/internal/medical/repository"
)

type MedicalPatientService interface {
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
