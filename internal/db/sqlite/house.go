package sqlite

import (
	"Houses/internal/model"
	"Houses/internal/utils"
	"errors"
	"gorm.io/gorm"
)

type HouseRepoImpl struct {
	db *gorm.DB
}

func NewHouseRepoImpl(db *gorm.DB) *HouseRepoImpl {
	return &HouseRepoImpl{db: db}
}

//////////////////////////////////////////////////////////////////////////////////////////////////

func (r *HouseRepoImpl) Create(house *model.House) error {
	if res := r.db.Create(house); res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *HouseRepoImpl) GetAll() (*[]model.House, error) {
	var houses []model.House
	if res := r.db.Find(&houses); res.Error != nil {
		return nil, res.Error
	}
	return &houses, nil
}

func (r *HouseRepoImpl) GetById(id uint) (*model.House, error) {
	var house model.House

	if res := r.db.Where("id = ?", id).First(&house); res.Error == nil {
		return &house, nil
	} else if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, utils.ErrHouseNotFound
	} else {
		return nil, res.Error
	}
}

func (r *HouseRepoImpl) Update(house *model.House) error {
	if res := r.db.Model(house).Updates(house); res.Error != nil {
		return res.Error
	}
	return nil
}
