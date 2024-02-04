package model

type User struct {
	Id        string
	Status    int
	Email     string
	Username  string
	Firstname string
	Lastname  string
	CreatedAt string
}

type UserDto struct {
	Status    int
	Email     string
	Username  string
	Firstname string
	Lastname  string
	CreatedAt string
}
