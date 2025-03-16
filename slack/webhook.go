package slack

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// PostToWebhook posts the given Message to the Slack webhook with the given URL.
func PostToWebhook(url string, message Message) error {
	return PostToWebhookWithClient(http.DefaultClient, url, message)
}

// PostToWebhookWithClient uses the given *http.Client to post the given Message to the Slack webhook with the given URL.
func PostToWebhookWithClient(client *http.Client, url string, message Message) error {
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return err
	}

	resp, err := client.Post(url, "application/json", strings.NewReader(string(messageJSON)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("slack: unexpected status code %d, with response '%s'", resp.StatusCode, data)
	}

	if string(data) != "ok" {
		return fmt.Errorf("slack: unexpected response '%s'", data)
	}

	return nil
}
