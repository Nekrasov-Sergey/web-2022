package app

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"main/internal/app/ds"
	"main/swagger/promos"
	"net/http"
	"strconv"
)

func (a *Application) StartServer() {
	log.Println("server start up")

	r := gin.Default()

	r.GET("/promos/get", a.GetPromos)

	r.POST("/promos/create", a.CreatePromo)

	r.POST("/promos/create/random", a.CreateRandomPromo)

	r.Run()

	log.Println("server down")
}

// GetPromos godoc
// @Summary      Get all records
// @Description  Get a list of all promos
// @Tags         Info
// @Produce      json
// @Success      200  {object}  ds.Promos
// @Failure 500 {object} promos.PromoError
// @Router       /promos/get [get]
func (a *Application) GetPromos(gCtx *gin.Context) {
	resp, err := a.repo.GetPromos()
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&promos.PromoError{
				Description: "can`t get a list",
				Error:       "db error",
				Type:        "internal",
			})
		return
	}
	gCtx.JSON(http.StatusOK, resp)
}

// CreatePromo godoc
// @Summary      Add a new promo
// @Description  Adding a new promo to database
// @Tags         Add
// @Produce      json
// @Param Store query string true "Магазин"
// @Param Discount query string true "Скидка"
// @Param Price query string true "Цена"
// @Param Quantity query uint64 true "Количество"
// @Param Promo query string true "Промокоды(запись в виде массива)"
// @Success      201  {object}  promos.PromoCreated
// @Failure 500 {object} promos.PromoError
// @Router       /promos/create [Post]
func (a *Application) CreatePromo(gCtx *gin.Context) {
	quantity, _ := strconv.ParseUint(gCtx.Query("Quantity"), 10, 64)
	promo := ds.Promos{
		Store:    gCtx.Query("Store"),
		Discount: gCtx.Query("Discount"),
		Price:    gCtx.Query("Price"),
		Quantity: quantity,
		Promo:    []byte(gCtx.Query("Promo")),
	}

	err := a.repo.AddPromo(promo)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&promos.PromoError{
				Description: "adding failed",
				Error:       "db error",
				Type:        "internal",
			})
		return
	}

	gCtx.JSON(http.StatusOK, promo)
}

func (a *Application) CreateRandomPromo(gCtx *gin.Context) {
	for i := 0; i < 5; i++ {
		promo, _ := a.repo.NewRandRecords()
		gCtx.JSON(http.StatusOK, promo)
	}
}
