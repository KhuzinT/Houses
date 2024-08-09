package service

import (
	"Houses/internal/model"
	"Houses/internal/repo"
)

type FlatService interface {
	GetById(id uint) (*model.Flat, error)
	GetFlats(houseId uint) (*[]model.Flat, error)
	GetApprovedFlats(houseId uint) (*[]model.Flat, error)
	CreateFlat(houseId uint, info model.FlatInfo) (*model.Flat, error)
	UpdateFlatStatus(id uint, moderatorId uint, status model.FlatStatus) error
}

type FlatServiceImpl struct {
	repo         repo.FlatRepo
	houseService HouseService
}

func NewFlatServiceImpl(repo repo.FlatRepo, service HouseService) *FlatServiceImpl {
	return &FlatServiceImpl{repo: repo, houseService: service}
}

//////////////////////////////////////////////////////////////////////////////////////////////////

func (s *FlatServiceImpl) GetById(id uint) (*model.Flat, error) {
	flat, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return flat, nil
}

func (s *FlatServiceImpl) GetFlats(houseId uint) (*[]model.Flat, error) {
	flats, err := s.repo.GetByHouseId(houseId)
	if err != nil {
		return nil, err
	}

	return flats, nil
}

func (s *FlatServiceImpl) GetApprovedFlats(houseId uint) (*[]model.Flat, error) {
	flats, err := s.GetFlats(houseId)
	if err != nil {
		return nil, err
	}

	var filtered []model.Flat
	for _, flat := range *flats {
		if flat.FlatStatus == model.Approved {
			filtered = append(filtered, flat)
		}
	}
	return &filtered, nil
}

func (s *FlatServiceImpl) CreateFlat(houseId uint, info model.FlatInfo) (*model.Flat, error) {
	_, err := s.houseService.GetById(houseId)
	if err != nil {
		return nil, err
	}

	err = s.houseService.AddFlat(houseId)
	if err != nil {
		return nil, err
	}

	flat := model.Flat{
		HouseID:    houseId,
		FlatInfo:   info,
		FlatStatus: model.Created,
	}
	if err := s.repo.Create(&flat); err != nil {
		return nil, err
	}

	return &flat, nil
}

func (s *FlatServiceImpl) UpdateFlatStatus(id uint, moderatorId uint, status model.FlatStatus) error {
	flat, err := s.repo.GetById(id)
	if err != nil {
		return err
	}

	flat.FlatStatus = status
	if status == model.OnModeration {
		flat.ModeratorID = moderatorId
	} else {
		flat.ModeratorID = 0
	}
	return s.repo.Update(flat)
}
