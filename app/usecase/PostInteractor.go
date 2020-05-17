package usecase

import (
	"../domain"
)

// PostInteractor is a type
type PostInteractor struct {
	Post       PostRepository
	StatusCode int
}

// Get is a function
// idをキーとしてUserを1人取得する．
func (interactor *PostInteractor) Get(id int) (post domain.Post, err error) {
	// Posts の取得
	foundPost, err := interactor.Post.FindByID(id)
	if err != nil {
		interactor.StatusCode = 404
		return domain.Post{}, err
	}

	interactor.StatusCode = 200
	return foundPost, nil
}

// GetAll is a function
// 全てのユーザーを取得する．
func (interactor *PostInteractor) GetAll() (posts []domain.Post, err error) {
	// Posts の取得
	foundPosts, err := interactor.Post.FindAll()
	if err != nil {
		interactor.StatusCode = 404
		return []domain.Post{}, err
	}
	interactor.StatusCode = 200
	return foundPosts, nil
}

// Create is a function
// 全てのユーザーを取得する．
func (interactor *PostInteractor) Create(argUser domain.Post) (post domain.Post, err error) {

	createdUser, err := interactor.Post.Create(argUser)
	if err != nil {
		interactor.StatusCode = 404
		return domain.Post{}, err
	}
	interactor.StatusCode = 200
	return createdUser, nil
}

// Update is a function
// 全てのユーザーを取得する．
func (interactor *PostInteractor) Update(argUser domain.Post) (post domain.Post, err error) {

	updatedUser, err := interactor.Post.Update(argUser)
	if err != nil {
		interactor.StatusCode = 404
		return domain.Post{}, err
	}
	interactor.StatusCode = 200
	return updatedUser, nil
}

// Delete is a function
// idをキーとしてUserを1人取得する．
func (interactor *PostInteractor) Delete(id int) (post domain.Post, err error) {
	// Posts の取得
	foundPost, err := interactor.Post.Delete(id)
	if err != nil {
		interactor.StatusCode = 404
		return domain.Post{}, err
	}

	interactor.StatusCode = 200
	return foundPost, nil
}
