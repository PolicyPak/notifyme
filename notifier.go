package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log"
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
	log.Printf("Message: \" %s \" has been send\n",MESSAGE)
}


func SendSlackNotification(webHookUrl string, msg string) error {
	format := "2006-01-02 15:04:05"
	slackBody, _ := json.Marshal(SlackRequestBody{Text: time.Now().Format(format)+": "+msg})
	req, err := http.NewRequest(http.MethodPost, webHookUrl, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Timeout: 10 * time.Second,Transport: tr}
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
