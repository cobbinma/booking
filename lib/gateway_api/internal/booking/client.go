package booking

import (
	"context"
	"fmt"
	"github.com/cobbinma/booking-platform/lib/gateway_api/graph"
	"github.com/cobbinma/booking-platform/lib/gateway_api/models"
	"github.com/cobbinma/booking-platform/lib/protobuf/autogen/lang/go/booking/api"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"time"
)

func NewBookingClient(url string, log *zap.SugaredLogger, token *oauth2.Token, options ...func(*bookingClient)) (graph.BookingService, func(log *zap.SugaredLogger), error) {
	bc := &bookingClient{
		client: nil,
		log:    log,
	}
	cl := func(log *zap.SugaredLogger) {}

	for i := range options {
		options[i](bc)
	}

	if bc.client == nil {
		creds, err := credentials.NewClientTLSFromFile("localhost.crt", "localhost")
		if err != nil {
			return nil, nil, fmt.Errorf("failed to load credentials : %w", err)
		}

		opts := []grpc.DialOption{
			grpc.WithPerRPCCredentials(oauth.NewOauthAccess(token)),
			grpc.WithTransportCredentials(creds),
		}
		conn, err := grpc.Dial(url, opts...)
		if err != nil {
			return nil, nil, fmt.Errorf("could not connect : %s", err)
		}

		cl = func(log *zap.SugaredLogger) {
			if err := conn.Close(); err != nil {
				log.Error("could not close connection : %s", err)
			}
		}
		bc.client = api.NewBookingAPIClient(conn)
	}

	return bc, cl, nil
}

func WithClient(client api.BookingAPIClient) func(*bookingClient) {
	return func(c *bookingClient) {
		c.client = client
	}
}

type bookingClient struct {
	client api.BookingAPIClient
	log    *zap.SugaredLogger
}

func (b bookingClient) CancelBooking(ctx context.Context, input models.CancelBookingInput) (*models.Booking, error) {
	cancelled, err := b.client.CancelBooking(ctx, &api.CancelBookingRequest{Id: input.ID})
	if err != nil {
		return nil, fmt.Errorf("could not cancel booking using client : %w", err)
	}

	startsAt, err := time.Parse(time.RFC3339, cancelled.StartsAt)
	if err != nil {
		return nil, fmt.Errorf("could not parse start time : %w", err)
	}

	endsAt, err := time.Parse(time.RFC3339, cancelled.EndsAt)
	if err != nil {
		return nil, fmt.Errorf("could not parse end time : %w", err)
	}

	var givenName, familyName *string
	if cancelled.GivenName != "" {
		givenName = &cancelled.GivenName
	}
	if cancelled.FamilyName != "" {
		familyName = &cancelled.FamilyName
	}

	return &models.Booking{
		ID:         cancelled.Id,
		VenueID:    cancelled.VenueId,
		Email:      cancelled.Email,
		People:     int(cancelled.People),
		StartsAt:   startsAt,
		EndsAt:     endsAt,
		Duration:   int(cancelled.Duration),
		TableID:    cancelled.TableId,
		GivenName:  givenName,
		FamilyName: familyName,
	}, err
}

