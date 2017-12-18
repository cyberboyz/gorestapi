package routers

import (
	"github.com/gin-gonic/gin"
	c "memperbaikikode/controllers"
)

func GetEngine() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	router := r.Group("/v1")
	router.POST("/register", c.RegisterUser)
	router.POST("/login", c.LoginUser)
	user := router.Group("/user").Use(c.AuthRequired)
	{
		user.GET("/", c.UserGet)
		user.GET("/:id", c.UserDetail)
		user.PUT("/:id", c.UserUpdate)
		user.DELETE("/:id", c.UserDelete)
	}
	bencana := router.Group("/bencana").Use(c.AuthRequired)
	{
		bencana.POST("/", c.BencanaCreate)
		bencana.GET("/", c.BencanaGet)
		bencana.GET("/:id", c.BencanaDetail)
		bencana.PUT("/:id", c.BencanaUpdate)
		// user.DELETE("/:id", c.BencanaDelete)
	}
	router.Use(c.AuthRequired).GET("/logout", c.LogoutUser)

	return r
}
