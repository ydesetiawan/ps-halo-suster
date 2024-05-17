package handler

import (
	"github.com/labstack/echo/v4"
	medicalPatientService "ps-halo-suster/internal/medical/patient/service"
	medicalRecordService "ps-halo-suster/internal/medical/record/service"
	userService "ps-halo-suster/internal/user/service"
	"ps-halo-suster/pkg/httphelper/response"
)

type MedicalHandler struct {
	userService           userService.UserService
	medicalPatientService medicalPatientService.MedicalPatientService
	medicalRecordService  medicalRecordService.MedicalRecordService
}

func NewMedicalHandler(userService userService.UserService,
	medicalPatientService medicalPatientService.MedicalPatientService,
	medicalRecordService medicalRecordService.MedicalRecordService) *MedicalHandler {
	return &MedicalHandler{
		userService:           userService,
		medicalPatientService: medicalPatientService,
		medicalRecordService:  medicalRecordService,
	}
}

func (h *MedicalHandler) CreateMedicalPatient(ctx echo.Context) *response.WebResponse {
	return &response.WebResponse{
		Status:  201,
		Message: "CreateMedicalPatient Successfully",
		Data:    nil,
	}
}

func (h *MedicalHandler) GetMedicalPatient(ctx echo.Context) *response.WebResponse {
	return &response.WebResponse{
		Status:  200,
		Message: "GetMedicalPatient Successfully",
		Data:    nil,
	}
}

func (h *MedicalHandler) CreateRecordPatient(ctx echo.Context) *response.WebResponse {
	return &response.WebResponse{
		Status:  201,
		Message: "CreateRecordPatient Successfully",
		Data:    nil,
	}
}

func (h *MedicalHandler) GetRecordPatient(ctx echo.Context) *response.WebResponse {
	return &response.WebResponse{
		Status:  200,
		Message: "GetRecordPatient Successfully",
		Data:    nil,
	}
}
