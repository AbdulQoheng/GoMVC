package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct{}

// UserController is a ....
type UserInterface interface {
	Index(context *gin.Context)
}

// NewUserController is creating anew instance of UserControlller
func UserController() UserInterface {
	return &userController{}
}

func (c *userController) Index(context *gin.Context) {
	context.HTML(http.StatusOK, "abort.html", gin.H{
		"title": "Main website",
	})

}
