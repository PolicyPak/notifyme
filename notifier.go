package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

var (
	ErrorReceived = errors.New("not ok message received from Slack environment")
	MESSAGE           = os.Getenv("MESSAGE")
	WEB_HOOK_URL      = os.Getenv("web_hook_url")
)

type SlackRequestBody struct {
	Text string `json:"text"`
}

func main() {
	if err := SendSlackNotification(WEB_HOOK_URL,MESSAGE); err !=nil {
		panic(fmt.Errorf("Error on sending notification %v", err))
	}
}


func SendSlackNotification(webHookUrl string, msg string) error {

	slackBody, _ := json.Marshal(SlackRequestBody{Text: msg})
	req, err := http.NewRequest(http.MethodPost, webHookUrl, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return ErrorReceived
	}
	return nil
}
