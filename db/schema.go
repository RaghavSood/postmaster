package db

var migrationTableCreate = `CREATE TABLE IF NOT EXISTS migrations (
  latest_version integer
); 

INSERT INTO migrations (latest_version) VALUES (0);`

var sesEventsTableCreate = `CREATE TABLE ses_events (
    event_type text NOT NULL,
    message_id text,
    recipients text [],
    mail jsonb,
    event_data jsonb,
    received_at timestamp without time zone default (now() at time zone 'utc'),
    sns_id text NOT NULL
);

UPDATE migrations SET latest_version=1 WHERE latest_version=0;`

var sesEventsIndices = `
CREATE INDEX idx_ses_events_types on ses_events (event_type);
CREATE INDEX idx_ses_events_message_id on ses_events (message_id);
CREATE UNIQUE INDEX idx_ses_events_sns_id on ses_events (sns_id);
CREATE INDEX idx_ses_events_recipients on ses_events USING GIN(recipients);
CREATE INDEX idx_ses_events_mail on ses_events USING GIN(mail);
CREATE INDEX idx_ses_events_event_data on ses_events USING GIN(event_data);

UPDATE migrations SET latest_version=2 WHERE latest_version=1;`

var sesEventsPrimaryKey = `
ALTER TABLE ses_events ADD COLUMN id BIGSERIAL PRIMARY KEY;

UPDATE migrations SET latest_version=3 WHERE latest_version=2;`

func (c *Client) migrationsExists() (bool, error) {
	var exists bool
	err := c.db.QueryRow("select exists ( select from information_schema.tables where table_name='migrations')").Scan(&exists)
	return exists, err
}

func (c *Client) latestMigration() (int, error) {
	var latest int
	err := c.db.QueryRow("SELECT latest_version FROM migrations LIMIT 1").Scan(&latest)
	return latest, err
}
