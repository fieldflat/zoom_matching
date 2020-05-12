package usecase

import (
	"../domain"
)

type UserRepository interface {
	FindByID(id int) (event domain.Users, err error)
}
