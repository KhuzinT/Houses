package service

import (
	"Houses/internal/model"
	"Houses/internal/repo"
	"time"
)

type HouseService interface {
	CreateHouse(info model.HouseInfo) (*model.House, error)
	GetById(id uint) (*model.House, error)
	AddFlat(houseId uint) error
}

type HouseServiceImpl struct {
	repo repo.HouseRepo
}

func NewHouseServiceImpl(repo repo.HouseRepo) *HouseServiceImpl {
	return &HouseServiceImpl{repo: repo}
}

//////////////////////////////////////////////////////////////////////////////////////////////////

func (s *HouseServiceImpl) CreateHouse(info model.HouseInfo) (*model.House, error) {
	house := model.House{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		HouseInfo: info,
	}

	if err := s.repo.Create(&house); err != nil {
		return nil, err
	}
	return &house, nil
}

func (s *HouseServiceImpl) GetById(id uint) (*model.House, error) {
	house, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return house, nil
}

func (s *HouseServiceImpl) AddFlat(houseId uint) error {
	house, err := s.repo.GetById(houseId)
	if err != nil {
		return err
	}

	// обновляем время последнего добавления жилья
	house.UpdatedAt = time.Now()
	return s.repo.Update(house)
}
