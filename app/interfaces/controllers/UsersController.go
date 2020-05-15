package controllers

import (
	"strconv"
	"time"

	"../../domain"
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

// [2]. methodを作成

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

// Create is a function
// 全てのユーザーを取得する．
func (controller *UsersController) Create(c Context) {
	var user domain.Users

	// formデータから構築する場合
	user.ID, _ = strconv.Atoi(c.PostForm("ID"))
	user.UserID = c.PostForm("uid")
	user.FirstName = c.PostForm("first_name")
	user.LastName = c.PostForm("last_name")
	user.DisplayName = c.PostForm("display_name")
	user.Email = c.PostForm("email")
	user.CreatedAt = time.Now()

	// JSONデータを用いて構築する場合
	// c.Bind(&user)

	returnUser, err := controller.Interactor.CreateUser(user)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewH(err.Error(), nil))
		return
	}

	c.JSON(controller.Interactor.StatusCode, NewH("success", returnUser))
}

// Delete is a function
// idをキーとしてUserを1人取得する．
func (controller *UsersController) Delete(c Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	user, err := controller.Interactor.Delete(id)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewH(err.Error(), nil))
		return
	}
	c.JSON(controller.Interactor.StatusCode, NewH("success", user))
}
