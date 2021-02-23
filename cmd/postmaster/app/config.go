package app

type Config struct {
	Database  string `mapstructure:"database" json:"database"`
	LogFile   string `mapstructure:"log_file" json:"log_file"`
	SESKey    string `mapstructure:"ses_key" json:"ses_key"`
	SESSecret string `mapstructure:"ses_secret" json:"ses_secret"`
}
