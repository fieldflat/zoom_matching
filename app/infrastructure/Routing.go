package infrastructure

import (
	"github.com/gin-gonic/gin"

	"../interfaces/controllers"
)

type Routing struct {
	DB  *DB
	Gin *gin.Engine
}

func NewRouting(db *DB) *Routing {
	r := &Routing{
		DB:  db,
		Gin: gin.Default(),
	}
	r.setRouting()
	return r
}

// [1]. Routingの設定
func (r *Routing) setRouting() {
	usersController := controllers.NewUsersController(r.DB)
	postController := controllers.NewPostController(r.DB)
	r.Gin.GET("/users/:id", func(c *gin.Context) { usersController.Get(c) })
	r.Gin.GET("/users", func(c *gin.Context) { usersController.GetAll(c) })
	r.Gin.POST("/users", func(c *gin.Context) { usersController.Create(c) })
	r.Gin.PUT("/users/:id", func(c *gin.Context) { usersController.Update(c) })
	r.Gin.DELETE("/users/:id", func(c *gin.Context) { usersController.Delete(c) })

	r.Gin.GET("/posts/:id", func(c *gin.Context) { postController.Get(c) })
	r.Gin.GET("/posts_by_uid", func(c *gin.Context) { postController.GetPostsByUID(c) })
	r.Gin.GET("/posts", func(c *gin.Context) { postController.GetAll(c) })
	r.Gin.POST("/posts", func(c *gin.Context) { postController.Create(c) })
	r.Gin.PUT("/posts/:id", func(c *gin.Context) { postController.Update(c) })
	r.Gin.DELETE("/posts/:id", func(c *gin.Context) { postController.Delete(c) })
}

func (r *Routing) Run(port string) {
	r.Gin.Run(port)
}
