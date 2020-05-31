package infrastructure

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/fieldflat/zoom_matching/app/interfaces/controllers"
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
	store := cookie.NewStore([]byte("secret"))
	r.Gin.Use(sessions.Sessions("login_session", store))

	usersController := controllers.NewUsersController(r.DB)
	postController := controllers.NewPostController(r.DB)
	participantController := controllers.NewParticipantController(r.DB)

	r.Gin.GET("/users/:id", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		usersController.Get(c)
	})

	r.Gin.GET("/users", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		usersController.GetAll(c)
	})

	r.Gin.POST("/users", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		usersController.Create(c)
	})

	r.Gin.PUT("/users/:id", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		usersController.Update(c)
	})

	r.Gin.DELETE("/users/:id", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		usersController.Delete(c)
	})

	r.Gin.GET("/posts/:id", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		postController.Get(c)
	})

	r.Gin.GET("/posts_by_uid", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		postController.GetPostsByUID(c)
	}) // query parameter

	r.Gin.GET("/posts", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		postController.GetAll(c)
	})

	r.Gin.POST("/posts", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		postController.Create(c)
	})

	r.Gin.PUT("/posts/:id", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		postController.Update(c)
	})

	r.Gin.DELETE("/posts/:id", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		postController.Delete(c)
	})

	r.Gin.POST("/participants", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		participantController.Create(c)
	})

	r.Gin.GET("/participants_by_uid", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		participantController.GetParticipantsByUID(c)
	}) // query parameter

	r.Gin.GET("/participants_by_post_id", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		participantController.GetParticipantsByPostID(c)
	}) // query parameter

	r.Gin.DELETE("/participants/:participant_id", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		participantController.Delete(c)
	})

	r.Gin.POST("/login", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		usersController.GetByUserID(c)
		// session := sessions.Default(c)
		// if session.Get("id") == nil {
		// 	id := 1
		// 	session.Set("id", id)
		// 	session.Save()
		// 	c.JSON(200, gin.H{"success": "OK", "id": session.Get("id")})
		// } else {
		// 	c.JSON(200, gin.H{"message": "already login", "id": session.Get("id")})
		// }
	})

	r.Gin.POST("/logout", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		session := sessions.Default(c)
		session.Set("id", nil)
		session.Save()

		c.JSON(200, gin.H{"message": "logout done"})
	})

	r.Gin.POST("/isLogin", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		session := sessions.Default(c)
		if session.Get("id") == nil {
			c.JSON(200, gin.H{"message": "not login"})
		} else {
			c.JSON(200, gin.H{"message": "login!", "id": session.Get("id")})
		}
	})
}

func (r *Routing) Run(port string) {
	r.Gin.Run(port)
}
