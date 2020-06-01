package db

import "github.com/pkg/errors"

func (c *Client) AutoMigrate() error {
	var migrationList = []string{
		migrationTableCreate,
		sesEventsTableCreate,
		sesEventsIndices,
		sesEventsPrimaryKey,
	}

	mExists, err := c.migrationsExists()
	if err != nil {
		return errors.Wrap(err, "failed to check if migrations table exists")
	}

	runFrom := 0

	if mExists {
		runFrom, err = c.latestMigration()
		if err != nil {
			return errors.Wrap(err, "failed to check latest migration version")
		}

		runFrom = runFrom + 1 // Run from the next migration available
	}

	for runFrom < len(migrationList) {
		tx, err := c.db.Begin()
		if err != nil {
			return errors.Wrap(err, "failed to get database transaction")
		}

		_, err = tx.Exec(migrationList[runFrom])
		if err != nil {
			return errors.Wrap(err, "failed to create migrations table")
		}

		err = tx.Commit()
		if err != nil {
			return errors.Wrap(err, "failed to commit migration")
		}

		runFrom = runFrom + 1
	}
	return nil
}
