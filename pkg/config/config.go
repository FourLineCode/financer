package config

type Config struct {
	Port        string
	DatabaseURL string
}

func GetConfig() *Config {
	return &Config{
		Port:        ":5000",
		DatabaseURL: "./db/db.sqlite",
	}
}
