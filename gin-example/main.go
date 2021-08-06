package main

import (
	"gin-example/controller"
	"gin-example/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	engine := gin.Default()

	engine.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong")
	})

	engine.Static("/resource", "./resource")
	engine.StaticFile("/favicon.ico", "./resource/favicon.ico")

	videoController := controller.NewVideoController()
	videoGroup := engine.Group("/video")
	videoGroup.Use(middleware.MyLogger())
	videoGroup.Use(middleware.MyAuth())

	videoGroup.GET("/", videoController.GetAll)
	videoGroup.POST("/", videoController.Create)
	videoGroup.PUT("/:id", videoController.Update)
	videoGroup.DELETE("/:id", videoController.Delete)

	log.Fatalln(engine.Run("localhost:8080"))
}
