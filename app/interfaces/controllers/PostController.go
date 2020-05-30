package controllers

import (
	"strconv"
	"time"

	"github.com/fieldflat/zoom_matching/app/domain"
	"github.com/fieldflat/zoom_matching/app/interfaces/database"
	"github.com/fieldflat/zoom_matching/app/usecase"
)

// PostController is a type
type PostController struct {
	Interactor usecase.PostInteractor
}

// NewPostController is a function
func NewPostController(db database.DB) *PostController {
	return &PostController{
		Interactor: usecase.PostInteractor{
			Post: &database.PostRepository{DB: db},
		},
	}
}

// [2]. methodを作成

// Get is a function
// idをキーとしてPostを1人取得する．
func (controller *PostController) Get(c Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	post, err := controller.Interactor.Get(id)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewH(err.Error(), nil))
		return
	}
	c.JSON(controller.Interactor.StatusCode, NewH("success", post))
}

// GetAll is a function
// 全てのユーザーを取得する．
func (controller *PostController) GetAll(c Context) {

	posts, err := controller.Interactor.GetAll()
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewH(err.Error(), nil))
		return
	}
	c.JSON(controller.Interactor.StatusCode, NewH("success", posts))
}

// GetPostsByUID is a function
// 全てのユーザーを取得する．
func (controller *PostController) GetPostsByUID(c Context) {

	uid := c.Query("uid")

	posts, err := controller.Interactor.GetPostsByUID(uid)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewH(err.Error(), nil))
		return
	}
	c.JSON(controller.Interactor.StatusCode, NewH("success", posts))
}

// Create is a function
// 全てのユーザーを取得する．
func (controller *PostController) Create(c Context) {
	var post domain.Post

	// formデータから構築する場合
	layout := "Mon Jan 2 15:04:05 MST 2006"

	post.UserID = c.PostForm("uid")
	post.ScheduledDate, _ = time.Parse(layout, c.PostForm("scheduled_date"))
	post.DeadlineDate, _ = time.Parse(layout, c.PostForm("deadline_date"))
	post.Title = c.PostForm("title")
	post.Contents = c.PostForm("contents")
	post.Fee, _ = strconv.Atoi(c.PostForm("fee"))

	// JSONデータを用いて構築する場合
	// c.Bind(&post)

	returnPost, err := controller.Interactor.Create(post)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewH(err.Error(), nil))
		return
	}

	c.JSON(controller.Interactor.StatusCode, NewH("success", returnPost))
}

// Update is a function
// idをキーとしてPostを1人取得する．
func (controller *PostController) Update(c Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	post, err := controller.Interactor.Get(id)

	// formデータから構築する場合
	layout := "Mon Jan 2 15:04:05 MST 2006"

	post.UserID = c.PostForm("uid")
	post.ScheduledDate, _ = time.Parse(layout, c.PostForm("scheduled_date"))
	post.DeadlineDate, _ = time.Parse(layout, c.PostForm("deadline_date"))
	post.Title = c.PostForm("title")
	post.Contents = c.PostForm("contents")
	post.Fee, _ = strconv.Atoi(c.PostForm("fee"))
	post.CreatedAt = time.Now()

	// JSONデータを用いて構築する場合
	// c.Bind(&post)

	updatedPost, err := controller.Interactor.Update(post)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewH(err.Error(), nil))
		return
	}
	c.JSON(controller.Interactor.StatusCode, NewH("success", updatedPost))
}

// Delete is a function
// idをキーとしてPostを1人取得する．
func (controller *PostController) Delete(c Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	post, err := controller.Interactor.Delete(id)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewH(err.Error(), nil))
		return
	}
	c.JSON(controller.Interactor.StatusCode, NewH("success", post))
}
