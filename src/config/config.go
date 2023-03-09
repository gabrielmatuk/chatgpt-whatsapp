package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	//Need to input your API KEY from ChatGPT in .env file.
	ApiKey = ""
)

func LoadingEnv() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	ApiKey = os.Getenv("API_TOKEN_GPT")
}
