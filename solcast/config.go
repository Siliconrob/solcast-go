package solcast

import (
	"os"
)

const BaseUrl = "https://api.solcast.com.au"
const Solcast_API_KeyName = "SOLCAST_API_KEY"

type Config struct {
	Url    string
	APIKey string
}

func Read() Config {
	return Config{
		Url:    BaseUrl,
		APIKey: os.Getenv(Solcast_API_KeyName),
	}
}
