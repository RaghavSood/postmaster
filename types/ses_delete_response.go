package types

type SESDeleteResonse struct {
	Email         string `json:"email"`
	Response      string `json:"output"`
	ResultMessage string `json:"result_message"`
}
