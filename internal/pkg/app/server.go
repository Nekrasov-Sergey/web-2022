package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "main/docs"
	"main/internal/app/ds"
	"main/swagger"
	"net/http"
	"strconv"
)

func (a *Application) StartServer() {
	log.Println("server start up")

	r := gin.Default()

	r.Use(CORSMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/store/:sort", a.GetStores)

	r.GET("/store/price/:uuid", a.GetPriceStore)

	r.POST("/store", a.CreateStore)

	r.POST("/store/random", a.CreateRandomStores)

	r.PUT("/store/:uuid", a.ChangePriceStore)

	r.DELETE("/store/:uuid", a.DeleteStore)

	//Запросы для корзины:
	r.GET("/cart", a.GetCart)

	r.GET("/store/1/:uuid", a.GetStore)

	r.GET("/store/promo/:quantity/:uuid", a.GetPromoStore)

	r.GET("/cart/increase/:store", a.IncreaseQuantity)

	r.GET("/cart/decrease/:store", a.DecreaseQuantity)

	r.GET("/cart/delete/:store", a.DeleteCart)

	_ = r.Run()

	log.Println("server down")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// GetStores 		godoc
// @Summary      	Get all stores
// @Description  	Get a list of all stores
// @Tags         	Info
// @Produce      	json
// @Success      	200 {object} ds.StoreDocs
// @Failure 		500 {object} swagger.StoreError
// @Router       	/store [get]
func (a *Application) GetStores(gCtx *gin.Context) {
	sort := gCtx.Param("sort")
	resp, err := a.repo.GetStores(sort)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&swagger.StoreError{
				Description: "Can't get a list of promo codes",
				Error:       swagger.Err500,
				Type:        swagger.TypeInternalReq,
			})
		return
	}

	gCtx.JSON(http.StatusOK, resp)
}

// GetPriceStore  	godoc
// @Summary      	Get price of store
// @Description  	Get price of store by UUID
// @Tags         	Info
// @Produce      	json
// @Param 			UUID path string true "UUID промо" format(uuid)
// @Success      	200 {object} swagger.StorePrice
// @Failure 	 	400 {object} swagger.StoreError
// @Failure 	 	404 {object} swagger.StoreError
// @Failure 	 	500 {object} swagger.StoreError
// @Router       	/store/price/{UUID} [get]
func (a *Application) GetPriceStore(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&swagger.StoreError{
				Description: "Invalid UUID format",
				Error:       swagger.Err400,
				Type:        swagger.TypeClientReq,
			})
		return
	}

	var promo ds.Store
	code, err := a.repo.GetPriceStore(UUID, &promo)
	if err != nil {
		if code == 404 {
			gCtx.JSON(
				http.StatusNotFound,
				&swagger.StoreError{
					Description: "UUID Not Found",
					Error:       swagger.Err404,
					Type:        swagger.TypeClientReq,
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&swagger.StoreError{
					Description: "Get promo price failed",
					Error:       swagger.Err500,
					Type:        swagger.TypeInternalReq,
				})
			return
		}
	}

	gCtx.JSON(
		http.StatusOK,
		&swagger.StorePrice{
			Price: promo.Price,
		})
}

