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

	r.GET("/promos", a.GetPromos)

	r.GET("/promos/:uuid", a.GetPromoPrice)

	r.POST("/promos/", a.CreatePromo)

	r.POST("/promos/random", a.CreateRandomPromo)

	r.PUT("/promos/:uuid/", a.ChangePrice)

	r.DELETE("/promos/:uuid/", a.DeletePromo)

	_ = r.Run()

	log.Println("server down")
}

// GetPromos 		godoc
// @Summary      	Get all records
// @Description  	Get a list of all promos
// @Tags         	Info
// @Produce      	json
// @Success      	200 {object} model.PromosDocs
// @Failure 		500 {object} promos.PromoError
// @Router       	/promos [get]
func (a *Application) GetPromos(gCtx *gin.Context) {
	resp, err := a.repo.GetPromos()
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&promos.PromoError{
				Description: "Can't get a list of promo codes",
				Error:       promos.Err500,
				Type:        promos.TypeInternalReq,
			})
		return
	}

	gCtx.JSON(http.StatusOK, resp)
}

// GetPromoPrice  	godoc
// @Summary      	Get price for a promo
// @Description  	Get the price using the promo uuid
// @Tags         	Info
// @Produce      	json
// @Param 			UUID query string true "UUID промо" format(uuid)
// @Success      	200 {object} promos.PromoPrice
// @Failure 	 	400 {object} promos.PromoError
// @Failure 	 	404 {object} promos.PromoError
// @Failure 	 	500 {object} promos.PromoError
// @Router       	/promos/:uuid [get]
func (a *Application) GetPromoPrice(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&promos.PromoError{
				Description: "Invalid UUID format",
				Error:       promos.Err400,
				Type:        promos.TypeClientReq,
			})
		return
	}

	var promo model.Promos
	code, err := a.repo.GetPromoPrice(UUID, &promo)
	if err != nil {
		if code == 404 {
			gCtx.JSON(
				http.StatusNotFound,
				&promos.PromoError{
					Description: "UUID Not Found",
					Error:       promos.Err404,
					Type:        promos.TypeClientReq,
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&promos.PromoError{
					Description: "Get promo price failed",
					Error:       promos.Err500,
					Type:        promos.TypeInternalReq,
				})
			return
		}
	}

	gCtx.JSON(
		http.StatusOK,
		&promos.PromoPrice{
			Price: promo.Price,
		})
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
// @Failure 		500 {object} promos.PromoError
// @Router  		/promos/ [post]
func (a *Application) CreatePromo(gCtx *gin.Context) {
	promo := model.Promos{}
	err := gCtx.BindJSON(&promo)
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&promos.PromoError{
				Description: "Invalid parameters",
				Error:       promos.Err400,
				Type:        promos.TypeClientReq,
			})
		return
	}

	err = a.repo.AddPromo(promo)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&promos.PromoError{
				Description: "Create failed",
				Error:       promos.Err500,
				Type:        promos.TypeInternalReq,
			})
		return
	}

	gCtx.JSON(
		http.StatusCreated,
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
// @Failure 			500 {object} promos.PromoError
// @Router       		/promos/random [post]
func (a *Application) CreateRandomPromo(gCtx *gin.Context) {
	amount := model.Amount{}
	err := gCtx.BindJSON(&amount)
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&promos.PromoError{
				Description: "The quantity is negative or not int",
				Error:       promos.Err400,
				Type:        promos.TypeClientReq,
			})
		return
	}

	for i := 0; i < int(amount.Amount); i++ {
		err = a.repo.NewRandRecords()
		if err != nil {
			gCtx.JSON(
				http.StatusInternalServerError,
				&promos.PromoError{
					Description: "Create random promo failed",
					Error:       promos.Err500,
					Type:        promos.TypeInternalReq,
				})
			return
		}
	}

	gCtx.JSON(
		http.StatusCreated,
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
// @Router       	/promos/:uuid/ [put]
func (a *Application) ChangePrice(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&promos.PromoError{
				Description: "Invalid UUID format",
				Error:       promos.Err400,
				Type:        promos.TypeClientReq,
			})
		return
	}

	promo := model.Promos{}
	err = gCtx.BindJSON(&promo)
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&promos.PromoError{
				Description: "The price is negative or not int",
				Error:       promos.Err400,
				Type:        promos.TypeClientReq,
			})
		return
	}

	code, err := a.repo.ChangePrice(UUID, promo.Price)
	if err != nil {
		if code == 404 {
			gCtx.JSON(
				http.StatusNotFound,
				&promos.PromoError{
					Description: "UUID Not Found",
					Error:       promos.Err404,
					Type:        promos.TypeClientReq,
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&promos.PromoError{
					Description: "Change failed",
					Error:       promos.Err500,
					Type:        promos.TypeInternalReq,
				})
			return
		}
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
// @Router       	/promos/:uuid [delete]
func (a *Application) DeletePromo(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&promos.PromoError{
				Description: "Invalid UUID format",
				Error:       promos.Err400,
				Type:        promos.TypeClientReq,
			})
		return
	}

	code, err := a.repo.DeletePromo(UUID)
	if err != nil {
		if code == 404 {
			gCtx.JSON(
				http.StatusNotFound,
				&promos.PromoError{
					Description: "UUID Not Found",
					Error:       promos.Err404,
					Type:        promos.TypeClientReq,
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&promos.PromoError{
					Description: "Delete failed",
					Error:       promos.Err500,
					Type:        promos.TypeInternalReq,
				})
			return
		}
	}

	gCtx.JSON(
		http.StatusOK,
		&promos.PromoDeleted{
			Success: true,
		})
}
