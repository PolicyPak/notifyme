package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	ErrorReceived = errors.New("not ok message received from Slack environment")
	url           = os.Getenv("url")
	webhookUrl    = os.Getenv("web_hook_url")
)

type SlackRequestBody struct {
	Text string `json:"text"`
}

func main() {
	urls := strings.Split(url, ",")
	CheckListOfURLs(urls)
}

func CheckListOfURLs(urls []string) {
	for _, url := range urls {
		fmt.Printf("Checking for %s\n", url)
		if err := checkStatus(url); err != nil {
			if err := SendSlackNotification(webhookUrl, err.Error()); err != nil {
				log.Fatalf("Error on sending slack notification ! ")
			}
		}

	}
}

func checkStatus(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	fmt.Println(resp.StatusCode)
	return nil
}

func SendSlackNotification(webhookUrl string, msg string) error {

	slackBody, _ := json.Marshal(SlackRequestBody{Text: msg})
	req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(slackBody))
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
