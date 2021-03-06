package main

import (
	"fmt"
	mw "github.com/cobbinma/booking-platform/lib/gateway_api/cmd/api/middleware"
	"github.com/cobbinma/booking-platform/lib/gateway_api/internal/auth0"
	"github.com/cobbinma/booking-platform/lib/gateway_api/internal/booking"
	"github.com/cobbinma/booking-platform/lib/gateway_api/internal/venue"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cobbinma/booking-platform/lib/gateway_api/graph"
	"github.com/cobbinma/booking-platform/lib/gateway_api/graph/generated"
)

const defaultPort = "9999"

func main() {
	_ = godotenv.Load()

	logger, err := zap.NewProduction()
	if err != nil {
		panic("could not start logger" + err.Error())
	}
	log := logger.Sugar()

	c, err := NewConfig()
	if err != nil {
		log.Fatalf("could not construct config : %s", err)
	}

	tokenClient, err := auth0.NewTokenClient(log, c.authDomain)
	if err != nil {
		log.Fatalf("could not create token client : %s", err)
	}

	venueToken, err := tokenClient.GetToken(log, "http://venue")
	if err != nil {
		log.Fatalf("could not get venue client : %s", err)
	}

	venueClient, closeVenueClient, err := venue.NewVenueClient(c.venueURL, log, venueToken)
	if err != nil {
		log.Fatalf("could not create venue client : %s", err)
	}
	defer closeVenueClient(log)

	bookingToken, err := tokenClient.GetToken(log, "http://booking")
	if err != nil {
		log.Fatalf("could not get booking client : %s", err)
	}

	bookingClient, closeBookingClient, err := booking.NewBookingClient(c.bookingURL, log, bookingToken)
	if err != nil {
		log.Fatalf("could not create booking client : %s", err)
	}
	defer closeBookingClient(log)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{Resolvers: graph.NewResolver(log, venueClient, bookingClient)}))
	e := echo.New()
	e.Use(mw.ZapLogger(logger))

	if c.allowCors {
		e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	}

	e.GET("/", echo.WrapHandler(playground.Handler("GraphQL playground", "/query")))
	e.POST("/query", echo.WrapHandler(srv), mw.Auth(c.authDomain, c.authApiId), mw.User(auth0.NewUserService(c.authDomain)))
	e.OPTIONS("/query", func(c echo.Context) error {
		headers := c.Request().Header
		for key, value := range headers {
			c.Response().Header().Set(key, value[0])
		}
		return c.NoContent(http.StatusOK)
	})

	log.Infof("connect to http://localhost:%s/ for GraphQL playground", c.port)
	e.Logger.Fatal(e.Start(":" + c.port))
}

type Config struct {
	port                string
	allowCors           bool
	authDomain          string
	authApiId           string
	venueURL            string
	authVenueAudience   string
	bookingURL          string
	authBookingAudience string
}

func NewConfig() (*Config, error) {
	missing := []string{}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	allowCors := false
	if _, present := os.LookupEnv("ALLOW_CORS"); present {
		allowCors = true
	}
	e := "AUTH0_DOMAIN"
	domain, present := os.LookupEnv(e)
	if !present {
		missing = append(missing, e)
	}
	e = "AUTH0_API_IDENTIFIER"
	apiId, present := os.LookupEnv(e)
	if !present {
		missing = append(missing, e)
	}
	e = "VENUE_API_ROOT"
	venueURL, present := os.LookupEnv(e)
	if !present {
		missing = append(missing, e)
	}
	e = "AUTH0_VENUE_API_IDENTIFIER"
	authVenueAudience, present := os.LookupEnv(e)
	if !present {
		missing = append(missing, e)
	}
	e = "AUTH0_BOOKING_API_IDENTIFIER"
	authBookingAudience, present := os.LookupEnv(e)
	if !present {
		missing = append(missing, e)
	}
	e = "BOOKING_API_ROOT"
	bookingURL, present := os.LookupEnv(e)
	if !present {
		missing = append(missing, e)
	}

	if len(missing) > 0 {
		return nil, fmt.Errorf("missing environment variables [%s]", strings.Join(missing, ", "))
	}

	return &Config{
		port:                port,
		allowCors:           allowCors,
		authDomain:          domain,
		authApiId:           apiId,
		venueURL:            venueURL,
		authVenueAudience:   authVenueAudience,
		bookingURL:          bookingURL,
		authBookingAudience: authBookingAudience,
	}, nil
}
