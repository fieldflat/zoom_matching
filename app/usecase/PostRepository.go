package usecase

import (
	"github.com/fieldflat/zoom_matching/app/domain"
)

// [4]. ここを編集
// PostRepository is a type
type PostRepository interface {
	FindByID(id int) (event domain.Post, err error)
	FindAll() (event []domain.Post, err error)
	Create(post domain.Post) (event domain.Post, err error)
	Update(post domain.Post) (event domain.Post, err error)
	Delete(id int) (event domain.Post, err error)
	FindPostsByUID(uid string) (event []domain.Post, err error)
}
