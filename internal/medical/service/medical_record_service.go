package service

import (
	medicalRepository "ps-halo-suster/internal/medical/repository"
	userRepository "ps-halo-suster/internal/user/repository"
)

type MedicalRecordService interface {
}

type medicalRecordService struct {
	userRepository           userRepository.UserRepository
	medicalPatientRepository medicalRepository.MedicalPatientRepository
	medicalRecordRepository  medicalRepository.MedicalRecordRepository
}

func NewMedicalRecordServiceImpl(userRepository userRepository.UserRepository,
	medicalPatientRepository medicalRepository.MedicalPatientRepository,
	medicalRecordRepository medicalRepository.MedicalRecordRepository) MedicalRecordService {
	return &medicalRecordService{
		userRepository:           userRepository,
		medicalPatientRepository: medicalPatientRepository,
		medicalRecordRepository:  medicalRecordRepository,
	}
}
