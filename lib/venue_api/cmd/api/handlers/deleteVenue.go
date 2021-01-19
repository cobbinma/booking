package handlers

import (
	"fmt"
	"github.com/cobbinma/booking-platform/lib/venue_api/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

func DeleteVenue(repository models.Repository) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		id, err := getVenueIDFromRequest(c)
		if err != nil {
			logrus.Error(fmt.Errorf("%s : %w", "could not get id from request", err))
			return c.JSON(http.StatusBadRequest, newErrorResponse(InvalidRequest, "invalid id"))
		}

		if err := repository.DeleteVenue(ctx, id); err != nil {
			m := "could not delete venue"
			logrus.Error(fmt.Errorf("%s : %w", m, err))
			return c.JSON(http.StatusInternalServerError, newErrorResponse(InternalError, m))
		}

		return c.NoContent(http.StatusNoContent)
	}
}
