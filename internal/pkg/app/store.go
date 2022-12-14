package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"main/internal/app/ds"
	"main/swagger"
	"net/http"
	"strconv"
	"strings"
)

// GetStores 		godoc
// @Summary      	Get all stores
// @Description  	Get a list of all stores
// @Tags         	Info
// @Produce      	json
// @Success      	200 {object} ds.StoreDocs
// @Failure 		500 {object} swagger.Error
// @Router       	/store [get]
func (a *Application) GetStores(gCtx *gin.Context) {
	resp, err := a.repo.GetStores()
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&swagger.Error{
				Description: "Can't get a list of promo codes",
				Error:       swagger.Err500,
				Type:        swagger.TypeInternalReq,
			})
		return
	}

	gCtx.JSON(http.StatusOK, resp)
}

// GetStore 		godoc
// @Summary      	Get store
// @Description  	Get store using its uuid
// @Tags         	Info
// @Produce      	json
// @Param 			UUID path string true "UUID магазина" format(uuid)
// @Success      	200 {object} ds.StoreDocs
// @Failure 		500 {object} swagger.Error
// @Router       	/store/{UUID} [get]
func (a *Application) GetStore(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	resp, err := a.repo.GetStore(UUID)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&swagger.Error{
				Description: "Can't get a list of promo codes",
				Error:       swagger.Err500,
				Type:        swagger.TypeInternalReq,
			})
		return
	}

	gCtx.JSON(http.StatusOK, resp)
}

// GetPromoStore			godoc
// @Summary     	Get a promo
// @Description 	Get a promo in store using its uuid
// @Tags         	Info
// @Produce      	json
// @Param 			UUID path string true "UUID магазина" format(uuid)
// @Param 			Quantity path string true "Кол-во"
// @Success      	200 {object} swagger.StorePromo
// @Failure 		400 {object} swagger.Error
// @Failure 		404 {object} swagger.Error
// @Failure 	 	500 {object} swagger.Error
// @Router       	/store/{UUID}/{Quantity} [get]
func (a *Application) GetPromoStore(gCtx *gin.Context) {
	jwtStr := gCtx.GetHeader("Authorization")
	userUUID := a.GetUserByToken(jwtStr)

	StoreUUID, err := uuid.Parse(gCtx.Param("uuid"))
	quantity, _ := strconv.Atoi(gCtx.Param("quantity"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&swagger.Error{
				Description: "Invalid UUID format",
				Error:       swagger.Err400,
				Type:        swagger.TypeClientReq,
			})
		return
	}

	code, Promo, err := a.repo.GetPromoStore(uint64(quantity), StoreUUID, userUUID)
	if err != nil {
		if code == 404 {
			gCtx.JSON(
				http.StatusNotFound,
				&swagger.Error{
					Description: "UUID Not Found",
					Error:       swagger.Err404,
					Type:        swagger.TypeClientReq,
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&swagger.Error{
					Description: "Delete failed",
					Error:       swagger.Err500,
					Type:        swagger.TypeInternalReq,
				})
			return
		}
	}

	gCtx.JSON(http.StatusOK, Promo)
}

// CreateStore		godoc
// @Summary     	Add a new promo
// @Description		Adding a new promo to database
// @Tags			Add
// @Produce      	json
// @Param 			Promo body ds.StoreDocs true "Магазин"
// @Success 		201 {object} swagger.Create
// @Failure 		400 {object} swagger.Error
// @Failure 		500 {object} swagger.Error
// @Router  		/store [post]
func (a *Application) CreateStore(gCtx *gin.Context) {
	store := ds.Store{}

	err := gCtx.BindJSON(&store)
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&swagger.Error{
				Description: "Invalid parameters",
				Error:       swagger.Err400,
				Type:        swagger.TypeClientReq,
			})
		return
	}

	if len(store.Promo) != int(store.Quantity) {
		store.Promo = strings.Split(store.Promo[0], ",")
	}

	err = a.repo.CreateStore(store)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&swagger.Error{
				Description: "Create failed",
				Error:       swagger.Err500,
				Type:        swagger.TypeInternalReq,
			})
		return
	}

	gCtx.JSON(
		http.StatusCreated,
		&swagger.Create{
			Success: true,
		})
}

