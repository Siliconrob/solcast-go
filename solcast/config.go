package solcast

import (
	"flag"
	"os"
	"log"
)

const BaseUrl = "https://api.solcast.com.au"

type Config struct {
	Url string
	APIKey string
}

func Read() Config {

	var currentKey string
	flag.StringVar(&currentKey, "key", "", "API key for Solcast library")
	flag.Parse()

	if currentKey == "" {
		currentKey = os.Getenv("SOLCAST_API_KEY")
	} else if len(currentKey) != 32 {
		log.Printf("Input Solcast API key is not proper format %s, reading from environment variables", currentKey)
		currentKey = os.Getenv("SOLCAST_API_KEY")
	}

	return Config{
		Url:BaseUrl,
		APIKey: currentKey,
	}
}