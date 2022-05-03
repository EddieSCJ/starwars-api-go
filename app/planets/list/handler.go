package list

import (
	"context"
	"net/http"
	"starwars-api-go/app/commons"
	"starwars-api-go/app/planets/model"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
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
	logger := log.Ctx(ctx.Request().Context())
	logger.Info().Msg("Starting List Planets")

	filter := model.Filter{}
	if err := ctx.Bind(&filter); err != nil {
		logger.Error().Err(err).Msg("Error binding filter")
		badRequestError := commons.NewBadRequest("wrong params")
		return ctx.JSON(badRequestError.Code, badRequestError)
	}

	planets, err := h.service.List(ctx.Request().Context(), filter)
	if err != nil {
		logger.Error().Err(err).Msg("Error listing planets")
		internalServerError := commons.NewInternalServerError(err.Error())
		return ctx.JSON(internalServerError.Code, internalServerError)
	}

	logger.Info().Msg("Finishing List Planets Successfully")
	return ctx.JSON(http.StatusOK, model.FromDomainList(planets))
}
