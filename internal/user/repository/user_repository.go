package repository

import "ps-halo-suster/internal/user/model"

type UserRepository interface {
	GetUserByNIP(nip string) (model.User, error)
	GetUserByNIPAndId(nip string, id string) (model.User, error)
	RegisterUser(user *model.User) (string, error)
}
