package model

type User struct {
	ID uint `gorm:"primary_key"`

	UserLogin

	UserType
}

type UserLogin struct {
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}

type UserType uint

const (
	Client UserType = iota
	Moderator
)
