package database

import (
	"errors"

	"../../domain"
)

type UserRepository struct {
	DB DB
}

// [5]. ここを編集

func (repo *UserRepository) FindByID(id int) (user domain.Users, err error) {
	user = domain.Users{}
	repo.DB.First(&user, id)
	if user.ID <= 0 {
		return domain.Users{}, errors.New("user is not found")
	}
	return user, nil
}

func (repo *UserRepository) FindAll() (user []domain.Users, err error) {
	users := []domain.Users{}
	repo.DB.Find(&users)
	return users, nil
}

func (repo *UserRepository) Create(user domain.Users) (event domain.Users, err error) {
	repo.DB.Create(&user)
	return user, nil
}

func (repo *UserRepository) Update(user domain.Users) (event domain.Users, err error) {
	repo.DB.Save(&user)
	if user.ID <= 0 {
		return domain.Users{}, errors.New("user is not found")
	}
	return user, nil
}

func (repo *UserRepository) Delete(id int) (user domain.Users, err error) {
	user = domain.Users{}
	repo.DB.First(&user, id)
	repo.DB.Delete(&user, id)
	if user.ID <= 0 {
		return domain.Users{}, errors.New("user is not found")
	}
	return user, nil
}
