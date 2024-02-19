package model

type User struct {
	ID        uint `gorm:"primarykey"`
	Status    int
	Email     string
	Username  string
	Password  string
	Firstname string
	Lastname  string
}

type UserDto struct {
	Status    int    `validate:"required"`
	Email     string `validate:"required"`
	Password  string `validate:"required"`
	Username  string `validate:"required"`
	Firstname string `validate:"required"`
	Lastname  string `validate:"required"`
}

func (u User) GetUserId() uint {
	return u.ID
}

func (u User) GetEmail() string {
	return u.Email
}
