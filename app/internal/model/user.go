package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint   `gorm:"primarykey"`
	Status    int    `gorm:"default:0"`
	Email     string `gorm:"index:idx_email,unique"`
	Username  string
	Password  string
	Firstname string
	Lastname  string
	AddressID uint
	Address   Address
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) GetUserId() uint {
	return u.ID
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	if u.Password != "" {
		u.Password = ""
	}
	return
}

type UserDto struct {
	Status    int        `json:"status,omitempty"`
	Email     string     `json:"email,omitempty"`
	Password  string     `json:"password,omitempty"`
	Username  string     `json:"username,omitempty"`
	Firstname string     `json:"firstname,omitempty"`
	Lastname  string     `json:"lastname,omitempty"`
	Address   AddressDto `json:"address"`
}
