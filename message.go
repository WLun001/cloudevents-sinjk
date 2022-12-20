package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func sendMessage(text string) {

	body, _ := json.Marshal(map[string]string{
		"chat_id": os.Getenv("TELEGRAM_CHAT_ID"),
		"text":    text,
	})

	requestURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage",
		os.Getenv("TELEGRAM_BOT_TOKEN"))
	req, err := http.Post(requestURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(req.Body)
}
