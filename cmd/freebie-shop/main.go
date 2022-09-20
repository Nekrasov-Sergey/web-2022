package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"main/internal/app/config"
	"main/internal/pkg/app"
	"os"
)

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
}
