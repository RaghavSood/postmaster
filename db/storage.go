package db

import (
	"fmt"

	"github.com/RaghavSood/postmaster/types"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

const PAGE_SIZE = 3

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

func (c *Client) GetEvents(from uint64, email string, event string, direction string) ([]*types.SESEvent, error) {
	found := []*types.SESEvent{}

	eArr := []string{email}
	if from == 0 {
		// sql package does not allow the highest bit to be set, therefore this is as high as we can go
		from = 1<<63 - 1
	}

	operator := "<"
	sortOrder := "DESC"
	if direction == "prev" {
		operator = ">"
		sortOrder = "ASC"
	}

	if direction == "last" {
		operator = ">"
		sortOrder = "ASC"
		from = 0
	}

	if direction == "first" {
		from = 1<<63 - 1
	}

	query := "SELECT * FROM ses_events WHERE id" + operator + "$1 AND ($2='{\"\"}' OR recipients@>$2::text[]) AND ($3='' OR event_type=$3) ORDER BY id " + sortOrder + " LIMIT $4"
	err := c.db.Select(&found, query, from, pq.Array(&eArr), event, PAGE_SIZE)

	if direction == "prev" || direction == "last" {
		found = reverseSESEvents(found)
	}

	return found, err
}

func (c *Client) GetMessageEvents(from uint64, messageId string, direction string) ([]*types.SESEvent, error) {
	found := []*types.SESEvent{}

	if from == 0 {
		// sql package does not allow the highest bit to be set, therefore this is as high as we can go
		from = 1<<63 - 1
	}

	operator := "<"
	sortOrder := "DESC"
	if direction == "prev" {
		operator = ">"
		sortOrder = "ASC"
	}

	query := "SELECT * FROM ses_events WHERE id" + operator + "$1 AND message_id=$2 ORDER BY id " + sortOrder + " LIMIT $3"
	err := c.db.Select(&found, query, from, messageId, PAGE_SIZE)

	if direction == "prev" {
		found = reverseSESEvents(found)
	}

	return found, err
}

func reverseSESEvents(in []*types.SESEvent) (out []*types.SESEvent) {
	out = make([]*types.SESEvent, len(in))

	copy(out, in)

	for i := len(out)/2 - 1; i >= 0; i-- {
		opp := len(out) - 1 - i
		out[i], out[opp] = out[opp], out[i]
	}

	return
}
