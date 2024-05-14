package repository

import "ps-halo-suster/internal/user/model"

type UserRepository interface {
	GetUserByPhoneNumber(phoneNumber string) (model.User, error)
	GetUserByPhoneNumberAndId(phoneNumber string, id int64) (model.User, error)
	RegisterUser(user *model.User) (int64, error)
}
