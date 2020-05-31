package db

var migrationTableCreate = `CREATE TABLE IF NOT EXISTS migrations (
  latest_version integer
); 

INSERT INTO migrations (latest_version) VALUES (0);`

var sesEventsTableCreate = `CREATE TABLE ses_events (
    event_type text NOT NULL,
    message_id text NULL,
    recipients text [],
    mail jsonb,
    event_data jsonb
);

UPDATE migrations SET latest_version=1 WHERE latest_version=0;`

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
