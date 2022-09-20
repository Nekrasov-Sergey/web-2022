package app

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
	log.Println("server start up")

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
			log.Printf("id received %s\n", id)
			intID, err := strconv.Atoi(id)
			if err != nil {
				log.Printf("can't convert id %v", err)
				c.Error(err)
				return
			}

			product, err := a.repo.GetProductByID(uint(intID))
			if err != nil {
				log.Printf("can't get product by id %v", err)
				c.Error(err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"product_price": product.Price,
			})
			return
		}

		create := c.Query("create")
		if create != "" {
			log.Printf("create received %s\n", create)
			createBool, err := strconv.ParseBool(create) // пытаемся привести это к чиселке
			if err != nil {                              // если не получилось
				log.Printf("can't convert create %v", err)
				c.Error(err)
				return
			}

			if createBool {
				a.repo.NewRandRecord()
				c.JSON(http.StatusOK, gin.H{
					"status": "ok",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": "create not true",
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

	log.Println("server down")
}
