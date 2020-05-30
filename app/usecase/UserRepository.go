package usecase

import (
	"github.com/fieldflat/zoom_matching/app/domain"
)

// [4]. ここを編集
// UserRepository is a type
type UserRepository interface {
	FindByID(id int) (event domain.Users, err error)
	FindByUserID(uid string) (event domain.Users, err error)
	FindAll() (event []domain.Users, err error)
	Create(user domain.Users) (event domain.Users, err error)
	Update(user domain.Users) (event domain.Users, err error)
	Delete(id int) (event domain.Users, err error)
}
