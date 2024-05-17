package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"ps-halo-suster/internal/medical/record/dto"
	"strings"
)

type medicalRecordRepository struct {
	db *sqlx.DB
}

func NewMedicalRecordRepositoryImpl(db *sqlx.DB) MedicalRecordRepository {
	return &medicalRecordRepository{db: db}
}

func (m *medicalRecordRepository) CreateRecord(request *dto.MedicalRecordReq) error {
	//TODO implement me
	panic("implement me")
}

// TODO need to testing
func queryGetRecords(params *dto.MedicalRecordReqParams) string {
	var filters []string

	// Add conditions based on the parameters
	if params.IdentityDetail.IdentityNumber != 0 {
		filters = append(filters, fmt.Sprintf("identity_number = %d", params.IdentityDetail.IdentityNumber))
	}
	if params.CreatedBy.UserID != "" {
		filters = append(filters, fmt.Sprintf("users.id = '%s'", params.CreatedBy.UserID))
	}
	if params.CreatedBy.NIP != "" {
		filters = append(filters, fmt.Sprintf("users.nip = '%s'", params.CreatedBy.NIP))
	}

	// Validate createdAt param
	var orderBy string
	if params.CreatedAt == "asc" {
		orderBy = "ORDER BY medical_records.created_at ASC"
	} else if params.CreatedAt == "desc" {
		orderBy = "ORDER BY medical_records.created_at DESC"
	}

	// Construct query
	query := "SELECT medical_records.id, medical_records.identity_number, medical_records.symptoms, medical_records.medications, medical_records.created_at FROM medical_records"
	query += " JOIN users ON medical_records.created_by = users.id"
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

func (m *medicalRecordRepository) GetRecords(params *dto.MedicalRecordReqParams) ([]dto.MedicalRecordResp, error) {
	query := queryGetRecords(params)
	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//TODO

	return nil, nil
}
