package app

type Config struct {
	Database string `mapstructure:"database" json:"database"`
	LogFile  string `mapstructure:"log_file" json:"log_file"`
}