func (b bookingClient) Bookings(ctx context.Context, filter models.BookingsFilter, pageInfo models.PageInfo) (*models.BookingsPage, error) {
	if pageInfo.Limit == nil {
		var max = 50
		pageInfo.Limit = &max
	}
	if filter.VenueID == nil {
		empty := ""
		filter.VenueID = &empty
	}
	resp, err := b.client.GetBookings(ctx, &api.GetBookingsRequest{
		VenueId: *filter.VenueID,
		Date:    filter.Date.Format(time.RFC3339),
		Page:    int32(pageInfo.Page),
		Limit:   int32(*pageInfo.Limit),
	})
	if err != nil {
		return nil, fmt.Errorf("could not get bookings from client : %w", err)
	}

	bookings := make([]*models.Booking, len(resp.Bookings))
	for i, b := range resp.Bookings {
		startsAt, err := time.Parse(time.RFC3339, b.StartsAt)
		if err != nil {
			return nil, fmt.Errorf("incorrect time format returned from booking client : %w", err)
		}
		endsAt, err := time.Parse(time.RFC3339, b.EndsAt)
		if err != nil {
			return nil, fmt.Errorf("incorrect time format returned from booking client : %w", err)
		}
		var givenName, familyName *string
		if b.GivenName != "" {
			givenName = &b.GivenName
		}
		if b.FamilyName != "" {
			familyName = &b.FamilyName
		}
		bookings[i] = &models.Booking{
			ID:         b.Id,
			VenueID:    b.VenueId,
			Email:      b.Email,
			People:     int(b.People),
			StartsAt:   startsAt,
			EndsAt:     endsAt,
			Duration:   int(b.Duration),
			TableID:    b.TableId,
			GivenName:  givenName,
			FamilyName: familyName,
		}
	}

	return &models.BookingsPage{
		Bookings:    bookings,
		HasNextPage: resp.HasNextPage,
		Pages:       int(resp.Pages),
	}, nil
}

func (b bookingClient) GetSlot(ctx context.Context, slot models.SlotInput) (*models.GetSlotResponse, error) {
	resp, err := b.client.GetSlot(ctx, &api.SlotInput{
		VenueId:  slot.VenueID,
		Email:    slot.Email,
		People:   (uint32)(slot.People),
		StartsAt: slot.StartsAt.Format(time.RFC3339),
		Duration: (uint32)(slot.Duration),
	})
	if err != nil {
		return nil, fmt.Errorf("could not get slot from booking api : %w", err)
	}

	var match *models.Slot
	if resp.Match != nil {
		startsAt, err := time.Parse(time.RFC3339, resp.Match.StartsAt)
		if err != nil {
			return nil, fmt.Errorf("could not parse start time : %w", err)
		}

		endsAt, err := time.Parse(time.RFC3339, resp.Match.EndsAt)
		if err != nil {
			return nil, fmt.Errorf("could not parse end time : %w", err)
		}

		match = &models.Slot{
			VenueID:  resp.Match.VenueId,
			Email:    resp.Match.Email,
			People:   (int)(resp.Match.People),
			StartsAt: startsAt,
			EndsAt:   endsAt,
			Duration: (int)(resp.Match.Duration),
		}
	}

	return &models.GetSlotResponse{
		Match:               match,
		OtherAvailableSlots: nil,
	}, nil
}

func (b bookingClient) CreateBooking(ctx context.Context, input models.BookingInput) (*models.Booking, error) {
	var givenName, familyName string
	if input.GivenName != nil {
		givenName = *input.GivenName
	}
	if input.FamilyName != nil {
		familyName = *input.FamilyName
	}

	resp, err := b.client.CreateBooking(ctx, &api.BookingInput{
		VenueId:    input.VenueID,
		Email:      input.Email,
		People:     (uint32)(input.People),
		StartsAt:   input.StartsAt.Format(time.RFC3339),
		Duration:   (uint32)(input.Duration),
		GivenName:  givenName,
		FamilyName: familyName,
	})
	if err != nil {
		return nil, fmt.Errorf("could not create booking in booking api : %w", err)
	}

	startsAt, err := time.Parse(time.RFC3339, resp.StartsAt)
	if err != nil {
		return nil, fmt.Errorf("could not parse start time : %w", err)
	}

	endsAt, err := time.Parse(time.RFC3339, resp.EndsAt)
	if err != nil {
		return nil, fmt.Errorf("could not parse end time : %w", err)
	}

	var gn, fn *string
	if resp.GivenName != "" {
		gn = &resp.GivenName
	}
	if resp.FamilyName != "" {
		fn = &resp.FamilyName
	}

	return &models.Booking{
		ID:         resp.Id,
		VenueID:    resp.VenueId,
		Email:      resp.Email,
		People:     (int)(resp.People),
		StartsAt:   startsAt,
		EndsAt:     endsAt,
		Duration:   (int)(resp.Duration),
		TableID:    resp.TableId,
		GivenName:  gn,
		FamilyName: fn,
	}, nil
}
