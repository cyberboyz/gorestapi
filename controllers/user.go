package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	m "memperbaikikode/models"
	"net/http"
	"strconv"
	"strings"
)

func AuthRequired(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")

	auth := &m.User{}
	if authorization == "" {
		response := &m.Response{
			Message: "Cannot access the resource : You need to authenticate",
		}
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}

	err := m.AuthorizeUser(auth, authorization)
	if err != nil {
		response := &m.Response{
			Message: "Unauthorized access",
		}
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}
}

func RegisterUser(c *gin.Context) {
	register := &m.User{}

	err := c.Bind(register)
	if err != nil {
		response := &m.Response{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	if register.Email == "" {
		response := &m.Response{
			Message: "Error : Email cannot be empty",
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	validation := ValidateFormatEmail(register.Email)

	if validation != "" {
		response := &m.Response{
			Message: validation,
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	register.Email = strings.ToLower(register.Email)

	if register.Password == "" {
		response := &m.Response{
			Message: "Error : Password cannot be empty",
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	register.Password = string(hash)

	if register.Name == "" {
		response := &m.Response{
			Message: "Error : Name cannot be empty",
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	register.Token = randToken(20)

	user, err := m.RegisterUser(register)
	if err != nil {
		response := &m.Response{
			Message: err.Error(),
		}
		c.JSON(http.StatusServiceUnavailable, response)
		c.Abort()
		return
	}

	response := &m.Response{
		Message:    "Registration is successful",
		Success:    true,
		StatusCode: http.StatusOK,
		Data: map[string]interface{}{
			"nama":   user.Name,
			"email":  user.Email,
			"age":    user.Age,
			"weight": user.Weight,
		},
	}
	c.JSON(http.StatusOK, response)
}

func LoginUser(c *gin.Context) {
	user := &m.User{}
	c.Bind(user)
	user.Email = strings.ToLower(user.Email)
	inputPassword := user.Password

	err := m.LoginUser(user)
	if err != nil {
		response := &m.Response{
			Message: err.Error(),
		}
		c.JSON(http.StatusServiceUnavailable, response)
		c.Abort()
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputPassword))
	if err != nil {
		response := &m.Response{
			Message:    "Error : Unauthorized User",
			Success:    false,
			StatusCode: http.StatusUnauthorized,
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	if user.Token == "" {
		user.Token = randToken(20)
		_, err := m.UpdateToken(user)
		if err != nil {
			response := &m.Response{
				Message: err.Error(),
			}
			c.JSON(http.StatusServiceUnavailable, response)
			c.Abort()
			return
		}
	}

	response := &m.Response{
		Message:    "Get user : Certain user detail has been shown",
		Success:    true,
		StatusCode: http.StatusOK,
		Data: map[string]interface{}{
			"nama":   user.Name,
			"email":  user.Email,
			"age":    user.Age,
			"weight": user.Weight,
			"token":  user.Token,
		},
	}

	c.JSON(http.StatusOK, response)
}

func LogoutUser(c *gin.Context) {
	logout := &m.User{}

	authorization := c.Request.Header.Get("Authorization")
	authorization = strings.TrimPrefix(authorization, "Bearer ")
	fmt.Println(authorization)
	fmt.Println(logout.Token)

	err := m.LogoutUser(logout, authorization)
	if err != nil {
		response := &m.Response{
			Message: err.Error(),
		}
		c.JSON(http.StatusServiceUnavailable, response)
		c.Abort()
		return
	}

	response := &m.Response{
		Message:    "Logout successful",
		Success:    true,
		StatusCode: http.StatusOK,
	}

	c.JSON(http.StatusOK, response)
}

func PostGet(c *gin.Context) {
	users, err := m.GetUsers()
	if err != nil {
		response := &m.Response{
			Message: err.Error(),
		}
		c.JSON(http.StatusServiceUnavailable, response)
		c.Abort()
		return
	}

	response := &m.Response{
		Message: "Get users",
		Data:    users,
	}

	c.JSON(http.StatusOK, response)
}

func PostDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response := &m.Response{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	user, err := m.GetUser(id)
	if err != nil {
		response := &m.Response{
			Message: err.Error(),
		}
		c.JSON(http.StatusServiceUnavailable, response)
		c.Abort()
		return
	}

	response := &m.Response{
		Message: "Get user",
		Data:    user,
	}

	c.JSON(http.StatusOK, response)
}

func PostCreate(c *gin.Context) {
	req := &m.User{}
	err := c.BindJSON(&req)
	if err != nil {
		response := &m.Response{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	user, err := m.CreateUser(req)
	if err != nil {
		response := &m.Response{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	response := &m.Response{
		Message: "User has been created",
		Data:    user,
	}

	c.JSON(http.StatusCreated, response)
}

func PostUpdate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response := &m.Response{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	req := &m.User{}
	err = c.BindJSON(&req)
	if err != nil {
		response := &m.Response{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	req.ID = id
	user, err := m.UpdateUser(req)
	if err != nil {
		response := &m.Response{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	response := &m.Response{
		Message: "User has been updated",
		Data:    user,
	}

	c.JSON(http.StatusOK, response)
}

func PostDelete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response := &m.Response{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	err = m.DeleteUser(id)
	if err != nil {
		response := &m.Response{
			Message: err.Error(),
		}
		c.JSON(http.StatusServiceUnavailable, response)
		c.Abort()
		return
	}

	response := &m.Response{
		Message: "User has been deleted",
	}

	c.JSON(http.StatusOK, response)
}
