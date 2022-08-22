package main

import (
	"context"
	"fmt"
	"log"
	"time"

	consumerBusiness "belajar-mongodb/business/consumer"
	"belajar-mongodb/config"
	consumerRepository "belajar-mongodb/modules/consumer"
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

	// consumerService.InsertOne(&consumerBusiness.Consumer{
	// 	ID:      1,
	// 	Name:    "Widya Bagus",
	// 	Address: "Bondowoso",
	// })

	result, err := consumerService.FindByID(1)
	if err != nil {
		panic("error: " + err.Error())
	}
	fmt.Printf("Result: %+v \n", result)
}
