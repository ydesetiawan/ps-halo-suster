package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"ps-halo-suster/internal/medical/patient/dto"
	"ps-halo-suster/internal/medical/patient/model"
	"strings"
)

type medicalPatientRepository struct {
	db *sqlx.DB
}

func NewMedicalPatientRepositoryImpl(db *sqlx.DB) MedicalPatientRepository {
	return &medicalPatientRepository{db: db}
}

func (m *medicalPatientRepository) CreatePatient(request *dto.MedicalPatientReq) error {
	//TODO implement me
	panic("implement me")
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

func (m *medicalPatientRepository) GetPatients(params *dto.MedicalPatientReqParams) ([]model.MedicalPatient, error) {
	query := buildMedicalPatientQuery(params)
	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//TODO

	return nil, nil

}
