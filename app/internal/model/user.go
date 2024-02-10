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
	Status    int
	Email     string
	Password  string
	Username  string
	Firstname string
	Lastname  string
	CreatedAt string
}

func (u User) GetUserId() string {
	return u.Id
}

func (u User) GetEmail() string {
	return u.Email
}
