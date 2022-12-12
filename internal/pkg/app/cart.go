package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"main/swagger"
	"net/http"
)

// GetCart 			godoc
// @Summary      	Get a whole cart
// @Description  	Get a list of the entire basket
// @Tags         	Info
// @Produce      	json
// @Success      	200 {object} ds.Cart
// @Failure 		500 {object} swagger.Error
// @Router       	/cart [get]
func (a *Application) GetCart(gCtx *gin.Context) {
	jwtStr := gCtx.GetHeader("Authorization")
	userUUID := a.GetUserByToken(jwtStr)

	resp, err := a.repo.GetCart(userUUID)
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

// GetCart1 		godoc
// @Summary      	Get store from the cart
// @Description  	Get one store from the shopping cart
// @Tags         	Info
// @Produce      	json
// @Param 			Store path string true "Магазин"
// @Success      	200 {object} ds.Cart
// @Failure 		400 {object} swagger.Error
// @Failure 		500 {object} swagger.Error
// @Router       	/cart/{Store} [get]
func (a *Application) GetCart1(gCtx *gin.Context) {
	jwtStr := gCtx.GetHeader("Authorization")
	userUUID := a.GetUserByToken(jwtStr)

	store, err := uuid.Parse(gCtx.Param("store"))
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

	resp, _ := a.repo.GetCart1(store, userUUID)

	gCtx.JSON(http.StatusOK, resp)
}

// IncreaseQuantity godoc
// @Summary      	Increase by 1 in the cart
// @Description  	Increase by 1 the number of promo codes in the cart
// @Tags         	Info
// @Produce      	json
// @Param 			Store path string true "Магазин"
// @Success      	200 {object} swagger.CartIncrease
// @Failure 		400 {object} swagger.Error
// @Failure 		500 {object} swagger.Error
// @Router       	/cart/increase/{Store} [get]
func (a *Application) IncreaseQuantity(gCtx *gin.Context) {
	jwtStr := gCtx.GetHeader("Authorization")
	userUUID := a.GetUserByToken(jwtStr)

	store, err := uuid.Parse(gCtx.Param("store"))
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

	quantity, err := a.repo.IncreaseQuantity(store, userUUID)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&swagger.Error{
				Description: "Change failed",
				Error:       swagger.Err500,
				Type:        swagger.TypeInternalReq,
			})
		return
	}

	gCtx.JSON(http.StatusOK, quantity)
}

// DecreaseQuantity godoc
// @Summary      	Decrease by 1 in the cart
// @Description  	Decrease by 1 the number of promo codes in the cart
// @Tags         	Info
// @Produce      	json
// @Param 			Store path string true "Магазин"
// @Success      	200 {object} swagger.CartDecrease
// @Failure 		400 {object} swagger.Error
// @Failure 		404 {object} swagger.Error
// @Failure 		500 {object} swagger.Error
// @Router       	/cart/decrease/{Store} [get]
func (a *Application) DecreaseQuantity(gCtx *gin.Context) {
	jwtStr := gCtx.GetHeader("Authorization")
	userUUID := a.GetUserByToken(jwtStr)

	store, err := uuid.Parse(gCtx.Param("store"))
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

	quantity, code, err := a.repo.DecreaseQuantity(store, userUUID)
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

	gCtx.JSON(http.StatusOK, quantity)
}

// DeleteCart		godoc
// @Summary     	Delete a store in the cart
// @Description 	Delete a store in the cart using its uuid
// @Tags         	Delete
// @Produce      	json
// @Param 			Store path string true "Магазин"
// @Success      	200 {object} swagger.Delete
// @Failure 		400 {object} swagger.Error
// @Failure 		404 {object} swagger.Error
// @Failure 	 	500 {object} swagger.Error
// @Router       	/cart/delete/{Store} [delete]
func (a *Application) DeleteCart(gCtx *gin.Context) {
	jwtStr := gCtx.GetHeader("Authorization")
	userUUID := a.GetUserByToken(jwtStr)

	store, err := uuid.Parse(gCtx.Param("store"))
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

	code, err := a.repo.DeleteCart(store, userUUID)
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
