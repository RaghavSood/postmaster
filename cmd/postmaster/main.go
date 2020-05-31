package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"os"

	"github.com/RaghavSood/postmaster/cmd/postmaster/app"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

var f *flag.FlagSet

func init() {
	f = flag.NewFlagSet("config", flag.ContinueOnError)
	f.Usage = func() {
		fmt.Println(f.FlagUsages())
		os.Exit(0)
	}

	f.String("config", "config.yml",
		"path to the config file")

	if err := f.Parse(os.Args[1:]); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Error loading flags")
	}
}

func main() {
	configFile, _ := f.GetString("config")
	if err := app.Serve(configFile); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Failed to serve postmaster")
		os.Exit(1)
	}
}
