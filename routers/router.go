package routers

import (
	"github.com/Rickykn/buddyku-app.git/database"
	"github.com/Rickykn/buddyku-app.git/handlers"
	"github.com/Rickykn/buddyku-app.git/middlewares"
	"github.com/Rickykn/buddyku-app.git/repositories"
	"github.com/Rickykn/buddyku-app.git/services"
	"github.com/gin-gonic/gin"
)

func Server() *gin.Engine {

	engine := gin.Default()
	errConnect := database.Connect()

	// engine.Static("/docs", "swaggerui")

	ur := repositories.NewUserRepository(&repositories.URConfig{
		DB: database.Get(),
	})

	ar := repositories.NewAdminRepository(&repositories.ARConfig{
		DB: database.Get(),
	})

	us := services.NewUserService(&services.USConfig{
		UserRepository: ur,
	})

	as := services.NewAdminService(&services.ASConfig{
		AdminRespository: ar,
	})

	h := handlers.New(&handlers.HandlerConfig{

		UserService:  us,
		AdminService: as,
	})

	users := engine.Group("/users")
	{

		users.POST("/register", h.RegisterUser)
		users.POST("/login", h.LoginUser)
		users.POST("/article", middlewares.AuthorizeJWT, h.UserCreateNewPost)
		users.GET("/get-point", middlewares.AuthorizeJWT, h.GetPointUser)
		users.GET("/detail-article/:id", middlewares.AuthorizeJWT, h.GetArticleDetail)
	}

	admin := engine.Group("/admins")

	{
		admin.POST("/register", h.RegisterAdmin)
		admin.POST("/login", h.LoginAdmin)

		admin.POST("/setpoint", middlewares.AuthorizeJWTAdmin, h.SetPointReward)
		admin.GET("/all-user", middlewares.AuthorizeJWTAdmin, h.GetAllUser)
	}

	if errConnect != nil {
		panic(errConnect)
	}
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return engine
}
