package repo

import "Houses/internal/model"

type FlatRepo interface {
	Create(flat *model.Flat) error
	GetAll() (*[]model.Flat, error)
	GetById(id uint) (*model.Flat, error)
	GetByHouseId(houseId uint) (*[]model.Flat, error)
	Update(flat *model.Flat) error
}
