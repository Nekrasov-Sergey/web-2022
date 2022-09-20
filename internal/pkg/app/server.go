package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Gift struct {
	Name  string
	Price int
}

type BirthdayGifts struct {
	Person string
	Gifts  []Gift
}

func (a *Application) StartServer() {
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
		id := c.Query("id") // получаем из запроса query string

		if id != "" {
			log.Printf("id recived %s\n", id)
			intID, err := strconv.Atoi(id) // пытаемся привести это к чиселке
			if err != nil {                // если не получилось
				log.Printf("cant convert id %v", err)
				c.Error(err)
				return
			}

			product, err := a.repo.GetProductByID(uint(intID))
			if err != nil { // если не получилось
				log.Printf("cant get product by id %v", err)
				c.Error(err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"product_price": product.Price,
			})
			return
		}
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
