package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/strawHat121/expense-tracker-telegram-bot/config"
)

func main() {

	apiKey := config.GetAPIKey()

	if apiKey == "" {
		fmt.Println("API_KEY not set")
		return
	}

	resp, err := http.Get("https://api.telegram.org/bot" + apiKey + "/getMe")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
	fmt.Println(resp.Status)
	fmt.Println(resp.ContentLength)
}
