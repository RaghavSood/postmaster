package types

type SESQuery struct {
	Email  string `form:"email"`
	Reason string `form:"reason"`
}
