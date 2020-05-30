package usecase

import (
	"github.com/fieldflat/zoom_matching/app/domain"
)

// ParticipantInteractor is a type
type ParticipantInteractor struct {
	Participant ParticipantRepository
	StatusCode  int
}

// Get is a function
// idをキーとしてUserを1人取得する．
func (interactor *ParticipantInteractor) Get(id int) (participant domain.Participant, err error) {
	// Participants の取得
	foundParticipant, err := interactor.Participant.FindByID(id)
	if err != nil {
		interactor.StatusCode = 404
		return domain.Participant{}, err
	}

	interactor.StatusCode = 200
	return foundParticipant, nil
}

// GetAll is a function
// 全てのユーザーを取得する．
func (interactor *ParticipantInteractor) GetAll() (participants []domain.Participant, err error) {
	// Participants の取得
	foundParticipants, err := interactor.Participant.FindAll()
	if err != nil {
		interactor.StatusCode = 404
		return []domain.Participant{}, err
	}
	interactor.StatusCode = 200
	return foundParticipants, nil
}

// GetParticipantsByUID is a function
// 全てのユーザーを取得する．
func (interactor *ParticipantInteractor) GetParticipantsByUID(uid string) (participants []domain.Participant, err error) {
	// Participants の取得
	foundParticipants, err := interactor.Participant.FindParticipantsByUID(uid)

	if err != nil {
		interactor.StatusCode = 404
		return []domain.Participant{}, err
	}
	interactor.StatusCode = 200
	return foundParticipants, nil
}

// GetParticipantsByPostID is a function
// 全てのユーザーを取得する．
func (interactor *ParticipantInteractor) GetParticipantsByPostID(post_id int) (participants []domain.Participant, err error) {
	// Participants の取得
	foundParticipants, err := interactor.Participant.FindParticipantsByPostID(post_id)

	if err != nil {
		interactor.StatusCode = 404
		return []domain.Participant{}, err
	}
	interactor.StatusCode = 200
	return foundParticipants, nil
}

// Create is a function
// 全てのユーザーを取得する．
func (interactor *ParticipantInteractor) Create(argUser domain.Participant) (participant domain.Participant, err error) {

	createdUser, err := interactor.Participant.Create(argUser)
	if err != nil {
		interactor.StatusCode = 404
		return domain.Participant{}, err
	}
	interactor.StatusCode = 200
	return createdUser, nil
}

// Update is a function
// 全てのユーザーを取得する．
func (interactor *ParticipantInteractor) Update(argUser domain.Participant) (participant domain.Participant, err error) {

	updatedUser, err := interactor.Participant.Update(argUser)
	if err != nil {
		interactor.StatusCode = 404
		return domain.Participant{}, err
	}
	interactor.StatusCode = 200
	return updatedUser, nil
}

// Delete is a function
// idをキーとしてUserを1人取得する．
func (interactor *ParticipantInteractor) Delete(id int) (participant domain.Participant, err error) {
	// Participants の取得
	foundParticipant, err := interactor.Participant.Delete(id)
	if err != nil {
		interactor.StatusCode = 404
		return domain.Participant{}, err
	}

	interactor.StatusCode = 200
	return foundParticipant, nil
}
