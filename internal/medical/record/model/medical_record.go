package model

type MedicalRecord struct {
	ID             string `db:"id"`
	IdentityNumber int64  `db:"identity_number"`
	Symptoms       string `db:"symptoms"`
	Medications    string `db:"medications"`
	CreatedBy      string `db:"created_by"`
}
