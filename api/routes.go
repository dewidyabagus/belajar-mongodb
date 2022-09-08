package api

import (
	"belajar-mongodb/api/v1/consumer"

	echo "github.com/labstack/echo/v4"
)

type Routes struct {
	Consumer *consumer.Controller
}

func NewRoutes(e *echo.Echo, routes *Routes) {
	v1 := e.Group("/v1")

	v1.POST("/consumers", routes.Consumer.InsertOne)
	v1.GET("/consumers", routes.Consumer.GetAll)
	v1.GET("/consumers/id/:id", routes.Consumer.FindByID)
	v1.PUT("/consumers/id/:id", routes.Consumer.UpdateByID)
	v1.DELETE("/consumers/id/:id", routes.Consumer.DeleteByID)
}
