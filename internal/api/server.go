package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Gift struct {
	Name  string
	Price int
}

type BirthdayGifts struct {
	Person string
	Gifts  []Gift
}

func StartServer() {
	log.Println("Server start up")

	list := BirthdayGifts{
		Person: "Sergey Nekrasov",
		Gifts: []Gift{
			{"Laptop", 45000},
			{"Bicycle", 20000},
			{"Sneakers", 6000},
		},
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.LoadHTMLGlob("templates/*")

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	r.GET("/gifts", func(c *gin.Context) {
		c.HTML(http.StatusOK, "gifts.tmpl", gin.H{
			"Person": list.Person,
			"Gifts":  list.Gifts,
		})
	})

	r.Static("/image", "./resources")

	r.Run()

	log.Println("Server down")
}
