package controllers

import (
	"strconv"

	"github.com/fieldflat/zoom_matching/app/domain"
	"github.com/fieldflat/zoom_matching/app/interfaces/database"
	"github.com/fieldflat/zoom_matching/app/usecase"
)

// ParticipantController is a type
type ParticipantController struct {
	Interactor usecase.ParticipantInteractor
}

// NewParticipantController is a function
func NewParticipantController(db database.DB) *ParticipantController {
	return &ParticipantController{
		Interactor: usecase.ParticipantInteractor{
			Participant: &database.ParticipantRepository{DB: db},
		},
	}
}

// [2]. methodを作成

// GetParticipantsByUID is a function
// 全てのユーザーを取得する．
func (controller *ParticipantController) GetParticipantsByUID(c Context) {

	uid := c.Query("uid")

	participants, err := controller.Interactor.GetParticipantsByUID(uid)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewH(err.Error(), nil))
		return
	}
	c.JSON(controller.Interactor.StatusCode, NewH("success", participants))
}

// GetParticipantsByPostID is a function
// 全てのユーザーを取得する．
func (controller *ParticipantController) GetParticipantsByPostID(c Context) {

	postID, _ := strconv.Atoi(c.Query("post_id"))

	participants, err := controller.Interactor.GetParticipantsByPostID(postID)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewH(err.Error(), nil))
		return
	}
	c.JSON(controller.Interactor.StatusCode, NewH("success", participants))
}

// Create is a function
// 全てのユーザーを取得する．
func (controller *ParticipantController) Create(c Context) {
	var participant domain.Participant

	participant.UserID = c.PostForm("uid")
	participant.PostID, _ = strconv.Atoi(c.PostForm("post_id"))

	// JSONデータを用いて構築する場合
	// c.Bind(&participant)

	returnParticipant, err := controller.Interactor.Create(participant)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewH(err.Error(), nil))
		return
	}

	c.JSON(controller.Interactor.StatusCode, NewH("success", returnParticipant))
}

// Delete is a function
// idをキーとしてParticipantを1人取得する．
func (controller *ParticipantController) Delete(c Context) {

	participant_id, _ := strconv.Atoi(c.Param("participant_id"))

	participant, err := controller.Interactor.Delete(participant_id)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewH(err.Error(), nil))
		return
	}
	c.JSON(controller.Interactor.StatusCode, NewH("success", participant))
}
