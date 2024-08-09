package model

import "time"

type House struct {
	ID uint `gorm:"primary_key"`

	CreatedAt time.Time
	UpdatedAt time.Time

	HouseInfo
}

type HouseInfo struct {
	Num       uint
	Year      uint
	Addr      string
	Developer string
}
