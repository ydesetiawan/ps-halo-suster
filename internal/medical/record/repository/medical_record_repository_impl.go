package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"ps-halo-suster/internal/medical/record/dto"
	"ps-halo-suster/internal/medical/record/model"
	"ps-halo-suster/pkg/errs"
	"strings"
)

type medicalRecordRepository struct {
	db *sqlx.DB
}

func NewMedicalRecordRepositoryImpl(db *sqlx.DB) MedicalRecordRepository {
	return &medicalRecordRepository{db: db}
}

const queryCreateRecord = ` WITH check_identity AS (
        SELECT identity_number
        FROM medical_patients
        WHERE identity_number = $1
    )
    INSERT INTO medical_records (id, identity_number, symptoms, medications, created_by)
    SELECT $2, $1, $3, $4, $5
    FROM check_identity
    WHERE EXISTS (SELECT 1 FROM check_identity);
    `

func (m *medicalRecordRepository) CreateRecord(mRecord *model.MedicalRecord) error {
	result, err := m.db.Exec(queryCreateRecord, mRecord.IdentityNumber, mRecord.ID, mRecord.Symptoms, mRecord.Medications, mRecord.CreatedBy)
	if err != nil {

		return errs.NewErrInternalServerErrors("Failed to execute query:", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errs.NewErrInternalServerErrors("Failed to get rows affected:", err)
	}

	if rowsAffected == 0 {
		return errs.NewErrDataNotFound("No record inserted, identity number does not exist in medical_patients.", mRecord.IdentityNumber, errs.ErrorData{})
	}

	return nil
}

// TODO need to testing
func buildMedicalRecordQuery(params *dto.MedicalRecordReqParams) string {
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
	query := buildMedicalRecordQuery(params)
	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//TODO

	return nil, nil
}
