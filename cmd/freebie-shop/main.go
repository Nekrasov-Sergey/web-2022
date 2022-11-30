package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"main/internal/pkg/app"
	"os"
)

// @title Freebie-shop
// @version 1.0
// @description Store with promo codes for various stores

// @contact.name Sergey Nekrasov
// @contact.url https://vk.com/serega_nekrasov
// @contact.email 79508031750@yandex.ru

// @host 127.0.0.1:8080
// @schemes http https
// @BasePath /

func main() {
	log.Println("application start")

	ctx := context.Background()

	application, err := app.New(ctx)
	if err != nil {
		log.Printf("can`t create application: %s", err)
		os.Exit(2)
	}

	err = application.Run()
	if err != nil {
		log.Printf("can`t run application: %s", err)
		os.Exit(2)
	}

	log.Println("application terminated")
}
