package repository

import "github.com/jmoiron/sqlx"

type medicalPatientRepository struct {
	db *sqlx.DB
}

func NewMedicalPatientRepositoryImpl(db *sqlx.DB) MedicalPatientRepository {
	return &medicalPatientRepository{db: db}
}
