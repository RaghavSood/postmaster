package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Client struct {
	db *sqlx.DB
}

func NewClient(dsn string) (*Client, error) {
	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		return nil, errors.Wrap(err, "could not connect to database")
	}

	return &Client{
		db: db,
	}, nil
}
