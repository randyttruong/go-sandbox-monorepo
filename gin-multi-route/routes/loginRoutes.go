package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginRoutes(router *gin.Engine) {
	loginGroup := router.Group("/login")
	{
		loginGroup.POST("/", postLoginQuery)
	}
	fmt.Println("vim-go")
}

func postLoginQuery(c *gin.Context) {
	var username string
	var password string
	var resp gin.H

	username = c.Query("username")
	password = c.Query("password")

	if len(username) == 0 || len(password) == 0 {
		resp = gin.H{"message": "Please provide the required credentials in order to login."}

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if username != "randyt" {
		resp = gin.H{"message": "Login unsuccessful"}

	} else {
		resp = gin.H{"message": "Login successful"}
	}

	c.JSON(http.StatusOK, resp)
	return
}

func postLoginArgs(c *gin.Context) {
	var username, password string
	var resp gin.H

	username = c.Param("username")
	password = c.Param("password")

	if len(username) == 0 || len(password) == 0 {
		resp = gin.H{"message": "Login unsuccessful. Please provide the required parameters."}

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if username != "randyt" {
		resp = gin.H{"message": "Login unsuccessful. Incorrect username"}
	} else {
		resp = gin.H{"message": "Login successful. Welcome."}
	}

	c.JSON(http.StatusOK, resp)
	return
}

// Note:
// `"json:{fieldName}"`
// is known as a "struct field tag"
// This is used by Go's encoding/json package
// IT essentially tells the JSON encoder/decoder
// to map the attribute LoginForm.username to
// the "username" field in the JSON
// as well as the LoginForm.password attribute
// to the "password" field in the JSON
type LoginForm struct {
	username string  `"json:username"` // Mandatory Field
	password string  `"json:password"` // Mandatory Field
	email    *string `"json:email"`    // Optional Field
}

func postLoginJSON(c *gin.Context) {
	var loginForm LoginForm
	var err error
	var resp gin.H

	err = c.BindJSON(&loginForm)

	if err != nil {
		resp = gin.H{"error": err.Error()}

		c.JSON(
			http.StatusBadRequest,
			resp,
		)
		return
	}

	return
}

func setDefaults(loginForm *LoginForm) {
	var defaultEmail string = "swag123@gmail.com"
	if loginForm.email == nil {
		loginForm.email = &defaultEmail

	}
}

// In order to provide default values,
// we need to first define the default values,
// then, instead of utilizing BindJSON, we can
// use ShouldBindJSON(), which functions
// the same as BindJSON(), but also utilizes
// json.Unmarshal(), which simply allows us
// to handle the case when a field doesn't exist
// within the JSON body
//

func postLoginJSONDefault(c *gin.Context) {
	var loginForm LoginForm
	var err error
	var resp gin.H

	err = c.ShouldBindJSON(loginForm)

	if err != nil {
		resp = gin.H{"error": err.Error()}

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	setDefaults(&loginForm)

	if loginForm.username == "randyt" {
		resp = gin.H{"message": "Login successful"}
	} else {
		resp = gin.H{"message": "Login unsuccessful"}
	}

	c.JSON(http.StatusOK, resp)

	return
}
