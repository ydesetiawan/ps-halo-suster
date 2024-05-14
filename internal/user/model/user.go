package model

import (
	"github.com/oklog/ulid/v2"
	"math/rand"
	"ps-halo-suster/internal/user/dto"
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

func NewUser(req dto.RegisterReq) *User {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())

	return &User{
		ID:       ulid.MustNew(ms, entropy).String(),
		Name:     req.Name,
		NIP:      req.NIP,
		Password: req.Password,
		Role:     req.Role,
	}
}
