package db

import (
	"fmt"

	"github.com/RaghavSood/postmaster/types"
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

func (c *Client) InsertEvent(event types.SESEvent) error {
	rows, err := c.db.NamedQuery("INSERT INTO ses_events (event_type, message_id, recipients, mail, event_data, sns_id) VALUES (:event_type, :message_id, :recipients, :mail, :event_data, :sns_id) ON CONFLICT (sns_id) DO NOTHING", event)
	if rows != nil {
		rows.Close()
	}

	return err
}

func (c *Client) ExistsSNSMessageID(messageID string) (bool, error) {
	return c.rowExists("SELECT sns_id FROM ses_events WHERE sns_id=$1", messageID)
}

func (c *Client) rowExists(query string, args ...interface{}) (bool, error) {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := c.db.QueryRow(query, args...).Scan(&exists)
	return exists, err
}
