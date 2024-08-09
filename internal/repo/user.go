package repo

import "Houses/internal/model"

type UserRepo interface {
	Create(user *model.User) error
	GetAll() (*[]model.User, error)
	GetById(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
}
