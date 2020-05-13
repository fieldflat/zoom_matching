package usecase

import (
	"../domain"
)

// UserInteractor is a type
type UserInteractor struct {
	User       UserRepository
	StatusCode int
}

// Get is a function
// idをキーとしてUserを1人取得する．
func (interactor *UserInteractor) Get(id int) (user domain.Users, err error) {
	// Users の取得
	foundUser, err := interactor.User.FindByID(id)
	if err != nil {
		interactor.StatusCode = 404
		return domain.Users{}, err
	}
	// user = foundUser.BuildForGet()
	interactor.StatusCode = 200
	return foundUser, nil
}

// GetAll is a function
// 全てのユーザーを取得する．
func (interactor *UserInteractor) GetAll() (user []domain.Users, err error) {
	// Users の取得
	foundUser, err := interactor.User.FindAll()
	if err != nil {
		interactor.StatusCode = 404
		return []domain.Users{}, err
	}
	// user = foundUser.BuildForGet()
	interactor.StatusCode = 200
	return foundUser, nil
}
