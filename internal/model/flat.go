package model

type Flat struct {
	ID uint `gorm:"primary_key"`

	HouseID     uint
	ModeratorID uint

	FlatInfo

	FlatStatus
}

type FlatInfo struct {
	Num   uint
	Price uint
	Rooms uint
}

type FlatStatus uint

const (
	Created FlatStatus = iota
	Approved
	Declined
	OnModeration
)
