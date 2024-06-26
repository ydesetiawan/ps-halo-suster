package repository

import (
	"ps-halo-suster/internal/user/dto"
	"ps-halo-suster/internal/user/model"
)

type UserRepository interface {
	GetUserByIDAndRole(id string, role string) (model.User, error)
	GetUserByNIPAndRole(nip string, role string) (model.User, error)
	GetUserByNIPAndId(nip string, id string) (model.User, error)
	RegisterUser(user *model.User) (string, error)
	UpdateUser(request *dto.UpdateUserReq) error
	DeleteUser(userId string) error
	GrantAccessUser(request *dto.GrantAccessReq) error
	GetNurses(params *dto.GetNurseParams) ([]dto.GetNurseResp, error)
}
