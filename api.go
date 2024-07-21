package yoopenai

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const apiURL = "https://api.openai.com/v1"

type Client struct {
	BaseURL string
	authToken string
	HTTPClient *http.Client
}

func newClient(authToken string) *Client {
	return &Client{
		BaseURL: apiURL,
		authToken: authToken,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func(c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.authToken))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()


	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}