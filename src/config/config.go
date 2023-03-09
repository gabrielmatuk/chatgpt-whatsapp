package config

import (
	"os"
)

var (
	//Need to input your API KEY from ChatGPT in .env file.
	ApiKey = ""
)

func LoadingEnv() {
	ApiKey = os.Getenv("API_TOKEN_GPT")
}
