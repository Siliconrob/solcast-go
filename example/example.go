package main

import (
	solcast "../solcast"
	"os"
	"log"
)

func main() {
	config := solcast.Read()
	if config.APIKey == "" {
		log.Printf("No valid Solcast API key available to use set environment variable or pass argument")
		os.Exit(-1)
	}
	log.Printf("API Key: %s", config.APIKey)
}