package main

import (
	"database/sql"
	"encoding/json"
	"kr-legal-dong-api/model"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
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

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.Exec("CREATE TABLE `si` (`code` varchar(10) NOT NULL PRIMARY KEY, `name` varchar(100) NOT NULL, `active` integer(1) NOT NULL)")

	for _, si := range si {
		db.Exec("INSERT INTO `si` (`code`, `name`, `active`) VALUES (?, ?, ?)", si.Code, si.Name, si.Active)
	}

	db.Exec("CREATE TABLE `gu` (`code` varchar(10) NOT NULL PRIMARY KEY, `si_code` varchar(10) NOT NULL, `name` varchar(100) NOT NULL, `active` integer(1) NOT NULL, FOREIGN KEY(`si_code`) REFERENCES `si`(`code`))")

	for _, gu := range gu {
		db.Exec("INSERT INTO `gu` (`code`, `si_code`, `name`, `active`) VALUES (?, ?, ?, ?)", gu.Code, gu.SiCode, gu.Name, gu.Active)
	}

	db.Exec("CREATE TABLE `dong` (`code` varchar(10) NOT NULL PRIMARY KEY, `gu_code` varchar(10) NOT NULL, `name` varchar(100) NOT NULL, `active` integer(1) NOT NULL, FOREIGN KEY(`gu_code`) REFERENCES `gu`(`code`))")

	for _, dong := range dong {
		db.Exec("INSERT INTO `dong` (`code`, `gu_code`, `name`, `active`) VALUES (?, ?, ?, ?)", dong.Code, dong.GuCode, dong.Name, dong.Active)
	}

	db.Exec("CREATE TABLE `detail` (`code` varchar(10) NOT NULL PRIMARY KEY, `dong_code` varchar(10) NOT NULL, `name` varchar(100) NOT NULL, `active` integer(1) NOT NULL, FOREIGN KEY(`dong_code`) REFERENCES `dong`(`code`))")

	for _, detail := range detail {
		db.Exec("INSERT INTO `detail` (`code`, `dong_code`, `name`, `active`) VALUES (?, ?, ?, ?)", detail.Code, detail.DongCode, detail.Name, detail.Active)
	}
}

func main() {
	initDB()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "kr-legal-dong-api")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, model.DefaultResponse{
			Message: "pong",
		})
	})

	r.Run()
}
