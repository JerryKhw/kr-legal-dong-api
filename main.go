package main

import (
	"encoding/json"
	"kr-legal-dong-api/model"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func initDB() {
	siBytes, err := os.ReadFile("./kr-legal-dong/data/si.json")
	if err != nil {
		panic(err)
	}

	var si []model.Si

	err = json.Unmarshal(siBytes, &si)
	if err != nil {
		panic(err)
	}

	guBytes, err := os.ReadFile("./kr-legal-dong/data/gu.json")
	if err != nil {
		panic(err)
	}

	var gu []model.Gu

	err = json.Unmarshal(guBytes, &gu)
	if err != nil {
		panic(err)
	}

	dongBytes, err := os.ReadFile("./kr-legal-dong/data/dong.json")
	if err != nil {
		panic(err)
	}

	var dong []model.Dong

	err = json.Unmarshal(dongBytes, &dong)
	if err != nil {
		panic(err)
	}

	detailBytes, err := os.ReadFile("./kr-legal-dong/data/detail.json")
	if err != nil {
		panic(err)
	}

	var detail []model.Detail

	err = json.Unmarshal(detailBytes, &detail)
	if err != nil {
		panic(err)
	}
}

func setEnv() {
	if os.Getenv("APP_MODE") != "prod" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	setEnv()

	initDB()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, os.Getenv("APP_NAME"))
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, model.DefaultResponse{
			Message: "pong",
		})
	})

	r.Run()
}
