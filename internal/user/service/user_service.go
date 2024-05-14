package service

import (
	"ps-halo-suster/internal/user/dto"
	"ps-halo-suster/internal/user/model"
	"ps-halo-suster/internal/user/repository"
	"ps-halo-suster/pkg/bcrypt"
	"ps-halo-suster/pkg/errs"
	"ps-halo-suster/pkg/middleware"
	"strconv"
)

type UserService interface {
	RegisterUser(*dto.RegisterReq) (*dto.RegisterResp, error)
	Login(*dto.LoginReq) (*dto.RegisterResp, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) RegisterUser(req *dto.RegisterReq) (*dto.RegisterResp, error) {
	hashedPassword, _ := bcrypt.HashPassword(req.Password)
	req.Password = hashedPassword
	req.Role = string(model.IT)
	id, err := s.userRepository.RegisterUser(model.NewUser(*req))

	if err != nil {
		return &dto.RegisterResp{}, err
	}
	token, _ := middleware.GenerateJWT(req.Role, id)
	return &dto.RegisterResp{
		UserId:      id,
		NIP:         req.NIP,
		Name:        req.Name,
		AccessToken: token,
	}, nil
}

func (s *userService) Login(req *dto.LoginReq) (*dto.RegisterResp, error) {
	//TODO validation request
	usr, err := s.userRepository.GetUserByNIP(strconv.Itoa(req.NIP))
	if err != nil {
		return &dto.RegisterResp{}, errs.NewErrDataNotFound("user not found ", req.NIP, errs.ErrorData{})
	}
	err = bcrypt.ComparePassword(req.Password, usr.Password)
	if err != nil {
		return &dto.RegisterResp{}, errs.NewErrBadRequest("password is wrong ")
	}

	token, _ := middleware.GenerateJWT(usr.Role, usr.ID)

	return &dto.RegisterResp{
		UserId:      usr.ID,
		NIP:         usr.NIP,
		AccessToken: token,
	}, nil
}
