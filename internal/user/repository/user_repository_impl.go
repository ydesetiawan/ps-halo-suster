package repository

import (
	"fmt"
	"ps-halo-suster/internal/user/dto"
	"ps-halo-suster/internal/user/model"
	"ps-halo-suster/pkg/errs"
	"strconv"
	"strings"
	"time"

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

const queryDeleteUser = `
    WITH deleted AS (
        DELETE FROM users
        WHERE id = $1
        RETURNING *
    )
    SELECT EXISTS(SELECT 1 FROM deleted)
    `

func (r *userRepositoryImpl) DeleteUser(userId string) error {
	var exists bool
	err := r.db.Get(&exists, queryDeleteUser, userId)
	if err != nil {
		return errs.NewErrDataConflict("execute query error [DeleteUser]: ", err.Error())
	}

	if !exists {
		return errs.NewErrDataNotFound("user not found ", userId, errs.ErrorData{})
	}
	return nil
}

const queryGrantAccessUser = `
    WITH updated AS (
        UPDATE users
        SET password = $1
        WHERE id = $2
        RETURNING *
    )
    SELECT EXISTS(SELECT 1 FROM updated)
    `

func (r *userRepositoryImpl) GrantAccessUser(request *dto.GrantAccessReq) error {
	var exists bool
	err := r.db.Get(&exists, queryGrantAccessUser, request.Password, request.ID)
	if err != nil {
		return errs.NewErrDataConflict("execute query error [GrantAccess]: ", err.Error())
	}

	if !exists {
		return errs.NewErrDataNotFound("user not found ", request.ID, errs.ErrorData{})
	}

	return nil
}

func queryGetNurses(params *dto.GetNurseParams) string {
	var filters []string
	if params.UserId != "" {
		filters = append(filters, fmt.Sprintf("id = '%s'", params.UserId))
	}
	if params.Name != "" {
		filters = append(filters, fmt.Sprintf("LOWER(name) LIKE '%%%s%%'", strings.ToLower(params.Name)))
	}
	if params.NIP != 0 {
		filters = append(filters, fmt.Sprintf("nip LIKE '%%%s%%'", strconv.Itoa(params.NIP)))
	}
	if params.Role != "" {
		if params.Role == "it" {
			filters = append(filters, "nip LIKE '615%'")
		} else if params.Role == "nurse" {
			filters = append(filters, "nip LIKE '303%'")
		}
	}

	// Validate createdAt param
	var orderBy string
	if params.CreatedAt == "asc" {
		orderBy = "ORDER BY created_at ASC"
	} else if params.CreatedAt == "desc" {
		orderBy = "ORDER BY created_at DESC"
	}

	// Construct query
	query := "SELECT id, nip, name, created_at FROM users"
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

func (r *userRepositoryImpl) GetNurses(params *dto.GetNurseParams) ([]dto.GetNurseResp, error) {
	query := queryGetNurses(params)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nurses []dto.GetNurseResp
	for rows.Next() {
		var nurse dto.GetNurseResp
		var nipString string
		err := rows.Scan(
			&nurse.UserId,
			&nipString,
			&nurse.Name,
			&nurse.CreatedAtTime)
		if err != nil {
			return nil, errs.NewErrInternalServerErrors("execute query error [GetNurses]: ", err.Error())
		}
		nurse.CreatedAt = nurse.CreatedAtTime.Format(time.RFC3339)
		nipInt, err := strconv.Atoi(nipString)
		nurse.NIP = nipInt
		nurses = append(nurses, nurse)
	}
	if err := rows.Err(); err != nil {
		return nil, errs.NewErrInternalServerErrors("execute query error [GetNurses]: ", err.Error())
	}

	if len(nurses) == 0 {
		nurses = []dto.GetNurseResp{}
	}

	return nurses, nil
}
