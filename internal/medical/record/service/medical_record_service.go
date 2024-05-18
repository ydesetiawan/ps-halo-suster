package service

import (
	patientRepository "ps-halo-suster/internal/medical/patient/repository"
	"ps-halo-suster/internal/medical/record/dto"
	"ps-halo-suster/internal/medical/record/model"
	"ps-halo-suster/internal/medical/record/repository"
	userRepository "ps-halo-suster/internal/user/repository"
	"ps-halo-suster/pkg/helper"
)

type MedicalRecordService interface {
	CreateRecord(request *dto.MedicalRecordReq) error
	GetRecords(params *dto.MedicalRecordReqParams) ([]dto.MedicalRecordResp, error)
}

type medicalRecordService struct {
	userRepository           userRepository.UserRepository
	medicalPatientRepository patientRepository.MedicalPatientRepository
	medicalRecordRepository  repository.MedicalRecordRepository
}

func NewMedicalRecordServiceImpl(userRepository userRepository.UserRepository,
	medicalPatientRepository patientRepository.MedicalPatientRepository,
	medicalRecordRepository repository.MedicalRecordRepository) MedicalRecordService {
	return &medicalRecordService{
		userRepository:           userRepository,
		medicalPatientRepository: medicalPatientRepository,
		medicalRecordRepository:  medicalRecordRepository,
	}
}

func (m *medicalRecordService) CreateRecord(request *dto.MedicalRecordReq) error {

	mRecord := &model.MedicalRecord{
		ID:             helper.GenerateULID(),
		IdentityNumber: request.IdentityNumber,
		Symptoms:       request.Symptoms,
		Medications:    request.Medications,
		CreatedBy:      request.UserId,
	}

	return m.medicalRecordRepository.CreateRecord(mRecord)
}

func (m *medicalRecordService) GetRecords(params *dto.MedicalRecordReqParams) ([]dto.MedicalRecordResp, error) {
	return m.medicalRecordRepository.GetRecords(params)
}
