package usecase

import (
	"../domain"
)

// [4]. ここを編集
// PostRepository is a type
type PostRepository interface {
	FindByID(id int) (event domain.Post, err error)
	FindAll() (event []domain.Post, err error)
	Create(user domain.Post) (event domain.Post, err error)
	Update(user domain.Post) (event domain.Post, err error)
	Delete(id int) (event domain.Post, err error)
	FindPostsByUID(uid string) (event []domain.Post, err error)
}
