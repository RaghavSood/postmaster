package types

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type SNSBase struct {
	Type             string    `json:"Type"`
	MessageID        string    `json:"MessageId"`
	TopicArn         string    `json:"TopicArn"`
	Subject          string    `json:"Subject,omitempty"`
	Timestamp        time.Time `json:"Timestamp"`
	SignatureVersion string    `json:"SignatureVersion"`
	Signature        string    `json:"Signature"`
	SigningCertURL   string    `json:"SigningCertURL"`
	UnsubscribeURL   string    `json:"UnsubscribeURL"`
}

func (b SNSBase) LogFields() log.Fields {
	return log.Fields{
		"Type":             b.Type,
		"MessageId":        b.MessageID,
		"TopicArn":         b.TopicArn,
		"Subject":          b.Subject,
		"Timestamp":        b.Timestamp,
		"SignatureVersion": b.SignatureVersion,
		"Signature":        b.Signature,
		"SigningCertURL":   b.SigningCertURL,
		"UnsubscribeURL":   b.UnsubscribeURL,
	}
}
