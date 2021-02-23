package types

import (
	"time"
)

type SESCheckResonse struct {
	Email string `json:"email_address"`

	LastUpdatedTime time.Time `json:"last_updated_time"`

	Reason string `json:"reason"`
}
