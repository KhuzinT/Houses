package repo

import "Houses/internal/model"

type HouseRepo interface {
	Create(house *model.House) error
	GetAll() (*[]model.House, error)
	GetById(id uint) (*model.House, error)
	Update(house *model.House) error
}
