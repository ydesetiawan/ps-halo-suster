package repository

import (
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

func (r *userRepositoryImpl) GetUserByNIP(nip string) (model.User, error) {
	var user model.User
	query := "select * from users where nip = $1 "
	err := r.db.Get(&user, query, nip)
	if err != nil {
		return model.User{}, errs.NewErrInternalServerErrors("execute query error [GetUserByNIP]: ", err.Error())
	}
	return user, err
}

func (r *userRepositoryImpl) GetUserByNIPAndId(nip string, id string) (model.User, error) {
	var user model.User
	query := "select * from users where nip = $1 and id = $2 "
	err := r.db.Get(&user, query, nip, id)
	if err != nil {
		return model.User{}, errs.NewErrInternalServerErrors("execute query error [GetUserByNIP]: ", err.Error())
	}
	return user, err
}

func (r *userRepositoryImpl) RegisterUser(user *model.User) (string, error) {
	var lastInsertId = ""
	query := "insert into users (id, nip, name, password, role) values($1, $2, $3, $4, $5) RETURNING id"

	err := r.db.QueryRowx(query, user.ID, user.NIP, user.Name, user.Password, user.Role).Scan(&lastInsertId)
	if err != nil {
		if strings.Contains(err.Error(), "users_nip_key") {
			return lastInsertId, errs.NewErrDataConflict("nip already exist", user.NIP)
		}
		return lastInsertId, errs.NewErrInternalServerErrors("execute query error [RegisterUser]: ", err.Error())
	}

	return lastInsertId, nil
}
