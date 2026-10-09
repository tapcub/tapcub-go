package tapcub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	APIKey   string
	Endpoint string
	HTTP     *http.Client
}

func New(apiKey, endpoint string) *Client {
	return &Client{APIKey: apiKey, Endpoint: strings.TrimRight(endpoint, "/"), HTTP: &http.Client{Timeout: 8 * time.Second}}
}

func (c *Client) Track(event string, payload map[string]any) error {
	if payload == nil {
		payload = map[string]any{}
	}
	payload["type"] = "track"
	payload["event"] = event
	if payload["context"] == nil {
		payload["context"] = map[string]any{"platform": "server", "library": map[string]any{"name": "tapcub-go", "version": "0.1.0"}}
	}
	return c.send([]map[string]any{payload})
}

func (c *Client) send(events []map[string]any) error {
	body, _ := json.Marshal(map[string]any{"events": events})
	req, err := http.NewRequest(http.MethodPost, c.Endpoint+"/api/v1/track", bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	res, err := c.HTTP.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode >= 300 {
		return fmt.Errorf("tapcub track %s", res.Status)
	}
	return nil
}
