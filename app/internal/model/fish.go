package model

import "time"

type Fish struct {
	ID        uint `gorm:"primarykey"`
	Status    int  `gorm:"default:0"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FishDto struct {
	Name      string
	CreatedAt string
}
