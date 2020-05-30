package controllers

import (
	"log"
	"strconv"

	"github.com/fieldflat/zoom_matching/app/domain"
	"github.com/fieldflat/zoom_matching/app/interfaces/database"
	"github.com/fieldflat/zoom_matching/app/usecase"
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

// GetByUserID is a function
// idをキーとしてUserを1人取得する．
func (controller *UsersController) GetByUserID(c Context) {

	uid := c.PostForm("uid")

	user, err := controller.Interactor.GetByUserID(uid)
	log.Print(user.Password)
	log.Print(c.PostForm("password"))
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewH(err.Error(), nil))
		return
	} else if user.Password != c.PostForm("password") {
		c.JSON(404, NewH("password incorrect", nil))
		return
	}
	c.JSON(controller.Interactor.StatusCode, NewH("success", user))
}

// Create is a function
// 全てのユーザーを取得する．
func (controller *UsersController) Create(c Context) {
	var user domain.Users

	// formデータから構築する場合
	user.UserID = c.PostForm("uid")
	user.FirstName = c.PostForm("first_name")
	user.LastName = c.PostForm("last_name")
	user.DisplayName = c.PostForm("display_name")
	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")

	// JSONデータを用いて構築する場合
	// c.Bind(&user)

	returnUser, err := controller.Interactor.CreateUser(user)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewH(err.Error(), nil))
		return
	}

	c.JSON(controller.Interactor.StatusCode, NewH("success", returnUser))
}

// Update is a function
// idをキーとしてUserを1人取得する．
func (controller *UsersController) Update(c Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.Get(id)

	// formデータから構築する場合
	user.UserID = c.PostForm("uid")
	user.FirstName = c.PostForm("first_name")
	user.LastName = c.PostForm("last_name")
	user.DisplayName = c.PostForm("display_name")
	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")

	// JSONデータを用いて構築する場合
	// c.Bind(&user)

	updatedUser, err := controller.Interactor.Update(user)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewH(err.Error(), nil))
		return
	}
	c.JSON(controller.Interactor.StatusCode, NewH("success", updatedUser))
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
