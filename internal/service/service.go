package service

import (
	"Houses/internal/repo"
)

type Service struct {
	User  UserService
	Flat  FlatService
	House HouseService
}

func NewService(repo *repo.Repo) *Service {
	houseService := NewHouseServiceImpl(repo.House)
	return &Service{
		User:  NewUserServiceImpl(repo.User),
		Flat:  NewFlatServiceImpl(repo.Flat, houseService),
		House: houseService,
	}
}
