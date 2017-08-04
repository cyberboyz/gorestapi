package routers

import (
	c "binar-academy/example-db-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func GetEngine() *gin.Engine {
	router := gin.Default()

	router.GET("/user", c.UserGet)
	router.GET("/user/:id", c.UserDetail)
	router.POST("/user", c.UserCreate)
	router.PUT("/user/:id", c.UserUpdate)
	router.DELETE("/user/:id", c.UserDelete)

	return router
}