// CreateRandomStores 	godoc
// @Summary      		Add a new random promo
// @Description  		Adding a new random promo to database
// @Tags        		Add
// @Produce      		json
// @Param				Quantity body ds.QuantityStores true "Количество"
// @Success     		201 {object} swagger.Create
// @Failure 			400 {object} swagger.Error
// @Failure 			500 {object} swagger.Error
// @Router       		/store/random [post]
func (a *Application) CreateRandomStores(gCtx *gin.Context) {
	quantity := ds.QuantityStores{}
	err := gCtx.BindJSON(&quantity)
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&swagger.Error{
				Description: "The quantity is negative or not int",
				Error:       swagger.Err400,
				Type:        swagger.TypeClientReq,
			})
		return
	}

	for i := 0; i < int(quantity.Quantity); i++ {
		err = a.repo.CreateRandomStores()
		if err != nil {
			gCtx.JSON(
				http.StatusInternalServerError,
				&swagger.Error{
					Description: "Create random promo failed",
					Error:       swagger.Err500,
					Type:        swagger.TypeInternalReq,
				})
			return
		}
	}

	gCtx.JSON(
		http.StatusCreated,
		&swagger.Create{
			Success: true,
		})
}

// ChangeStore		godoc
// @Summary      	Change promo price
// @Description  	Change the promo price using its uuid
// @Tags         	Change
// @Produce      	json
// @Param 			UUID path string true "UUID магазина" format(uuid)
// @Param 			Price body ds.PriceStore true "Новая цена"
// @Success      	200 {object} swagger.Change
// @Failure 		400 {object} swagger.Error
// @Failure 		404 {object} swagger.Error
// @Failure 	 	500 {object} swagger.Error
// @Router       	/store/{UUID} [put]
func (a *Application) ChangeStore(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&swagger.Error{
				Description: "Invalid UUID format",
				Error:       swagger.Err400,
				Type:        swagger.TypeClientReq,
			})
		return
	}

	store := ds.Store{}
	err = gCtx.BindJSON(&store)
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&swagger.Error{
				Description: "The price is negative or not int",
				Error:       swagger.Err400,
				Type:        swagger.TypeClientReq,
			})
		return
	}

	if len(store.Promo) != int(store.Quantity) {
		store.Promo = strings.Split(store.Promo[0], ",")
	}

	code, err := a.repo.ChangeStore(UUID, store)
	if err != nil {
		if code == 404 {
			gCtx.JSON(
				http.StatusNotFound,
				&swagger.Error{
					Description: "UUID Not Found",
					Error:       swagger.Err404,
					Type:        swagger.TypeClientReq,
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&swagger.Error{
					Description: "Change failed",
					Error:       swagger.Err500,
					Type:        swagger.TypeInternalReq,
				})
			return
		}
	}

	gCtx.JSON(
		http.StatusOK,
		&swagger.Change{
			Success: true,
		})
}

// DeleteStore		godoc
// @Summary     	Delete a store
// @Description 	Delete a store using its uuid
// @Tags         	Delete
// @Produce      	json
// @Param 			UUID path string true "UUID магазина" format(uuid)
// @Success      	200 {object} swagger.Delete
// @Failure 		400 {object} swagger.Error
// @Failure 		404 {object} swagger.Error
// @Failure 	 	500 {object} swagger.Error
// @Router       	/store/{UUID} [delete]
func (a *Application) DeleteStore(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&swagger.Error{
				Description: "Invalid UUID format",
				Error:       swagger.Err400,
				Type:        swagger.TypeClientReq,
			})
		return
	}

	code, err := a.repo.DeleteStore(UUID)
	if err != nil {
		if code == 404 {
			gCtx.JSON(
				http.StatusNotFound,
				&swagger.Error{
					Description: "UUID Not Found",
					Error:       swagger.Err404,
					Type:        swagger.TypeClientReq,
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&swagger.Error{
					Description: "Delete failed",
					Error:       swagger.Err500,
					Type:        swagger.TypeInternalReq,
				})
			return
		}
	}

	gCtx.JSON(
		http.StatusOK,
		&swagger.Delete{
			Success: true,
		})
}
