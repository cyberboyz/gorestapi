package main

import (
	"fmt"
	"github.com/appleboy/gofight"
	"github.com/buger/jsonparser"
	"github.com/stretchr/testify/assert"
	"memperbaikikode/routers"
	"net/http"
	"testing"
)

var token string

func TestRestfulAPI(t *testing.T) {
	r := gofight.New()

	r.POST("/v1/register").
		SetDebug(true).
		SetJSON(gofight.D{
			"email":    "abc@gmail.com",
			"password": "Om3th1N6n1C3",
			"name":     "Fattah Azzuhry",
		}).
		Run(routers.GetEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			data := []byte(r.Body.String())

			message, _ := jsonparser.GetString(data, "message")

			assert.Equal(t, "Registration is successful", message)
		})

	r.POST("/v1/login").
		SetDebug(true).
		SetJSON(gofight.D{
			"email":    "abc@gmail.com",
			"password": "Om3th1N6n1C3",
		}).
		Run(routers.GetEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			data := []byte(r.Body.String())

			message, _ := jsonparser.GetString(data, "message")
			success, _ := jsonparser.GetBoolean(data, "success")
			status_code, _ := jsonparser.GetInt(data, "status_code")
			token, _ = jsonparser.GetString(data, "data", "token")

			assert.Equal(t, "Get user : Certain user detail has been shown", message)
			assert.Equal(t, true, success)
			assert.Equal(t, int64(200), status_code)
			assert.Equal(t, http.StatusOK, r.Code)
			fmt.Println("Token adalah ", token)
		})

	r.GET("/v1/user/1").
		SetDebug(true).
		SetHeader(gofight.H{
			"Authorization": token,
		}).
		Run(routers.GetEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			// assert.Equal(t, "Get user", r.Body.String())
			data := []byte(r.Body.String())

			message, _ := jsonparser.GetString(data, "message")
			success, _ := jsonparser.GetBoolean(data, "success")

			assert.Equal(t, "Get user", message)
			assert.Equal(t, true, success)
			assert.Equal(t, http.StatusOK, r.Code)
		})

	r.PUT("/v1/user/1").
		SetDebug(true).
		SetHeader(gofight.H{
			"Authorization": token,
		}).
		Run(routers.GetEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			// assert.Equal(t, "Get user", r.Body.String())
			data := []byte(r.Body.String())

			message, _ := jsonparser.GetString(data, "message")
			success, _ := jsonparser.GetBoolean(data, "success")

			assert.Equal(t, "User has been updated", message)
			assert.Equal(t, true, success)
			assert.Equal(t, http.StatusOK, r.Code)
		})

	// r.POST("/v1/bencana").
	// 	SetDebug(true).
	// 	SetHeader(gofight.H{
	// 		"Authorization": token,
	// 	}).
	// 	SetJSON(gofight.D{
	// 		"kota":      "Jakarta",
	// 		"latitude":  "-6.1477",
	// 		"longitude": "116.2329",
	// 		"tewas":     "54",
	// 		"luka":      "16",
	// 	}).
	// 	Run(routers.GetEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
	// 		assert.Equal(t, "{\"success\":true,\"status_code\":200,\"message\":\"Registration is successful\",\"data\":{\"age\":0,\"email\":\"fattahazzuhry@gmail.com\",\"nama\":\"Fattah Azzuhry\",\"weight\":0}}\n", r.Body.String())
	// 		assert.Equal(t, http.StatusOK, r.Code)
	// 	})

	// r.GET("/v1/bencana/1").
	// 	SetDebug(true).
	// 	SetHeader(gofight.H{
	// 		"Authorization": token,
	// 	}).
	// 	Run(routers.GetEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
	// 		// assert.Equal(t, "Get user", r.Body.String())
	// 		data := []byte(r.Body.String())

	// 		message, _ := jsonparser.GetString(data, "message")
	// 		success, _ := jsonparser.GetBoolean(data, "success")

	// 		assert.Equal(t, "Get user", message)
	// 		assert.Equal(t, true, success)
	// 		assert.Equal(t, http.StatusOK, r.Code)
	// 	})

	// r.PUT("/v1/bencana/1").
	// 	SetDebug(true).
	// 	SetHeader(gofight.H{
	// 		"Authorization": token,
	// 	}).
	// 	Run(routers.GetEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
	// 		// assert.Equal(t, "Get user", r.Body.String())
	// 		data := []byte(r.Body.String())

	// 		message, _ := jsonparser.GetString(data, "message")
	// 		success, _ := jsonparser.GetBoolean(data, "success")

	// 		assert.Equal(t, "User has been updated", message)
	// 		assert.Equal(t, true, success)
	// 		assert.Equal(t, http.StatusOK, r.Code)
	// 	})

	// r.GET("/v1/logout").
	// 	SetDebug(true).
	// 	SetHeader(gofight.H{
	// 		"Authorization": token,
	// 	}).
	// 	Run(routers.GetEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
	// 		assert.Equal(t, "{\"success\":true,\"status_code\":200,\"message\":\"Logout successful\"}\n", r.Body.String())
	// 		assert.Equal(t, http.StatusOK, r.Code)
	// 	})

}
