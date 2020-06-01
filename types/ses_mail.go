package types

import "time"

type SESMail struct {
	Timestamp        time.Time     `json:"timestamp"`
	Source           string        `json:"source"`
	SendingAccountID string        `json:"sendingAccountId"`
	MessageID        string        `json:"messageId"`
	Destination      []string      `json:"destination"`
	HeadersTruncated bool          `json:"headersTruncated"`
	Headers          []MailHeaders `json:"headers"`
	CommonHeaders    CommonHeaders `json:"commonHeaders"`
	Tags             Tags          `json:"tags"`
}

type MailHeaders struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CommonHeaders struct {
	From      []string `json:"from"`
	Date      string   `json:"date"`
	To        []string `json:"to"`
	MessageID string   `json:"messageId"`
	Subject   string   `json:"subject"`
}

type Tags struct {
	SesOperation        []string `json:"ses:operation"`
	SesConfigurationSet []string `json:"ses:configuration-set"`
	SesSourceIP         []string `json:"ses:source-ip"`
	SesFromDomain       []string `json:"ses:from-domain"`
	SesCallerIdentity   []string `json:"ses:caller-identity"`
}
