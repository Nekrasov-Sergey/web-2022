package app

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "main/docs"
	"main/internal/app/role"
)

func (a *Application) StartServer() {
	log.Println("server start up")

	r := gin.Default()

	r.Use(CORSMiddleware())

	// Запрос для свагера
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запросы для магазина:
	r.GET("/store", a.GetStores)

	r.GET("/store/:uuid", a.GetStore)

	r.GET("/store/:uuid/:quantity", a.GetPromoStore)

	r.POST("/store", a.CreateStore)

	r.POST("/store/random", a.CreateRandomStores)

	r.PUT("/store/:uuid", a.ChangePriceStore)

	r.DELETE("/store/:uuid", a.DeleteStore)

	// Запросы для корзины:
	r.GET("/cart", a.GetCart)

	r.GET("/cart/:store", a.GetCart1)

	r.GET("/cart/increase/:store", a.IncreaseQuantity)

	r.GET("/cart/decrease/:store", a.DecreaseQuantity)

	r.DELETE("/cart/delete/:store", a.DeleteCart)

	// Запросы для авторизации
	r.POST("/login", a.Login)

	r.POST("/sign_up", a.Register)

	r.GET("/logout", a.Logout)

	r.Use(a.WithAuthCheck(role.Manager, role.Admin, role.Buyer)).GET("/ping", a.Ping)

	_ = r.Run()

	log.Println("server down")
}
