package repository

import (
	"fmt"
	"ps-halo-suster/internal/medical/patient/dto"
	"ps-halo-suster/internal/medical/patient/model"
	"ps-halo-suster/pkg/errs"
	"strings"

	"github.com/jmoiron/sqlx"
)

type medicalPatientRepositoryImpl struct {
	db *sqlx.DB
}

func NewMedicalPatientRepositoryImpl(db *sqlx.DB) MedicalPatientRepository {
	return &medicalPatientRepositoryImpl{db: db}
}

func (r *medicalPatientRepositoryImpl) CreatePatient(medicalPatient model.MedicalPatient) (model.MedicalPatient, error) {
  query := `INSERT INTO medical_patients (identity_number, name, phone_number, birth_date, gender, identity_card_scan_img) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(
    query, 
    medicalPatient.IdentityNumber, 
    medicalPatient.Name, 
    medicalPatient.PhoneNumber, 
    medicalPatient.BirthDate, 
    medicalPatient.Gender, 
    medicalPatient.IdentityCardScanImg)
	
  if err != nil {
    if strings.Contains(err.Error(), "medical_patients_pkey") {
      return medicalPatient, errs.NewErrDataConflict("identity number already exists", medicalPatient.IdentityNumber)
    }

		return medicalPatient, errs.NewErrInternalServerErrors("execute query error [RegisterUser]: ", err.Error())
	}

	return medicalPatient, nil
}

func buildMedicalPatientQuery(params *dto.MedicalPatientReqParams) string {
	var filters []string

	// Add conditions based on the parameters
	if params.IdentityNumber != 0 {
		filters = append(filters, fmt.Sprintf("identity_number = %d", params.IdentityNumber))
	}
	if params.Name != "" {
		filters = append(filters, fmt.Sprintf("LOWER(name) LIKE '%%%s%%'", strings.ToLower(params.Name)))
	}
	if params.PhoneNumber != "" {
		filters = append(filters, fmt.Sprintf("phone_number LIKE '%%%s%%'", params.PhoneNumber))
	}

	// Validate createdAt param
	var orderBy string
	if params.CreatedAt == "asc" {
		orderBy = "ORDER BY created_at ASC"
	} else if params.CreatedAt == "desc" {
		orderBy = "ORDER BY created_at DESC"
	}

	// Construct query
	query := "SELECT * FROM medical_patients"
	if len(filters) > 0 {
		query += " WHERE " + strings.Join(filters, " AND ")
	}
	if orderBy != "" {
		query += " " + orderBy
	}
	if params.Limit == 0 {
		params.Limit = 5
	}
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", params.Limit, params.Offset)

	return query
}

func (m *medicalPatientRepositoryImpl) GetPatients(params *dto.MedicalPatientReqParams) ([]model.MedicalPatient, error) {
	query := buildMedicalPatientQuery(params)
	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//TODO

	return nil, nil

}
