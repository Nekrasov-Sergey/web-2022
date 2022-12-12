package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"main/internal/app/ds"
	"main/swagger"
	"net/http"
)

//func (a *Application) AddOrder(gCtx *gin.Context) {
//	jwtStr := gCtx.GetHeader("Authorization")
//	userUUID := a.GetUserByToken(jwtStr)
//	order := ds.Order{}
//	order.UserUUID = userUUID
//	err := gCtx.BindJSON(&order)
//	if err != nil {
//		gCtx.JSON(
//			http.StatusBadRequest,
//			&swagger.Error{
//				Description: "Invalid parameters",
//				Error:       swagger.Err400,
//				Type:        swagger.TypeClientReq,
//			})
//		return
//	}
//	err = a.repo.AddOrder(order)
//	if err != nil {
//		gCtx.JSON(
//			http.StatusInternalServerError,
//			&swagger.Error{
//				Description: "Create failed",
//				Error:       swagger.Err500,
//				Type:        swagger.TypeInternalReq,
//			})
//		return
//	}
//	gCtx.JSON(
//		http.StatusOK,
//		&swagger.Create{
//			Success: true,
//		})
//
//}

func (a *Application) GetOrders(gCtx *gin.Context) {
	stDate := gCtx.Query("start_date")
	endDate := gCtx.Query("end_date")
	status := gCtx.Query("status")
	resp, err := a.repo.GetOrders(stDate, endDate, status)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&swagger.Error{
				Description: "can`t get a list",
				Error:       swagger.Err500,
				Type:        swagger.TypeInternalReq,
			})
		return
	}
	gCtx.JSON(http.StatusOK, resp)
}

func (a *Application) ChangeStatus(gCtx *gin.Context) {
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
	order := ds.Order{}
	err = gCtx.BindJSON(&order)
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
	resp, err := a.repo.ChangeStatus(UUID, order.Status)
	if err != nil {
		if resp == 404 {
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
		&swagger.Create{
			Success: true,
		})

}
