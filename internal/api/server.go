package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PromoCodes []struct {
	Store     string
	PromoCode string
	Discount  string
}

func StartServer() {
	log.Println("Server start up")
	promo := PromoCodes{
		{"OZON", "FRESH500", "500р"},
		{"Летуаль", "STYLE", "1000р"},
		{"ДОДО", "4093", "20%"},
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

	r.GET("/promo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "promo.tmpl", gin.H{
			"Promo": promo,
		})
	})

	r.Static("/image", "./resources")

	r.Run()

	log.Println("Server down")
}
