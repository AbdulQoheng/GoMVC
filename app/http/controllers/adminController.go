package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type adminController struct{}

// AdminController is a ....
type AdminInterface interface {
	Index(context *gin.Context)
}

// NewUserController is creating anew instance of UserControlller
func AdminController() AdminInterface {
	return &adminController{}
}

func (c *adminController) Index(context *gin.Context) {
	context.HTML(http.StatusOK, "admin.dashboard.html", gin.H{
		"title": "Main website",
	})

}
