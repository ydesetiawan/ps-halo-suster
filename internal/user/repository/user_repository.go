package repository

import "ps-halo-suster/internal/user/model"

type UserRepository interface {
	GetUserByIDAndRole(id string, role string) (model.User, error)
	GetUserByNIPAndRole(nip string, role string) (model.User, error)
	GetUserByNIPAndId(nip string, id string) (model.User, error)
	RegisterUser(user *model.User) (string, error)
}
