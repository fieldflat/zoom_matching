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

func (r *Routing) setRouting() {
	usersController := controllers.NewUsersController(r.DB)
	r.Gin.GET("/users/:id", func(c *gin.Context) { usersController.Get(c) })
	r.Gin.GET("/users", func(c *gin.Context) { usersController.GetAll(c) })
}

func (r *Routing) Run(port string) {
	r.Gin.Run(port)
}