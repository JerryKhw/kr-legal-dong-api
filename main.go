package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "kr-legal-dong-api/db"
	"kr-legal-dong-api/docs"
	"kr-legal-dong-api/model"
	v1 "kr-legal-dong-api/v1"
)

func main() {
	r := gin.Default()

	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = "kr-legal-dong-api.fly.dev"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Title = "대한민국 법정동 API"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "kr-legal-dong-api")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, model.DefaultResponse{
			Message: "pong",
		})
	})

	v1.SetRoute(r)

	r.Run()
}
