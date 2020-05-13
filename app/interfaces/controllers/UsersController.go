package controllers

import (
	"strconv"

	"../../usecase"
	"../database"
)

// UsersController is a type
type UsersController struct {
	Interactor usecase.UserInteractor
}

// NewUsersController is a function
func NewUsersController(db database.DB) *UsersController {
	return &UsersController{
		Interactor: usecase.UserInteractor{
			User: &database.UserRepository{DB: db},
		},
	}
}

// Get is a function
// idをキーとしてUserを1人取得する．
func (controller *UsersController) Get(c Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	user, err := controller.Interactor.Get(id)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewH(err.Error(), nil))
		return
	}
	c.JSON(controller.Interactor.StatusCode, NewH("success", user))
}

// GetAll is a function
// 全てのユーザーを取得する．
func (controller *UsersController) GetAll(c Context) {

	users, err := controller.Interactor.GetAll()
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewH(err.Error(), nil))
		return
	}
	c.JSON(controller.Interactor.StatusCode, NewH("success", users))
}
