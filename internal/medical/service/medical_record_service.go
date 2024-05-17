package service

import (
	"ps-halo-suster/internal/medical/dto"
	medicalRepository "ps-halo-suster/internal/medical/repository"
	userRepository "ps-halo-suster/internal/user/repository"
)

type MedicalRecordService interface {
	CreateRecord(request *dto.MedicalRecordReq) error
	GetRecords(params *dto.MedicalRecordReqParams) ([]dto.MedicalRecordResp, error)
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

func (m *medicalRecordService) CreateRecord(request *dto.MedicalRecordReq) error {
	//TODO implement me
	panic("implement me")
}

func (m *medicalRecordService) GetRecords(params *dto.MedicalRecordReqParams) ([]dto.MedicalRecordResp, error) {
	return m.medicalRecordRepository.GetRecords(params)
}
