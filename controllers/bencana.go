package controllers

import (
	"github.com/gin-gonic/gin"
	m "memperbaikikode/models"
	"net/http"
	"strconv"
)

func BencanaGet(c *gin.Context) {
	bencana, err := m.GetSeveralBencana()
	if err != nil {
		response := &m.Response{
			Message: err.Error(),
		}
		c.JSON(http.StatusServiceUnavailable, response)
		c.Abort()
		return
	}

	response := &m.Response{
		Message: "Get several disasters",
		Data:    bencana,
	}

	c.JSON(http.StatusOK, response)
}

func BencanaDetail(c *gin.Context) {
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

	bencana, err := m.GetBencana(id)
	if err != nil {
		response := &m.Response{
			Message: err.Error(),
		}
		c.JSON(http.StatusServiceUnavailable, response)
		c.Abort()
		return
	}

	response := &m.Response{
		Message: "Get disaster",
		Data:    bencana,
	}

	c.JSON(http.StatusOK, response)
}

func BencanaCreate(c *gin.Context) {
	req := &m.Bencana{}
	err := c.BindJSON(&req)
	if err != nil {
		response := &m.Response{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	bencana, err := m.CreateBencana(req)
	if err != nil {
		response := &m.Response{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	response := &m.Response{
		Message: "Disaster has been created",
		Data:    bencana,
	}

	c.JSON(http.StatusCreated, response)
}

func BencanaUpdate(c *gin.Context) {
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

	req := &m.Bencana{}
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
	bencana, err := m.UpdateBencana(req)
	if err != nil {
		response := &m.Response{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	response := &m.Response{
		Message: "Disaster has been updated",
		Data:    bencana,
	}

	c.JSON(http.StatusOK, response)
}

// func BencanaDelete(c *gin.Context) {
// 	idStr := c.Param("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		response := &m.Response{
// 			Message: err.Error(),
// 		}
// 		c.JSON(http.StatusBadRequest, response)
// 		c.Abort()
// 		return
// 	}

// 	err = m.DeleteBencana(id)
// 	if err != nil {
// 		response := &m.Response{
// 			Message: err.Error(),
// 		}
// 		c.JSON(http.StatusServiceUnavailable, response)
// 		c.Abort()
// 		return
// 	}

// 	response := &m.Response{
// 		Message: "Bencana has been deleted",
// 	}

// 	c.JSON(http.StatusOK, response)
// }
