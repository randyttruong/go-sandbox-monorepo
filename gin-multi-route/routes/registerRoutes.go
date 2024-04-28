package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	registerGroup := router.Group("/register")
	{
		registerGroup.POST("/", postRegisterQuery)
	}
}

func postRegisterQuery(c *gin.Context) {
	var username, password, email, firstname, lastname string
	var resp gin.H

	username = c.Query("username")
	password = c.Query("password")
	email = c.Query("email")
	firstname = c.Query("firstname")
	lastname = c.Query("lastname")

	if len(username) == 0 || len(password) == 0 || len(email) == 0 || len(firstname) == 0 || len(lastname) == 0 {
		resp = gin.H{"message": "Registration failed. Plesae provide all necessary arguments."}

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp = gin.H{"message": "Registration succssful."}

	c.JSON(http.StatusOK, resp)
	return
}

func postRegisterParams(c *gin.Context) {
	var username, password, email, firstname, lastname string
	var resp gin.H

	username = c.Param("username")
	password = c.Param("password")
	email = c.Param("email")
	firstname = c.Param("firstname")
	lastname = c.Param("lastname")

	if len(username) == 0 || len(password) == 0 || len(email) == 0 || len(firstname) == 0 || len(lastname) == 0 {
		resp = gin.H{"message": "Registration failed. Please provide the required parameters."}

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp = gin.H{"message": "Registration successful. Please login"}
	c.JSON(http.StatusOK, resp)
	return
}
