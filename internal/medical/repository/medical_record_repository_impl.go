package repository

import "github.com/jmoiron/sqlx"

type medicalRecordRepository struct {
	db *sqlx.DB
}

func NewMedicalRecordRepositoryImpl(db *sqlx.DB) MedicalRecordRepository {
	return &medicalRecordRepository{db: db}
}
