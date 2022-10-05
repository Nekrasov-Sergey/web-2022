package app

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type PromoCodes []struct {
	Store     string
	PromoCode string
	Discount  string
}

func (a *Application) StartServer() {
	log.Println("server start up")

	promo := PromoCodes{
		{"OZON", "FRESH500", "500р"},
		{"Летуаль", "STYLE", "1000р"},
		{"ДОДО", "4093", "20%"},
	}

	r := gin.Default()

	r.GET("/ping/:name", a.Ping)

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

			promo, err := a.repo.GetPromoByID(uint(intID))
			if err != nil {
				log.Printf("can't get promo by id %v", err)
				c.Error(err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"Price": promo.Price,
				"Promo": promo.Promo,
				"Store": promo.Store,
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
				a.repo.NewRandRecords()
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

	r.GET("/promo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "promo.tmpl", gin.H{
			"Promo": promo,
		})
	})

	r.Static("/image", "./resources")

	r.Run()

	log.Println("server down")
}

type pingReq struct{}
type pingResp struct {
	Status string `json:"status"`
}

// Ping godoc
// @Summary      Show hello text
// @Description  very very friendly response
// @Tags         Tests
// @Produce      json
// @Success      200  {object}  pingResp
// @Router       /ping/{name} [get]

func (a *Application) Ping(gCtx *gin.Context) {
	name := gCtx.Param("name")
	gCtx.String(http.StatusOK, "Hello %s", name)
}
