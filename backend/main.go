package main

import (
	"url-shortener-be/controllers"
	"url-shortener-be/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	
	db.ConnectDB()
	
	r.GET("/url/:id", controllers.GetShortUrl)
	r.POST("/url", controllers.CreateShortUrl)
	r.POST("/url/:id/visit", controllers.VisitShortUrl)
	
	r.Run(":3001")
}
