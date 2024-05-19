package repository

import (
	"fmt"
	"ps-halo-suster/internal/medical/record/dto"
	"ps-halo-suster/internal/medical/record/model"
	"ps-halo-suster/pkg/errs"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
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
    INSERT INTO medical_records (id, identity_number, symptoms, medications, created_by, created_at)
    SELECT $2, $1, $3, $4, $5, $6
    FROM check_identity
    WHERE EXISTS (SELECT 1 FROM check_identity);
    `

func (m *medicalRecordRepository) CreateRecord(mRecord *model.MedicalRecord) error {
	result, err := m.db.Exec(queryCreateRecord, mRecord.IdentityNumber, mRecord.ID, mRecord.Symptoms, mRecord.Medications, mRecord.CreatedBy, time.Now())
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

func buildMedicalRecordQuery(params *dto.MedicalRecordReqParams) string {
	var filters []string

	// Add conditions based on the parameters
	if params.IdentityNumber != 0 {
		filters = append(filters, fmt.Sprintf("mr.identity_number = %d", params.IdentityNumber))
	}
	if params.UserID != "" {
		filters = append(filters, fmt.Sprintf("uc.id = '%s'", params.UserID))
	}
	if params.NIP != "" {
		filters = append(filters, fmt.Sprintf("uc.nip = '%s'", params.NIP))
	}

	// Validate createdAt param
	var orderBy string
	if params.CreatedAt == "asc" {
		orderBy = "ORDER BY mr.created_at ASC"
	} else if params.CreatedAt == "desc" || params.CreatedAt == "" {
		orderBy = "ORDER BY mr.created_at DESC"
	}

	// Construct query using CTE
	query := `
		WITH user_check AS (
			SELECT id, nip, name
			FROM users
			WHERE 1 = 1`
	if params.UserID != "" {
		query += fmt.Sprintf(" AND id = '%s'", params.UserID)
	}
	if params.NIP != "" {
		query += fmt.Sprintf(" AND nip = '%s'", params.NIP)
	}
	query += `
		)
		SELECT mr.identity_number, mr.symptoms, mr.medications, mr.created_at, 
		       uc.nip, uc.name AS user_name, uc.id AS user_id,
		       mp.phone_number, mp.name, mp.birth_date, mp.gender, mp.identity_card_scan_img
		FROM medical_records mr
		JOIN user_check uc ON mr.created_by = uc.id
		JOIN medical_patients mp ON mr.identity_number = mp.identity_number`

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

	var records []dto.MedicalRecordResp
	for rows.Next() {
		var record dto.MedicalRecordResp
		var createdAt time.Time
		err := rows.Scan(
			&record.IdentityDetail.IdentityNumber,
			&record.Symptoms,
			&record.Medications,
			&createdAt,
			&record.CreatedBy.NIP,
			&record.CreatedBy.Name,
			&record.CreatedBy.UserID,
			&record.IdentityDetail.PhoneNumber,
			&record.IdentityDetail.Name,
			&record.IdentityDetail.BirthDate,
			&record.IdentityDetail.Gender,
			&record.IdentityDetail.IdentityCardScanImg,
		)
		if err != nil {
			return nil, err
		}
		record.CreatedAt = createdAt.Format(time.RFC3339Nano)
		records = append(records, record)
	}
	if len(records) == 0 {
		records = []dto.MedicalRecordResp{}
	}

	return records, nil
}
