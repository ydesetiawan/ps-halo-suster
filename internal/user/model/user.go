package model

import (
	"time"
)

type User struct {
	ID                  string    `db:"id" json:"id"`
	NIP                 int       `db:"nip" json:"nip"`
	Name                string    `db:"name" json:"name"`
	Password            string    `db:"password" json:"password"`
	Role                string    `db:"role" json:"role"`
	IdentityCardScanImg *string   `db:"identity_card_scan_img" json:"identityCardScanImg"`
	CreatedAt           time.Time `db:"created_at" json:"created_at"`
}

type Role string

const (
	IT    Role = "it"
	NURSE Role = "nurse"
)
