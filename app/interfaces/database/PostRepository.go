package database

import (
	"errors"

	"github.com/fieldflat/zoom_matching/app/domain"
)

type PostRepository struct {
	DB DB
}

// [5]. ここを編集

func (repo *PostRepository) FindByID(id int) (post domain.Post, err error) {
	post = domain.Post{}
	repo.DB.First(&post, id)
	if post.ID <= 0 {
		return domain.Post{}, errors.New("post is not found")
	}
	return post, nil
}

func (repo *PostRepository) FindAll() (post []domain.Post, err error) {
	posts := []domain.Post{}
	repo.DB.Find(&posts)
	return posts, nil
}

func (repo *PostRepository) FindPostsByUID(uid string) (post []domain.Post, err error) {
	posts := []domain.Post{}
	repo.DB.Where("user_id = ?", uid).Find(&posts)
	return posts, nil
}

func (repo *PostRepository) Create(post domain.Post) (event domain.Post, err error) {
	repo.DB.Create(&post)
	if post.ID <= 0 {
		return domain.Post{}, errors.New("post creation is failed")
	}
	return post, nil
}

func (repo *PostRepository) Update(post domain.Post) (event domain.Post, err error) {
	repo.DB.Save(&post)
	if post.ID <= 0 {
		return domain.Post{}, errors.New("post is not found")
	}
	return post, nil
}

func (repo *PostRepository) Delete(id int) (post domain.Post, err error) {
	post = domain.Post{}
	repo.DB.First(&post, id)
	repo.DB.Delete(&post, id)
	if post.ID <= 0 {
		return domain.Post{}, errors.New("post is not found")
	}
	return post, nil
}
