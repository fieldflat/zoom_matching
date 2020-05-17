package usecase

import (
	"../domain"
)

// [4]. ここを編集
// ParticipantRepository is a type
type ParticipantRepository interface {
	FindByID(id int) (event domain.Participant, err error)
	FindAll() (event []domain.Participant, err error)
	Create(post domain.Participant) (event domain.Participant, err error)
	Update(post domain.Participant) (event domain.Participant, err error)
	Delete(id int) (event domain.Participant, err error)
	FindParticipantsByUID(uid string) (event []domain.Participant, err error)
	FindParticipantsByPostID(post_id int) (event []domain.Participant, err error)
}
