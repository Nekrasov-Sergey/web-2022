package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"main/swagger"
	"net/http"
)

func (a *Application) GetCart(gCtx *gin.Context) {
	resp, err := a.repo.GetCart()
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

func (a *Application) DeleteCart(gCtx *gin.Context) {
	store, err := uuid.Parse(gCtx.Param("store"))
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

	code, err := a.repo.DeleteCart(store)
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

func (a *Application) GetStore(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	resp, err := a.repo.GetStore(UUID)
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

func (a *Application) IncreaseQuantity(gCtx *gin.Context) {
	store, err := uuid.Parse(gCtx.Param("store"))
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

	quantity, err := a.repo.IncreaseQuantity(store)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&swagger.StoreError{
				Description: "Change failed",
				Error:       swagger.Err500,
				Type:        swagger.TypeInternalReq,
			})
		return
	}

	gCtx.JSON(http.StatusOK, quantity)
}

func (a *Application) DecreaseQuantity(gCtx *gin.Context) {
	store, err := uuid.Parse(gCtx.Param("store"))
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

	quantity, code, err := a.repo.DecreaseQuantity(store)
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

	gCtx.JSON(http.StatusOK, quantity)
}
