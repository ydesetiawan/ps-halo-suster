package repository

import (
	"ps-halo-suster/internal/user/dto"
	"ps-halo-suster/internal/user/model"
	"ps-halo-suster/pkg/errs"
	"strings"

	"github.com/jmoiron/sqlx"
)

type userRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepositoryImpl(db *sqlx.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

const (
	queryGetUserByIDAndRole  = `SELECT * FROM users WHERE id = $1 AND role = $2`
	queryGetUserByNIPAndRole = `SELECT * FROM users WHERE nip = $1 AND role = $2`
	queryGetUserByNIPAndID   = `SELECT * FROM users WHERE nip = $1 AND id = $2`
	queryInsertUser          = `INSERT INTO users (id, nip, name, password, role, identity_card_scan_img) VALUES($1, $2, $3, $4, $5, $6) RETURNING id`
)

func (r *userRepositoryImpl) GetUserByIDAndRole(id string, role string) (model.User, error) {
	var user model.User
	err := r.db.Get(&user, queryGetUserByIDAndRole, id, role)
	if err != nil {
		return model.User{}, errs.NewErrInternalServerErrors("execute query error [GetUserByIDAndRole]: ", err.Error())
	}
	return user, err
}

func (r *userRepositoryImpl) GetUserByNIPAndRole(nip string, role string) (model.User, error) {
	var user model.User
	err := r.db.Get(&user, queryGetUserByNIPAndRole, nip, role)
	if err != nil {
		return model.User{}, errs.NewErrInternalServerErrors("execute query error [GetUserByNIPAndRole]: ", err.Error())
	}
	return user, err
}

func (r *userRepositoryImpl) GetUserByNIPAndId(nip string, id string) (model.User, error) {
	var user model.User
	err := r.db.Get(&user, queryGetUserByNIPAndID, nip, id)
	if err != nil {
		return model.User{}, errs.NewErrInternalServerErrors("execute query error [GetUserByNIPAndId]: ", err.Error())
	}
	return user, err
}

func (r *userRepositoryImpl) RegisterUser(user *model.User) (string, error) {
	var lastInsertId = ""
	err := r.db.QueryRowx(queryInsertUser, user.ID, user.NIP, user.Name, user.Password, user.Role, user.IdentityCardScanImg).Scan(&lastInsertId)
	if err != nil {
		if strings.Contains(err.Error(), "users_nip_key") {
			return lastInsertId, errs.NewErrDataConflict("nip already exist", user.NIP)
		}
		return lastInsertId, errs.NewErrInternalServerErrors("execute query error [RegisterUser]: ", err.Error())
	}

	return lastInsertId, nil
}

const queryUpdateUser = `
    WITH updated AS (
        UPDATE users
        SET name = $1, nip = $2
        WHERE id = $3
        RETURNING *
    )
    SELECT EXISTS(SELECT 1 FROM updated)
    `

func (r *userRepositoryImpl) UpdateUser(request *dto.UpdateUserReq) error {
	var exists bool
	err := r.db.Get(&exists, queryUpdateUser, request.Name, request.NIP, request.ID)
	if err != nil {
		return errs.NewErrDataConflict("execute query error [UpdateUser]: ", err.Error())
	}

	if !exists {
		return errs.NewErrDataNotFound("user not found ", request.ID, errs.ErrorData{})
	}

	return nil
}
