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

	interactor.StatusCode = 200
	return foundUser, nil
}

// GetAll is a function
// 全てのユーザーを取得する．
func (interactor *UserInteractor) GetAll() (user []domain.Users, err error) {
	// Users の取得
	foundUsers, err := interactor.User.FindAll()
	if err != nil {
		interactor.StatusCode = 404
		return []domain.Users{}, err
	}
	interactor.StatusCode = 200
	return foundUsers, nil
}

// GetByUserID is a function
// idをキーとしてUserを1人取得する．
func (interactor *UserInteractor) GetByUserID(uid string) (user domain.Users, err error) {
	// Users の取得
	foundUser, err := interactor.User.FindByUserID(uid)
	if err != nil {
		interactor.StatusCode = 404
		return domain.Users{}, err
	}

	interactor.StatusCode = 200
	return foundUser, nil
}

// CreateUser is a function
// 全てのユーザーを取得する．
func (interactor *UserInteractor) CreateUser(argUser domain.Users) (user domain.Users, err error) {

	createdUser, err := interactor.User.Create(argUser)
	if err != nil {
		interactor.StatusCode = 404
		return domain.Users{}, err
	}
	interactor.StatusCode = 200
	return createdUser, nil
}

// Update is a function
// 全てのユーザーを取得する．
func (interactor *UserInteractor) Update(argUser domain.Users) (user domain.Users, err error) {

	updatedUser, err := interactor.User.Update(argUser)
	if err != nil {
		interactor.StatusCode = 404
		return domain.Users{}, err
	}
	interactor.StatusCode = 200
	return updatedUser, nil
}

// Delete is a function
// idをキーとしてUserを1人取得する．
func (interactor *UserInteractor) Delete(id int) (user domain.Users, err error) {
	// Users の取得
	foundUser, err := interactor.User.Delete(id)
	if err != nil {
		interactor.StatusCode = 404
		return domain.Users{}, err
	}

	interactor.StatusCode = 200
	return foundUser, nil
}
