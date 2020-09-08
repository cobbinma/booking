package main

import (
	"context"
	"github.com/cobbinma/booking/lib/booking_api/cmd/api/handlers"
	"github.com/cobbinma/booking/lib/booking_api/config"
	"github.com/cobbinma/booking/lib/booking_api/gateways/tableAPI"
	"github.com/cobbinma/booking/lib/booking_api/repositories/postgres"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	dbClient, closeDB, err := postgres.NewDBClient()
	if err != nil {
		log.Fatal("could not create database client : ", err)
	}
	defer func() {
		if err := closeDB(); err != nil {
			log.Error("could not close database : ", err)
		}
	}()

	repository := postgres.NewPostgres(dbClient)
	if err := repository.Migrate(context.Background()); err != nil {
		log.Fatal("could not migrate : ", err)
	}

	tableClient := tableAPI.NewTableAPI()

	e := echo.New()

	if allowedOrigin := config.GetAllowOrigin(); allowedOrigin != "" {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{allowedOrigin},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		}))
	}

	e.Use(middleware.Logger())

	mw := handlers.VenueMiddleware

	h := handlers.NewHandlers(repository, tableClient)

	e.GET("/healthz", h.Health)
	e.POST("/venues/:venue_id/bookings", mw(h.CreateBooking))
	e.POST("/venues/:venue_id/slot", mw(h.BookingQuery))
	e.DELETE("/venues/:venue_id/bookings/:id", mw(h.DeleteBooking))
	e.GET("/venues/:venue_id/bookings/date/:date", mw(h.GetBookingsByDate))

	e.Logger.Fatal(e.Start(config.Port()))
}