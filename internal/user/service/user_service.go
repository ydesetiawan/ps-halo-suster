package service

import (
	"ps-halo-suster/internal/user/dto"
	"ps-halo-suster/internal/user/model"
	"ps-halo-suster/internal/user/repository"
	"ps-halo-suster/pkg/bcrypt"
	"ps-halo-suster/pkg/errs"
	"ps-halo-suster/pkg/helper"
	"ps-halo-suster/pkg/middleware"
	"strconv"
)

type UserService interface {
	GetUserByIdAndRole(id string, role string) (model.User, error)
	RegisterUser(user *model.User) (*dto.RegisterResp, error)
	Login(*dto.LoginReq) (*dto.RegisterResp, error)
	UpdateNurse(user *dto.UpdateUserReq) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) GetUserByIdAndRole(id string, role string) (model.User, error) {
	user, err := s.userRepository.GetUserByIDAndRole(id, role)
	if err != nil {
		return model.User{}, errs.NewErrUnauthorized("user has not role IT")
	}
	return user, err
}

func (s *userService) RegisterUser(user *model.User) (*dto.RegisterResp, error) {
	if string(model.IT) == user.Role {
		hashedPassword, _ := bcrypt.HashPassword(user.Password)
		user.Password = hashedPassword
	}
	id, err := s.userRepository.RegisterUser(user)

	if err != nil {
		return &dto.RegisterResp{}, err
	}
	token, _ := middleware.GenerateJWT(id, user.Role)
	return &dto.RegisterResp{
		UserId:      id,
		NIP:         user.NIP,
		Name:        user.Name,
		AccessToken: token,
	}, nil
}

func (s *userService) Login(req *dto.LoginReq) (*dto.RegisterResp, error) {
	response := &dto.RegisterResp{}

	nip := strconv.Itoa(req.NIP)
	if model.IT == req.Role && !helper.ValidatePrefixIT(nip) {
		return response, errs.NewErrDataNotFound("user is not from it (nip not starts with 615) ", req.NIP, errs.ErrorData{})
	} else if model.NURSE == req.Role && !helper.ValidatePrefixNurse(nip) {
		return response, errs.NewErrDataNotFound("user is not from nurse (nip not starts with 303) ", req.NIP, errs.ErrorData{})
	}

	user, err := s.userRepository.GetUserByNIPAndRole(nip, string(req.Role))
	if err != nil {
		return response, errs.NewErrDataNotFound("user not found ", req.NIP, errs.ErrorData{})
	}
	err = bcrypt.ComparePassword(req.Password, user.Password)
	if err != nil {
		return response, errs.NewErrBadRequest("password is wrong ")
	}

	token, _ := middleware.GenerateJWT(user.ID, user.Role)

	response = &dto.RegisterResp{
		UserId:      user.ID,
		NIP:         user.NIP,
		Name:        user.Name,
		AccessToken: token,
	}

	return response, nil
}

func (s *userService) UpdateNurse(request *dto.UpdateUserReq) error {
	return s.userRepository.UpdateUser(request)
}
