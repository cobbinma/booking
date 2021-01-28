package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/cobbinma/booking-platform/lib/protobuf/autogen/lang/go/venue/api"
	"github.com/cobbinma/booking-platform/lib/protobuf/autogen/lang/go/venue/models"
	"github.com/cobbinma/booking-platform/lib/venue_api/cmd/api/middleware"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"os"
)

func main() {
	_ = godotenv.Load()

	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("could not start logger : %s", err)
		os.Exit(-1)
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Error("could not sync logger")
		}
	}()
	log := logger.Sugar()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	ensureValidToken, err := middleware.EnsureValidToken()
	if err != nil {
		log.Fatalf("could not construct auth server inceptor : %s", err)
	}

	cert, err := tls.LoadX509KeyPair("localhost.crt", "localhost.key")
	if err != nil {
		log.Fatalf("failed to load cert : %s", err)
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("could not listen : %s", err)
	}

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		grpc_middleware.WithUnaryServerChain(grpc_zap.UnaryServerInterceptor(logger), ensureValidToken),
	}

	s := grpc.NewServer(opts...)
	api.RegisterVenueAPIServer(s, &VenueService{})
	api.RegisterTableAPIServer(s, &TableService{})

	log.Infof("starting gRPC listener on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %s", err)
	}
}

var _ api.VenueAPIServer = (*VenueService)(nil)

type VenueService struct{}

func (s VenueService) GetVenue(ctx context.Context, request *api.GetVenueRequest) (*models.Venue, error) {
	hours := []*models.OpeningHoursSpecification{
		{
			DayOfWeek: 1,
			Opens:     "10:00",
			Closes:    "19:00",
		},
		{
			DayOfWeek: 2,
			Opens:     "10:00",
			Closes:    "19:00",
		},
		{
			DayOfWeek: 3,
			Opens:     "10:00",
			Closes:    "19:00",
		},
		{
			DayOfWeek: 4,
			Opens:     "10:00",
			Closes:    "19:00",
		},
		{
			DayOfWeek: 5,
			Opens:     "10:00",
			Closes:    "19:00",
		},
		{
			DayOfWeek: 6,
			Opens:     "10:00",
			Closes:    "23:00",
		},
		{
			DayOfWeek: 7,
			Opens:     "10:00",
			Closes:    "20:00",
		},
	}
	return &models.Venue{
		Id:                  request.Id,
		Name:                "Hop and Vine",
		OpeningHours:        hours,
		SpecialOpeningHours: nil,
	}, nil
}

var _ api.TableAPIServer = (*TableService)(nil)

type TableService struct{}

func (t TableService) GetTables(ctx context.Context, request *api.GetTablesRequest) (*api.GetTablesResponse, error) {
	panic("implement me")
}

func (t TableService) AddTable(ctx context.Context, request *api.AddTableRequest) (*models.Table, error) {
	panic("implement me")
}
