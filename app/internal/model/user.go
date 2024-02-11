package model

type User struct {
	Id        string
	Status    int
	Email     string
	Username  string
	Password  string
	Firstname string
	Lastname  string
	CreatedAt string
}

type UserDto struct {
	Status    int    `validate:"required"`
	Email     string `validate:"required"`
	Password  string `validate:"required"`
	Username  string `validate:"required"`
	Firstname string `validate:"required"`
	Lastname  string `validate:"required"`
	CreatedAt string `validate:"required"`
}

func (u User) GetUserId() string {
	return u.Id
}

func (u User) GetEmail() string {
	return u.Email
}
