package models

import "time"

type User struct {
	Name string
	Login string
	Email string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDeleted bool
}

type UserRepository interface {
	Create (user *User) error
	Select (query string) (*User, error)
	Update (user *User, query string) (*User, error)
	Delete (user *User) error
}
