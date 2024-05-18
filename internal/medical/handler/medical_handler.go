package handler

import (
	patientDto "ps-halo-suster/internal/medical/patient/dto"
	medicalPatientService "ps-halo-suster/internal/medical/patient/service"
	recordDto "ps-halo-suster/internal/medical/record/dto"
	medicalRecordService "ps-halo-suster/internal/medical/record/service"
	"ps-halo-suster/pkg/base/handler"
	"ps-halo-suster/pkg/helper"
	"ps-halo-suster/pkg/httphelper/response"

	"github.com/labstack/echo/v4"
)

type MedicalHandler struct {
	medicalPatientService medicalPatientService.MedicalPatientService
	medicalRecordService  medicalRecordService.MedicalRecordService
}

func NewMedicalHandler(
	medicalPatientService medicalPatientService.MedicalPatientService,
	medicalRecordService medicalRecordService.MedicalRecordService) *MedicalHandler {
	return &MedicalHandler{
		medicalPatientService: medicalPatientService,
		medicalRecordService:  medicalRecordService,
	}
}

func (h *MedicalHandler) CreateMedicalPatient(ctx echo.Context) *response.WebResponse {
	var request = new(patientDto.MedicalPatientReq)
	err := ctx.Bind(&request)

	err = patientDto.ValidateMedicalPatientReq(request)
	helper.Panic400IfError(err)

	medicalPatient, err := h.medicalPatientService.CreatePatient(request)

	helper.PanicIfError(err, "CreateMedicalPatient failed")

	return &response.WebResponse{
		Status:  201,
		Message: "CreateMedicalPatient Successfully",
		Data:    medicalPatient,
	}
}

func (h *MedicalHandler) GetMedicalPatient(ctx echo.Context) *response.WebResponse {
	var params = new(patientDto.MedicalPatientReqParams)
	err := ctx.Bind(params)

	results, err := h.medicalPatientService.GetPatients(params)
	helper.PanicIfError(err, "failed to GetMedicalPatient")

	return &response.WebResponse{
		Status:  200,
		Message: "success",
		Data:    results,
	}
}

func (h *MedicalHandler) CreateRecordPatient(ctx echo.Context) *response.WebResponse {
	var request = new(recordDto.MedicalRecordReq)
	err := ctx.Bind(&request)

	err = recordDto.ValidateMedicalRecordReq(request)
	helper.Panic400IfError(err)

	userId, err := handler.GetUserId(ctx)
	helper.PanicIfError(err, "user unauthorized")

	request.UserId = userId
	err = h.medicalRecordService.CreateRecord(request)
	helper.PanicIfError(err, "CreateRecordPatient failed")

	return &response.WebResponse{
		Status:  201,
		Message: "CreateRecordPatient Successfully",
		Data:    nil,
	}
}

func (h *MedicalHandler) GetRecordPatient(ctx echo.Context) *response.WebResponse {
	var params = new(recordDto.MedicalRecordReqParams)
	err := ctx.Bind(params)

	results, err := h.medicalRecordService.GetRecords(params)
	helper.PanicIfError(err, "failed to GetRecordPatient")

	return &response.WebResponse{
		Status:  200,
		Message: "GetRecordPatient Successfully",
		Data:    results,
	}
}
