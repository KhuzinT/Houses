package sqlite

import (
	"Houses/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLite(uri string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.User{}, &model.Flat{}, &model.House{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
