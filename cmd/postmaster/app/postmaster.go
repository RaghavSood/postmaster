package app

import (
	"os"

	"github.com/RaghavSood/postmaster/db"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var config Config

type Postmaster struct {
	db *db.Client
}

func Serve(configPath string) error {
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		return errors.Wrap(err, "error loading config file")
	}

	if err := viper.Unmarshal(&config); err != nil {
		return errors.Wrap(err, "error parsing config")
	}

	if config.LogFile != "" {
		file, err := os.OpenFile(config.LogFile, os.O_APPEND|os.O_WRONLY, 0666)
		if err == nil {
			log.SetFormatter(&log.JSONFormatter{})
			log.SetOutput(file)
		} else {
			log.SetOutput(os.Stdout)
			log.Info("Failed to log to file, logging to stdout")
		}
	}

	return runApp()
}

func runApp() error {
	dbClient, err := db.NewClient(config.Database)
	if err != nil {
		return errors.Wrap(err, "could not connect to database")
	}

	err = dbClient.AutoMigrate()
	if err != nil {
		return errors.Wrap(err, "database migration failed")
	}

	return nil
}
