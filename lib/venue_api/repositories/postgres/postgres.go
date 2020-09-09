package postgres

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/cobbinma/booking/lib/venue_api/models"
)

var _ models.Repository = (*postgres)(nil)

type postgres struct {
	dbClient DBClient
}

func NewPostgres(client DBClient) models.Repository {
	return &postgres{dbClient: client}
}

func (p *postgres) CreateVenue(ctx context.Context, venue models.VenueInput) error {
	tx, err := p.dbClient.BeginX()
	if err != nil {
		return fmt.Errorf("%s : %w", "could not begin transaction", err)
	}

	rows, err := tx.NamedQuery("INSERT INTO venues (name) VALUES (:name) RETURNING id", venue)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("%s : %w", "could not perform named query to insert into venues", err)
	}

	var venueID int
	if rows.Next() {
		if err := rows.Scan(&venueID); err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("%s : %w", "could not scan row", err)
		}
	}
	if err := rows.Close(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("%s : %w", "could not close rows", err)
	}

	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert("opening_hours").
		Columns("venue_id", "day_of_week", "opens", "closes")

	for i := range venue.OpeningHours {
		oh := venue.OpeningHours[i]
		builder = builder.Values(venueID, oh.DayOfWeek, oh.Opens.Time(), oh.Closes.Time())
	}
	builder = builder.Suffix("ON CONFLICT (venue_id, day_of_week) DO UPDATE SET opens = EXCLUDED.opens, closes = EXCLUDED.closes")

	sql, args, err := builder.ToSql()
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("%s : %w", "could not build sql statement", err)
	}
	_, err = tx.Exec(sql, args...)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("%s : %w", "could not execute sql statement", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("%s : %w", "could not commit transaction", err)
	}

	return nil
}
