package list

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"starwars-api-go/app/commons"
	"starwars-api-go/app/planets/model"
)

type PlanetService interface {
	List(ctx context.Context, filter model.Filter) ([]model.Planet, error)
}

type Handler struct {
	service PlanetService
}

func NewHandler(service PlanetService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) List(ctx echo.Context) error {
	filter := model.Filter{}
	if err := ctx.Bind(&filter); err != nil {
		badRequestError := commons.NewBadRequest("wrong params")
		return ctx.JSON(badRequestError.Code, badRequestError)
	}

	planets, err := h.service.List(ctx.Request().Context(), filter)
	if err != nil {
		internalServerError := commons.NewInternalServerError(err.Error())
		return ctx.JSON(internalServerError.Code, internalServerError)
	}

	return ctx.JSON(http.StatusOK, model.FromDomainList(planets))
}
