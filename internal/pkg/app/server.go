package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"main/internal/app/model"
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

	r.PUT("/promos/change/price", a.ChangePrice)

	r.DELETE("/promos/delete", a.DeletePromo)

	r.Run()

	log.Println("server down")
}

// GetPromos 		godoc
// @Summary      	Get all records
// @Description  	Get a list of all promos
// @Tags         	Info
// @Produce      	json
// @Success      	200 {object} model.PromosDocs
// @Failure 		500 {object} promos.PromoError
// @Router       	/promos/get [get]
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

// CreatePromo		godoc
// @Summary     	Add a new promo
// @Description		Adding a new promo to database
// @Tags			Add
// @Produce      	json
// @Param 			Store query string true "Магазин"
// @Param 			Discount query string true "Скидка"
// @Param 			Price query string true "Цена"
// @Param 			Quantity query uint64 true "Количество"
// @Param 			Promo query []string true "Промокоды(запись в виде массива)"
// @Success 		201 {object} promos.PromoCreated
// @Failure 		500 {object} promos.PromoError
// @Router  		/promos/create [Post]
func (a *Application) CreatePromo(gCtx *gin.Context) {
	quantity, _ := strconv.ParseUint(gCtx.Query("Quantity"), 10, 64)
	promo := model.Promos{
		Store:    gCtx.Query("Store"),
		Discount: gCtx.Query("Discount"),
		Price:    gCtx.Query("Price"),
		Quantity: quantity,
		Promo:    pq.StringArray{gCtx.Query("Promo")},
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
	gCtx.JSON(
		http.StatusOK,
		&promos.PromoCreated{
			Success: true,
		})
}

// CreateRandomPromo 	godoc
// @Summary      		Add a new random promo
// @Description  		Adding a new random promo to database
// @Tags        		Add
// @Produce      		json
// @Param				Quantity query int64 true "Количество"
// @Success     		201 {object} promos.PromoCreated
// @Failure 			500 {object} promos.PromoError
// @Router       		/promos/create/random [Post]
func (a *Application) CreateRandomPromo(gCtx *gin.Context) {
	quantity, _ := strconv.ParseInt(gCtx.Query("Quantity"), 10, 64)
	for i := 0; i < int(quantity); i++ {
		err := a.repo.NewRandRecords()
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
	}
	gCtx.JSON(
		http.StatusOK,
		&promos.PromoCreated{
			Success: true,
		})
}

// ChangePrice		godoc
// @Summary      	Change promo price
// @Description  	Change the promo price using its uuid
// @Tags         	Change
// @Produce      	json
// @Param 			UUID query string true "UUID промо"
// @Param 			Price query string true "Новая цена"
// @Success      	200  {object}  promos.PromoChanged
// @Failure 	 	500 {object} promos.PromoError
// @Router       	/promos/change/price [put]
func (a *Application) ChangePrice(gCtx *gin.Context) {
	inputUuid, _ := uuid.Parse(gCtx.Query("UUID"))
	newPrice := gCtx.Query("Price")
	err := a.repo.ChangePrice(inputUuid, newPrice)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&promos.PromoError{
				Description: "update failed",
				Error:       "db error",
				Type:        "internal",
			})
		return
	}
	gCtx.JSON(
		http.StatusOK,
		&promos.PromoChanged{
			Success: true,
		})
}

// DeletePromo		godoc
// @Summary     	Delete a promo
// @Description 	Delete a promo using its uuid
// @Tags         	Delete
// @Produce      	json
// @Param 			UUID query string true "UUID промо"
// @Success      	200 {object} promos.PromoDeleted
// @Failure 	 	500 {object} promos.PromoError
// @Router       	/promos/delete [delete]
func (a *Application) DeletePromo(gCtx *gin.Context) {
	uuid := gCtx.Query("UUID")
	err := a.repo.DeletePromo(uuid)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&promos.PromoError{
				Description: "delete failed",
				Error:       "db error",
				Type:        "internal",
			})
		return
	}
	gCtx.JSON(
		http.StatusOK,
		&promos.PromoDeleted{
			Success: true,
		})
}
