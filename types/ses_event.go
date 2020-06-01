package types

import (
	"encoding/json"
	"time"

	"github.com/RaghavSood/postmaster/common"
	sqltypes "github.com/jmoiron/sqlx/types"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type SESNotification struct {
	SNSBase
	Event SESEvent `json:"Message"`
}

type SESEvent struct {
	EventType  string            `json:"eventType" db:"event_type"`
	MessageID  string            `json:"messageId" db:"message_id"`
	Recipients pq.StringArray    `json:"recipients" db:"recipients"`
	Mail       sqltypes.JSONText `json:"mail" db:"mail"`
	EventData  sqltypes.JSONText `json:"event_data" db:"event_data"`
	ReceivedAt time.Time         `json:"received_at" db:"received_at"`
	SNSID      string            `json:"sns_id" db:"sns_id"`
}

type fauxSESEvent SESEvent

type events struct {
	Bounce    json.RawMessage `json:"bounce,omitempty"`
	Complaint json.RawMessage `json:"complaint,omitempty"`
	Delivery  json.RawMessage `json:"delivery,omitempty"`
	Send      json.RawMessage `json:"send,omitempty"`
	Reject    json.RawMessage `json:"reject,omitempty"`
	Open      json.RawMessage `json:"open,omitempty"`
	Click     json.RawMessage `json:"click,omitempty"`
	Failure   json.RawMessage `json:"failure,omitempty"`
}

func (e *SESEvent) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	var f fauxSESEvent
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}

	*e = SESEvent(f)

	var mailBody SESMail
	if err := json.Unmarshal(f.Mail, &mailBody); err != nil {
		return err
	}

	var eventData events
	if err := json.Unmarshal([]byte(s), &eventData); err != nil {
		return err
	}

	e.MessageID = mailBody.MessageID
	e.ReceivedAt = time.Now().UTC()
	e.Recipients = mailBody.Destination
	e.EventData = common.ConcatMany([][]byte{eventData.Bounce, eventData.Complaint, eventData.Delivery, eventData.Send, eventData.Reject, eventData.Open, eventData.Click, eventData.Failure})

	return nil
}

func (e SESEvent) LogFields() log.Fields {
	return log.Fields{
		"event_type":  e.EventType,
		"message_id":  e.MessageID,
		"recipients":  e.Recipients,
		"mail":        string(e.Mail),
		"event_data":  string(e.EventData),
		"received_at": e.ReceivedAt,
		"sns_id":      e.SNSID,
	}
}
