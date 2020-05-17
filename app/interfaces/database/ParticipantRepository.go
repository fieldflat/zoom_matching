package database

import (
	"errors"

	"../../domain"
)

type ParticipantRepository struct {
	DB DB
}

// [5]. ここを編集

func (repo *ParticipantRepository) FindByID(id int) (participant domain.Participant, err error) {
	participant = domain.Participant{}
	repo.DB.First(&participant, id)
	if participant.ID <= 0 {
		return domain.Participant{}, errors.New("participant is not found")
	}
	return participant, nil
}

func (repo *ParticipantRepository) FindAll() (participant []domain.Participant, err error) {
	participants := []domain.Participant{}
	repo.DB.Find(&participants)
	return participants, nil
}

func (repo *ParticipantRepository) FindParticipantsByUID(uid string) (participant []domain.Participant, err error) {
	participants := []domain.Participant{}
	repo.DB.Where("user_id = ?", uid).Find(&participants)
	return participants, nil
}

func (repo *ParticipantRepository) FindParticipantsByPostID(post_id int) (participant []domain.Participant, err error) {
	participants := []domain.Participant{}
	repo.DB.Where("post_id = ?", post_id).Find(&participants)
	return participants, nil
}

func (repo *ParticipantRepository) Create(participant domain.Participant) (event domain.Participant, err error) {
	repo.DB.Create(&participant)
	if participant.ID <= 0 {
		return domain.Participant{}, errors.New("participant creation is failed")
	}
	return participant, nil
}

func (repo *ParticipantRepository) Update(participant domain.Participant) (event domain.Participant, err error) {
	repo.DB.Save(&participant)
	if participant.ID <= 0 {
		return domain.Participant{}, errors.New("participant is not found")
	}
	return participant, nil
}

func (repo *ParticipantRepository) Delete(id int) (participant domain.Participant, err error) {
	participant = domain.Participant{}
	repo.DB.First(&participant, id)
	repo.DB.Delete(&participant, id)
	if participant.ID <= 0 {
		return domain.Participant{}, errors.New("participant is not found")
	}
	return participant, nil
}
