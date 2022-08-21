package main

import (
	"github.com/AbdulQoheng/go-mvc/routes"
)

func main() {
	router := routes.Routes()
	router.Run()
}

// package main

// import (
// 	"github.com/AbdulQoheng/go-mvc/app/http/controllers/api"
// 	"github.com/AbdulQoheng/go-mvc/app/http/middleware"
// 	"github.com/AbdulQoheng/go-mvc/app/models"
// 	"github.com/AbdulQoheng/go-mvc/config"
// 	"github.com/AbdulQoheng/go-mvc/service"
// 	"github.com/gin-gonic/gin"

// 	"gorm.io/gorm"
// )

// var (
// 	db             *gorm.DB            = config.SetupDatabaseConnection()
// 	userRepository models.UserModels   = models.NewUserModels(db)
// 	bookRepository models.BookModels   = models.NewBookModels(db)
// 	jwtService     service.JWTService  = service.NewJWTService()
// 	userService    service.UserService = service.NewUserService(userRepository)
// 	bookService    service.BookService = service.NewBookService(bookRepository)
// 	authService    service.AuthService = service.NewAuthService(userRepository)
// 	authController api.AuthController  = api.NewAuthController(authService, jwtService)
// 	userController api.UserController  = api.NewUserController(userService, jwtService)
// 	bookController api.BookController  = api.NewBookController(bookService, jwtService)
// )

// func main() {
// 	defer config.CloseDatabaseConnection(db)
// 	r := gin.Default()

// 	authRoutes := r.Group("api/auth")
// 	{
// 		authRoutes.POST("/login", authController.Login)
// 		authRoutes.POST("/register", authController.Register)
// 	}

// 	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
// 	{
// 		userRoutes.GET("/profile", userController.Profile)
// 		userRoutes.PUT("/profile", userController.Update)
// 	}

// 	bookRoutes := r.Group("api/books", middleware.AuthorizeJWT(jwtService))
// 	{
// 		bookRoutes.GET("/", bookController.All)
// 		bookRoutes.POST("/", bookController.Insert)
// 		bookRoutes.GET("/:id", bookController.FindByID)
// 		bookRoutes.PUT("/:id", bookController.Update)
// 		bookRoutes.DELETE("/:id", bookController.Delete)
// 	}

// 	r.Run(":8000")
// }
