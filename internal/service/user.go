package service

import (
	"Houses/internal/model"
	"Houses/internal/repo"
	"Houses/internal/utils"
)

type UserService interface {
	Login(login model.UserLogin) error
	Register(login model.UserLogin, utype model.UserType) error
	GetUserByEmail(email string) (*model.User, error)
}

type UserServiceImpl struct {
	repo repo.UserRepo
}

func NewUserServiceImpl(repo repo.UserRepo) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

//////////////////////////////////////////////////////////////////////////////////////////////////

func (s *UserServiceImpl) Login(login model.UserLogin) error {
	user, err := s.repo.GetByEmail(login.Email)
	if err != nil {
		return err
	}

	err = utils.ComparePassHash(login.Password, user.Password)
	if err != nil {
		return utils.ErrPasswordNotMatch
	}

	return nil
}

func (s *UserServiceImpl) Register(login model.UserLogin, utype model.UserType) error {
	if err := utils.CheckUserLogin(login); err != nil {
		return err
	}

	hash, err := utils.GenPassHash(login.Password)
	if err != nil {
		return err
	}

	user := model.User{
		UserLogin: model.UserLogin{
			Email:    login.Email,
			Password: hash,
		},
		UserType: utype,
	}

	return s.repo.Create(&user)
}

func (s *UserServiceImpl) GetUserByEmail(email string) (*model.User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