// GetPromoStore			godoc
// @Summary     	Get a promo
// @Description 	Get a promo in store using its uuid
// @Tags         	Info
// @Produce      	json
// @Param 			UUID path string true "UUID промо" format(uuid)
// @Success      	200 {object} swagger.StorePromo
// @Failure 		400 {object} swagger.StoreError
// @Failure 		404 {object} swagger.StoreError
// @Failure 	 	500 {object} swagger.StoreError
// @Router       	/store/promo/{UUID} [get]
func (a *Application) GetPromoStore(gCtx *gin.Context) {
	quantity, _ := strconv.Atoi(gCtx.Param("quantity"))
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&swagger.StoreError{
				Description: "Invalid UUID format",
				Error:       swagger.Err400,
				Type:        swagger.TypeClientReq,
			})
		return
	}

	code, Promo, err := a.repo.GetPromoStore(uint64(quantity), UUID)
	if err != nil {
		if code == 404 {
			gCtx.JSON(
				http.StatusNotFound,
				&swagger.StoreError{
					Description: "UUID Not Found",
					Error:       swagger.Err404,
					Type:        swagger.TypeClientReq,
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&swagger.StoreError{
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
// @Success 		201 {object} swagger.StoreCreated
// @Failure 		400 {object} swagger.StoreError
// @Failure 		500 {object} swagger.StoreError
// @Router  		/store [post]
func (a *Application) CreateStore(gCtx *gin.Context) {
	promo := ds.Store{}
	err := gCtx.BindJSON(&promo)
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&swagger.StoreError{
				Description: "Invalid parameters",
				Error:       swagger.Err400,
				Type:        swagger.TypeClientReq,
			})
		return
	}

	err = a.repo.CreateStore(promo)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&swagger.StoreError{
				Description: "Create failed",
				Error:       swagger.Err500,
				Type:        swagger.TypeInternalReq,
			})
		return
	}

	gCtx.JSON(
		http.StatusCreated,
		&swagger.StoreCreated{
			Success: true,
		})
}

// CreateRandomStores 	godoc
// @Summary      		Add a new random promo
// @Description  		Adding a new random promo to database
// @Tags        		Add
// @Produce      		json
// @Param				Quantity body ds.QuantityStores true "Количество"
// @Success     		201 {object} swagger.StoreCreated
// @Failure 			400 {object} swagger.StoreError
// @Failure 			500 {object} swagger.StoreError
// @Router       		/store/random [post]
func (a *Application) CreateRandomStores(gCtx *gin.Context) {
	quantity := ds.QuantityStores{}
	err := gCtx.BindJSON(&quantity)
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&swagger.StoreError{
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
				&swagger.StoreError{
					Description: "Create random promo failed",
					Error:       swagger.Err500,
					Type:        swagger.TypeInternalReq,
				})
			return
		}
	}

	gCtx.JSON(
		http.StatusCreated,
		&swagger.StoreCreated{
			Success: true,
		})
}

// ChangePriceStore		godoc
// @Summary      	Change promo price
// @Description  	Change the promo price using its uuid
// @Tags         	Change
// @Produce      	json
// @Param 			UUID path string true "UUID промо" format(uuid)
// @Param 			Price body ds.PriceStore true "Новая цена"
// @Success      	200 {object} swagger.StoreChanged
// @Failure 		400 {object} swagger.StoreError
// @Failure 		404 {object} swagger.StoreError
// @Failure 	 	500 {object} swagger.StoreError
// @Router       	/store/{UUID} [put]
func (a *Application) ChangePriceStore(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&swagger.StoreError{
				Description: "Invalid UUID format",
				Error:       swagger.Err400,
				Type:        swagger.TypeClientReq,
			})
		return
	}

	promo := ds.Store{}
	err = gCtx.BindJSON(&promo)
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&swagger.StoreError{
				Description: "The price is negative or not int",
				Error:       swagger.Err400,
				Type:        swagger.TypeClientReq,
			})
		return
	}

	code, err := a.repo.ChangePriceStore(UUID, promo.Price)
	if err != nil {
		if code == 404 {
			gCtx.JSON(
				http.StatusNotFound,
				&swagger.StoreError{
					Description: "UUID Not Found",
					Error:       swagger.Err404,
					Type:        swagger.TypeClientReq,
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&swagger.StoreError{
					Description: "Change failed",
					Error:       swagger.Err500,
					Type:        swagger.TypeInternalReq,
				})
			return
		}
	}

	gCtx.JSON(
		http.StatusOK,
		&swagger.StoreChanged{
			Success: true,
		})
}

// DeleteStore		godoc
// @Summary     	Delete a store
// @Description 	Delete a store using its uuid
// @Tags         	Delete
// @Produce      	json
// @Param 			UUID path string true "UUID промо" format(uuid)
// @Success      	200 {object} swagger.StoreDeleted
// @Failure 		400 {object} swagger.StoreError
// @Failure 		404 {object} swagger.StoreError
// @Failure 	 	500 {object} swagger.StoreError
// @Router       	/store/{UUID} [delete]
func (a *Application) DeleteStore(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&swagger.StoreError{
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
				&swagger.StoreError{
					Description: "UUID Not Found",
					Error:       swagger.Err404,
					Type:        swagger.TypeClientReq,
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&swagger.StoreError{
					Description: "Delete failed",
					Error:       swagger.Err500,
					Type:        swagger.TypeInternalReq,
				})
			return
		}
	}

	gCtx.JSON(
		http.StatusOK,
		&swagger.StoreDeleted{
			Success: true,
		})
}
