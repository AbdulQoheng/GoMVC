package routes

import (
	"net/http"

	"github.com/AbdulQoheng/go-mvc/config"
	"github.com/gin-gonic/gin"
)

type routess struct {
	route *gin.Engine
}

func Routes() routess {

	r := routess{
		route: gin.Default(),
	}

	v1 := r.route.Group("/")
	// web
	r.addWeb(v1)
	r.addApi(v1)

	return r
}
func (r routess) Run() error {

	route := r.route
	// route.LoadHTMLGlob("gomvc/view/*.html")
	// route.LoadHTMLGlob("resources/views/*.html")
	defer config.CloseDatabaseConnection(db)
	files := []string{
		"gomvc/view/abort.html",
		"resources/views/template/layout.html",
		"resources/views/template/header.html",
		"resources/views/template/footer.html",
		"resources/views/welcome.html",
		"resources/views/admin/admin.dashboard.html",
	}
	route.LoadHTMLFiles(files...)
	route.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "abort.html", gin.H{
			"status": http.StatusNotFound,
			"msg":    "Not Found",
		})
	})

	return route.Run(":8000")
}
