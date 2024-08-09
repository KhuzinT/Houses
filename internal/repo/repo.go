package repo

import "gorm.io/gorm"
import "Houses/internal/db/sqlite"

type Repo struct {
	User  UserRepo
	Flat  FlatRepo
	House HouseRepo
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{
		User:  sqlite.NewUserRepoImpl(db),
		Flat:  sqlite.NewFlatRepoImpl(db),
		House: sqlite.NewHouseRepoImpl(db),
	}
}
