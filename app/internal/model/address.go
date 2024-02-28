package model

import "time"

type Address struct {
	ID        uint `gorm:"primarykey"`
	City      string
	Postcode  string
	Country   string
	Street    string
	Number    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AddressDto struct {
	City     string `json:"city,omitempty"`
	Postcode string `json:"postcode,omitempty"`
	Country  string `json:"country,omitempty"`
	Street   string `json:"street,omitempty"`
	Number   int    `json:"number,omitempty"`
}
