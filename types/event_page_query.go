package types

type EventPageQuery struct {
	From        uint64 `form:"from"`
	EventFilter string `form:"event_filter"`
	EmailFitler string `form:"email_filter"`
	MessageId   string `form:"message_id"`
	Direction   string `form:"direction"`
}
