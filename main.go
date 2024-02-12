package main

import (
	"EyeOnUrls/controllers"
	"EyeOnUrls/docs"
	"EyeOnUrls/models"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RunWebServer() {
	r := gin.Default()
	docs.SwaggerInfo.Title = "EyeOnUrls API"
	docs.SwaggerInfo.Description = "This is a simple URL shortener API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"

	r.GET("/:short_url", controllers.FetchUrl)
	r.POST("/", controllers.CreateUrl)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err := r.Run()
	if err != nil {
		panic(err)
	}

	return
}

func main() {
	models.ConnectDatabase()
	RunWebServer()
}
