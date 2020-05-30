package database

import (
	"errors"

	"github.com/fieldflat/zoom_matching/app/domain"
)

type UserRepository struct {
	DB DB
}

// [5]. ここを編集

func (repo *UserRepository) FindByID(id int) (user domain.Users, err error) {
	user = domain.Users{}
	repo.DB.First(&user, id)
	if user.ID <= 0 {
		return domain.Users{}, errors.New("user is not found (FindByID)")
	}
	return user, nil
}

func (repo *UserRepository) FindAll() (user []domain.Users, err error) {
	users := []domain.Users{}
	repo.DB.Find(&users)
	return users, nil
}

func (repo *UserRepository) FindByUserID(uid string) (user domain.Users, err error) {
	user = domain.Users{}
	repo.DB.Where("user_id = ?", uid).First(&user)
	if user.ID <= 0 {
		return domain.Users{}, errors.New("user is not found (FindByUserID)")
	}
	return user, nil
}

func (repo *UserRepository) Create(user domain.Users) (event domain.Users, err error) {
	repo.DB.Create(&user)
	if user.ID <= 0 {
		return domain.Users{}, errors.New("user creation is failed")
	}
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
