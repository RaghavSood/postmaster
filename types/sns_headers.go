package types

type SNSHeaders struct {
	MessageType     string `header:"x-amz-sns-message-type"`
	MessageID       string `header:"x-amz-sns-message-id"`
	TopicARN        string `header:"x-amz-sns-topic-arn"`
	SubscriptionARN string `header:"x-amz-sns-subscription-arn"`
}
