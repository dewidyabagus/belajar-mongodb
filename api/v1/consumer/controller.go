package consumer

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"belajar-mongodb/api/common"
	"belajar-mongodb/api/v1/consumer/request"
	"belajar-mongodb/api/v1/consumer/response"
	"belajar-mongodb/business/consumer"
)

const (
	invalidObjId     = "invalid consumer id"
	malformedPayload = "malformed request syntax"
)

type Controller struct {
	Service consumer.Servicer
}

func NewController(service consumer.Servicer) *Controller {
	return &Controller{service}
}

func (c *Controller) InsertOne(ctx echo.Context) error {
	payload := new(request.NewConsumer)
	if err := ctx.Bind(payload); err != nil {
		return common.ResErrorController(ctx, http.StatusBadRequest, malformedPayload)
	}

	if err := c.Service.InsertOne(ctx.Request().Context(), payload.ToBusinessConsumer()); err != nil {
		return common.ResBusinessError(ctx, err)
	}

	return common.ResSuccessCreated(ctx, "success new consumer")
}

func (c *Controller) FindByID(ctx echo.Context) error {
	id := ctx.Param("id")

	if !primitive.IsValidObjectID(id) {
		return common.ResErrorController(ctx, http.StatusBadRequest, invalidObjId)
	}

	result, err := c.Service.FindByID(ctx.Request().Context(), id)
	if err != nil {
		return common.ResBusinessError(ctx, err)
	}

	return common.ResSuccessWithData(ctx, response.ResSingleConsumer(result))
}

func (c *Controller) GetAll(ctx echo.Context) error {
	results, err := c.Service.GetAll(ctx.Request().Context())
	if err != nil {
		return common.ResBusinessError(ctx, err)
	}

	return common.ResSuccessWithData(ctx, response.ResManyConsumer(results))
}

func (c *Controller) UpdateByID(ctx echo.Context) error {
	id := ctx.Param("id")
	if !primitive.IsValidObjectID(id) {
		return common.ResErrorController(ctx, http.StatusBadRequest, invalidObjId)
	}

	payload := new(request.UpConsumer)
	if err := ctx.Bind(payload); err != nil {
		return common.ResErrorController(ctx, http.StatusBadRequest, malformedPayload)
	}

	if err := c.Service.UpdateByID(ctx.Request().Context(), id, payload.ToBusinessConsumer()); err != nil {
		return common.ResBusinessError(ctx, err)
	}

	return common.ResSuccessOK(ctx, "success update data")
}

func (c *Controller) DeleteByID(ctx echo.Context) error {
	id := ctx.Param("id")
	if !primitive.IsValidObjectID(id) {
		return common.ResErrorController(ctx, http.StatusBadRequest, invalidObjId)
	}

	if err := c.Service.DeleteByID(ctx.Request().Context(), id); err != nil {
		return common.ResBusinessError(ctx, err)
	}

	return common.ResSuccessOK(ctx, "success delete data")
}
