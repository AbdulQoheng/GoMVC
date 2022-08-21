package routes

import (
	"github.com/AbdulQoheng/go-mvc/app/http/controllers"
	"github.com/gin-gonic/gin"
)

func (r routess) addWeb(rg *gin.RouterGroup) {
	webs := rg.Group("/")

	webs.GET("/", controllers.WelcomeController().Index)
	webs.GET("/user", controllers.UserController().Index)
	webs.GET("/admin", controllers.AdminController().Index)
}
