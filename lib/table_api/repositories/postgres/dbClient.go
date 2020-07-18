package postgres

import (
	"database/sql"
	"fmt"
	"github.com/cobbinma/booking/lib/table_api/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBClient interface {
	DB() *sql.DB
}

type dbClient struct {
	db *sqlx.DB
}

func NewDBClient() (*dbClient, func() error, error) {
	dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=%s",
		config.DBHost,
		config.DBName,
		config.DBUser,
		config.DBPassword,
		config.DBSSLMode)

	driver := "postgres"

	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		return nil, nil, fmt.Errorf("could not open database : %w", err)
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)

	dbc := &dbClient{db: db}

	return dbc, dbc.Close, nil
}

func (dbc *dbClient) Close() error {
	return dbc.DB().Close()
}

func (dbc *dbClient) DB() *sql.DB {
	return dbc.db.DB
}