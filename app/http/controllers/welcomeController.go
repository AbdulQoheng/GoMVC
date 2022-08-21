package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type welcomeController struct{}

// UserController is a ....
type WelcomeInterface interface {
	Index(context *gin.Context)
}

// NewUserController is creating anew instance of UserControlller
func WelcomeController() WelcomeInterface {
	return &welcomeController{}
}

func (c *welcomeController) Index(context *gin.Context) {
	context.HTML(http.StatusOK, "welcome.html", gin.H{
		"title": "Main website",
	})

}
