package routers

import (
	"github.com/gin-gonic/gin"
	c "memperbaikikode/controllers"
)

func GetEngine() *gin.Engine {
	router := gin.Default()

	router.POST("/register", c.RegisterUser)
	router.POST("/login", c.LoginUser)
	post := router.Group("/post").Use(c.AuthRequired)
	{
		post.GET("/", c.PostGet)
		post.GET("/:id", c.PostDetail)
		post.POST("/", c.PostCreate)
		post.PUT("/:id", c.PostUpdate)
		post.DELETE("/:id", c.PostDelete)
	}
	router.Use(c.AuthRequired).GET("/logout", c.LogoutUser)

	return router
}
