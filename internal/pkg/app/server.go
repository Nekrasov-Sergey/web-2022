package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"main/internal/app/model"
	"main/swagger/promos"
	"net/http"
)

func (a *Application) StartServer() {
	log.Println("server start up")

	r := gin.Default()

	r.GET("/promos/get", a.GetPromos)

	r.GET("/promos/get", a.GetPromos)

	r.POST("/promos/create", a.CreatePromo)

	r.POST("/promos/create/random", a.CreateRandomPromo)

	r.PUT("/promos/change/price", a.ChangePrice)

	r.DELETE("/promos/delete", a.DeletePromo)

	_ = r.Run()

	log.Println("server down")
}

// GetPromos 		godoc
// @Summary      	Get all records
// @Description  	Get a list of all promos
// @Tags         	Info
// @Produce      	json
// @Success      	200 {object} model.PromosDocs
// @Failure 		400 {object} promos.PromoError
// @Failure 		404 {object} promos.PromoError
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
// @Param 			Discount query uint64 true "Скидка"
// @Param 			Price query uint64 true "Цена"
// @Param 			Quantity query uint64 true "Количество"
// @Param 			Promo query []string true "Промокоды(запись в виде массива)"
// @Success 		201 {object} promos.PromoCreated
// @Failure 		400 {object} promos.PromoError
// @Failure 		404 {object} promos.PromoError
// @Failure 		500 {object} promos.PromoError
// @Router  		/promos/create [post]
func (a *Application) CreatePromo(gCtx *gin.Context) {
	promo := model.Promos{}
	if err := gCtx.BindJSON(&promo); err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&promos.PromoError{
				Description: "adding failed",
				Error:       "db error",
				Type:        "internal",
			})
		return
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
// @Param				Quantity query uint64 true "Количество"
// @Success     		201 {object} promos.PromoCreated
// @Failure 			400 {object} promos.PromoError
// @Failure 			404 {object} promos.PromoError
// @Failure 			500 {object} promos.PromoError
// @Router       		/promos/create/random [post]
func (a *Application) CreateRandomPromo(gCtx *gin.Context) {
	promo := model.Promos{}
	if err := gCtx.BindJSON(&promo); err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&promos.PromoError{
				Description: "adding failed",
				Error:       "db error",
				Type:        "internal",
			})
		return
	}
	for i := 0; i < int(promo.Quantity); i++ {
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
// @Param 			UUID query string true "UUID промо" format(uuid)
// @Param 			Price query uint64 true "Новая цена"
// @Success      	200 {object} promos.PromoChanged
// @Failure 		400 {object} promos.PromoError
// @Failure 		404 {object} promos.PromoError
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
// @Param 			UUID query string true "UUID промо" format(uuid)
// @Success      	200 {object} promos.PromoDeleted
// @Failure 		400 {object} promos.PromoError
// @Failure 		404 {object} promos.PromoError
// @Failure 	 	500 {object} promos.PromoError
// @Router       	/promos/delete [delete]
func (a *Application) DeletePromo(gCtx *gin.Context) {
	UUID := gCtx.Query("UUID")
	err := a.repo.DeletePromo(UUID)
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
