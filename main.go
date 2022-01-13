package main

import (
	"gin_webservice/handlers"
	"gin_webservice/models"

	_ "gin_webservice/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title User API
// @description Simplest User API.
// @schemes http
// @host localhost:8080
func main() {
	engine := setupServer()
	models.ConnectDatabase()
	engine.Run()
}

func setupServer() *gin.Engine {
	engine := gin.Default()
	setRoutes(engine)
	return engine
}

func setRoutes(engine *gin.Engine) {
	user_grp := engine.Group("/users")
	{
		user_grp.GET("user", handlers.GetAllUsers)
		user_grp.POST("user", handlers.CreateUser)
		user_grp.GET("user/:id", handlers.GetUserByID)
		user_grp.PUT("user/:id", handlers.UpdateUser)
		user_grp.DELETE("user/:id", handlers.DeleteUser)
	}

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
