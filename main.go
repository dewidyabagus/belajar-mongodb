package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"belajar-mongodb/api"
	consumerController "belajar-mongodb/api/v1/consumer"
	consumerBusiness "belajar-mongodb/business/consumer"
	"belajar-mongodb/config"
	consumerRepository "belajar-mongodb/modules/consumer"

	echo "github.com/labstack/echo/v4"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	app := config.LoadConfig()
	emptyCtx := context.Background()

	mongoClient, err := config.MongoConnect(emptyCtx, app.MongoConfig)
	if err != nil {
		panic(err.Error())
	}
	defer func() {
		ctxWT, cancel := context.WithTimeout(emptyCtx, time.Second*10)
		defer cancel()

		mongoClient.Client().Disconnect(ctxWT)
	}()

	consumerModule := consumerRepository.NewRepository(mongoClient)
	consumerService := consumerBusiness.NewService(consumerModule)
	consumerHandler := consumerController.NewController(consumerService)

	// Echo Framework Setup
	e := echo.New()

	routes := &api.Routes{
		Consumer: consumerHandler,
	}
	api.NewRoutes(e, routes)

	go func() {
		if err := e.Start(fmt.Sprintf("%s:%d", app.AppHost, app.ListenPort)); err != nil {
			log.Println("HTTP Service Shutdown now...")

			os.Exit(0)
		}
	}()

	quit := make(chan os.Signal, 10)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctxWT, cancel := context.WithTimeout(emptyCtx, time.Second*10)
	defer cancel()

	if err := e.Shutdown(ctxWT); err != nil {
		log.Println("Failed gracefully shutdown service")
	}
}
