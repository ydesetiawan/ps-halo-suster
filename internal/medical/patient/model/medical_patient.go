package model

type MedicalPatient struct {
  IdentityNumber      int    `db:"identity_number"`
  Name                string `db:"name"`
  PhoneNumber         string `db:"phone_number"`
  BirthDate           string `db:"birth_date"`
  Gender              string `db:"gender"`
  IdentityCardScanImg string `db:"identity_card_scan_img"`
}
