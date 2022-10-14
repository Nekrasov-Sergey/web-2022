package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"main/internal/app/config"
	"main/internal/pkg/app"
	"os"
)

// @title Freebie-shop
// @version 1.0
// @description Store with promo codes for various stores

// @contact.name Sergey Nekrasov
// @contact.url https://vk.com/serega_nekrasov
// @contact.email 79508031750@yandex.ru

// @license.name AS IS (NO WARRANTY)

// @host 127.0.0.1:8080
// @schemes http https
// @BasePath /

func main() {
	log.Println("app start")

	ctx := context.Background()

	_, err := config.NewConfig(ctx) //Config пока не используется
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("can't init config")
		os.Exit(2)
	}

	application, err := app.New(ctx)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("can`t create app")
		os.Exit(2)
	}

	err = application.Run(ctx)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("can`t run app")
		os.Exit(2)
	}

	log.Println("app terminated")
}
