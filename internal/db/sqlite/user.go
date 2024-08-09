package sqlite

import (
	"Houses/internal/model"
	"Houses/internal/utils"
	"errors"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepoImpl(db *gorm.DB) *UserRepoImpl {
	return &UserRepoImpl{db: db}
}

//////////////////////////////////////////////////////////////////////////////////////////////////

func (r *UserRepoImpl) Create(user *model.User) error {
	if res := r.db.Create(user); res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *UserRepoImpl) GetAll() (*[]model.User, error) {
	var users []model.User
	if res := r.db.Find(&users); res.Error != nil {
		return nil, res.Error
	}
	return &users, nil
}

func (r *UserRepoImpl) GetById(id uint) (*model.User, error) {
	var user model.User

	if res := r.db.Where("id = ?", id).First(&user); res.Error == nil {
		return &user, nil
	} else if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, utils.ErrUserNotFound
	} else {
		return nil, res.Error
	}
}

func (r *UserRepoImpl) GetByEmail(email string) (*model.User, error) {
	var user model.User

	if res := r.db.Where("email = ?", email).First(&user); res.Error == nil {
		return &user, nil
	} else if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, utils.ErrUserNotFound
	} else {
		return nil, res.Error
	}
}
