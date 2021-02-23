package app

import (
	"net/http"
	"os"

	"github.com/RaghavSood/postmaster/db"
	"github.com/RaghavSood/postmaster/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sesv2"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var config Config

type Postmaster struct {
	sesClient *sesv2.SESV2
	db        *db.Client
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

	creds := credentials.NewStaticCredentials(config.SESKey, config.SESSecret, "")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-central-1"),
		Credentials: creds,
	})

	svc := sesv2.New(sess)

	postmaster := &Postmaster{
		db:        dbClient,
		sesClient: svc,
	}

	err = postmaster.run()
	if err != nil {
		return errors.Wrap(err, "postmaster stopped running")
	}

	return nil
}

func InjectSES(sesClient *sesv2.SESV2) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("ses_client", sesClient)
		c.Next()
	}
}

func InjectDatabase(dbc *db.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("database", dbc)
		c.Next()
	}
}

func (p *Postmaster) run() error {
	err := p.db.AutoMigrate()
	if err != nil {
		return errors.Wrap(err, "database migration failed")
	}

	router := gin.Default()

	router.Use(InjectDatabase(p.db))
	router.Use(InjectSES(p.sesClient))

	router.POST("/sns_hook", processHook)
	router.GET("/api/events", getEvents)
	router.GET("/api/message", getMessageEvents)
	router.GET("/api/suppression/check", getSuppressionListCheck)
	router.GET("/api/suppression/delete", getSuppressionListDelete)
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "dashboard/")
	})
	router.StaticFS("/dashboard", AssetFile())

	router.Run()

	return nil
}

func getSuppressionListCheck(c *gin.Context) {
	var query types.SESQuery

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	input := sesv2.GetSuppressedDestinationInput{
		EmailAddress: &query.Email,
	}

	sesClient, ok := c.MustGet("ses_client").(*sesv2.SESV2)
	if !ok {
		log.Warn("Could not get SES Client")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get SES Client"})
		return
	}

	output, err := sesClient.GetSuppressedDestination(&input)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Warn("Could not check email status")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not check email status"})
		return
	}
	checkResponse := types.SESCheckResonse{
		Email:           *output.SuppressedDestination.EmailAddress,
		LastUpdatedTime: *output.SuppressedDestination.LastUpdateTime,
		Reason:          *output.SuppressedDestination.Reason,
	}
	c.JSON(http.StatusOK, gin.H{"results": checkResponse})
}

func getSuppressionListDelete(c *gin.Context) {
	var query types.SESQuery

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	input := sesv2.DeleteSuppressedDestinationInput{
		EmailAddress: &query.Email,
	}

	sesClient, ok := c.MustGet("ses_client").(*sesv2.SESV2)
	if !ok {
		log.Warn("Could not get SES Client")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get SES Client"})
		return
	}

	output, err := sesClient.DeleteSuppressedDestination(&input)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Warn("Could not delete email from suppression list")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete email from suppression list"})
		return
	}
	checkResponse := types.SESDeleteResonse{
		Response: output.String(),
	}
	c.JSON(http.StatusOK, gin.H{"results": checkResponse})
}

func getMessageEvents(c *gin.Context) {
	var query types.EventPageQuery

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if query.Direction == "" {
		query.Direction = "next"
	}

	if query.Direction != "next" && query.Direction != "prev" && query.Direction != "first" && query.Direction != "last" {
		c.JSON(http.StatusBadRequest, "invalid direction")
	}

	dbConn, ok := c.MustGet("database").(*db.Client)
	if !ok {
		log.Warn("Could not get database")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get database connection"})
		return
	}

	eventList, err := dbConn.GetMessageEvents(query.From, query.MessageId, query.Direction)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Warn("Could not get events")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get events"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": eventList})

}

func getEvents(c *gin.Context) {
	var query types.EventPageQuery

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if query.Direction == "" {
		query.Direction = "next"
	}

	if query.Direction != "next" && query.Direction != "prev" && query.Direction != "first" && query.Direction != "last" {
		c.JSON(http.StatusBadRequest, "invalid direction")
	}

	dbConn, ok := c.MustGet("database").(*db.Client)
	if !ok {
		log.Warn("Could not get database")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get database connection"})
		return
	}

	eventList, err := dbConn.GetEvents(query.From, query.EmailFitler, query.EventFilter, query.Direction)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Warn("Could not get events")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get events"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": eventList})

}

func processHook(c *gin.Context) {
	var headers types.SNSHeaders

	if err := c.ShouldBindHeader(&headers); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	switch headers.MessageType {
	case "Notification":
		var notif types.SESNotification

		if err := c.ShouldBindJSON(&notif); err == nil {
			notif.Event.SNSID = headers.MessageID

			dbConn, ok := c.MustGet("database").(*db.Client)
			if !ok {
				log.Warn("Could not get database")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get database connection"})
				return
			}

			err = dbConn.InsertEvent(notif.Event)
			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
				}).Warn("Could not insert event")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get save event into database"})
				return

			}

			c.JSON(http.StatusOK, gin.H{"success": true})
		} else {
			log.WithFields(log.Fields{
				"error": err,
			}).Warn("Could not parse SNS notification")
			c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse webhook"})
			return
		}
	case "SubscriptionConfirmation":
		var subscription types.SNSSubscription

		if err := c.ShouldBindJSON(&subscription); err == nil {
			log.WithFields(subscription.LogFields()).Info("Received SNS webhook")

			err, status := confirmSubscription(subscription.SubscribeURL)
			if err != nil {
				log.WithFields(log.Fields{
					"error":  err,
					"status": status,
				}).Warn("Could not confirm subscription")
				c.JSON(http.StatusBadRequest, gin.H{"error": "could not confirm subscription"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"success": true})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse webhook"})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, "")
	}

}

func confirmSubscription(subcribeURL string) (error, int) {
	response, err := http.Get(subcribeURL)
	if err != nil {
		return errors.Wrap(err, "could not confirm subscription"), response.StatusCode
	}

	return nil, response.StatusCode
}
