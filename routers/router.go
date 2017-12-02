package routers

import (
	"github.com/gin-gonic/gin"
	c "memperbaikikode/controllers"
)

func GetEngine() *gin.Engine {
	router := gin.Default()

	router.POST("/register", c.RegisterUser)
	post := router.Group("/post").Use(c.AuthRequired)
	{
		post.GET("/", c.PostGet)
		post.GET("/:id", c.PostDetail)
		post.POST("/", c.PostCreate)
		post.PUT("/:id", c.PostUpdate)
		post.DELETE("/:id", c.PostDelete)
	}

	return router
}
