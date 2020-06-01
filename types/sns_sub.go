package types

import (
	log "github.com/sirupsen/logrus"
)

type SNSSubscription struct {
	SNSBase
	Token        string `json:"Token"`
	Message      string `json:"Message"`
	SubscribeURL string `json:"SubscribeURL"`
}

func (s SNSSubscription) LogFields() log.Fields {
	return log.Fields{
		"Type":             s.Type,
		"MessageId":        s.MessageID,
		"Token":            s.Token,
		"TopicArn":         s.TopicArn,
		"Message":          s.Message,
		"SubscribeURL":     s.SubscribeURL,
		"Timestamp":        s.Timestamp,
		"SignatureVersion": s.SignatureVersion,
		"Signature":        s.Signature,
		"SigningCertURL":   s.SigningCertURL,
	}
}
