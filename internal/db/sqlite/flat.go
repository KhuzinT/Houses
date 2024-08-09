package sqlite

import (
	"Houses/internal/model"
	"Houses/internal/utils"
	"errors"
	"gorm.io/gorm"
)

type FlatRepoImpl struct {
	db *gorm.DB
}

func NewFlatRepoImpl(db *gorm.DB) *FlatRepoImpl {
	return &FlatRepoImpl{db: db}
}

//////////////////////////////////////////////////////////////////////////////////////////////////

func (r *FlatRepoImpl) Create(flat *model.Flat) error {
	if res := r.db.Create(flat); res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *FlatRepoImpl) GetAll() (*[]model.Flat, error) {
	var flats []model.Flat
	if res := r.db.Find(&flats); res.Error != nil {
		return nil, res.Error
	}
	return &flats, nil
}

func (r *FlatRepoImpl) GetById(id uint) (*model.Flat, error) {
	var flat model.Flat

	if res := r.db.Where("id = ?", id).First(&flat); res.Error == nil {
		return &flat, nil
	} else if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, utils.ErrFlatNotFound
	} else {
		return nil, res.Error
	}
}

func (r *FlatRepoImpl) GetByHouseId(houseId uint) (*[]model.Flat, error) {
	var flats []model.Flat

	if res := r.db.Where("house_id = ?", houseId).Find(&flats); res.Error == nil {
		return &flats, nil
	} else if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, utils.ErrFlatNotFound
	} else {
		return nil, res.Error
	}
}

func (r *FlatRepoImpl) Update(flat *model.Flat) error {
	if res := r.db.Model(flat).Updates(flat); res.Error != nil {
		return res.Error
	}
	return nil
}
