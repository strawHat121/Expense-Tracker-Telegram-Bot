package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/strawHat121/expense-tracker-telegram-bot/config"
	"github.com/strawHat121/expense-tracker-telegram-bot/db"
)

type UpdateResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	MessageId int    `json:"message_id"`
	Text      string `json:"text"`
}

func main() {

	apiKey := config.GetAPIKey()

	if apiKey == "" {
		fmt.Println("API_KEY not set")
		return
	}

	database := db.InitDB("Expenses.db")

	defer database.Close()

	resp, err := http.Get("https://api.telegram.org/bot" + apiKey + "/getUpdates")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var telegramResponse UpdateResponse

	body, err := io.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &telegramResponse); err != nil {
		log.Fatal("Error in decoding JSON", err)
	}

	for _, update := range telegramResponse.Result {

		var amount int
		var comment string

		splittedString := strings.Split(update.Message.Text, " ")

		category := splittedString[0]
		if len(splittedString) > 1 {
			amount, _ = strconv.Atoi(splittedString[1])
			comment = splittedString[2]
		}

		fmt.Println(category, amount, comment)

	}

	if err != nil {
		log.Fatal(err)
	}
}
