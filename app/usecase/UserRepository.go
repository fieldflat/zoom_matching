package usecase

import (
	"../domain"
)

// UserRepository is a type
type UserRepository interface {
	FindByID(id int) (event domain.Users, err error)
	FindAll() (event []domain.Users, err error)
}
