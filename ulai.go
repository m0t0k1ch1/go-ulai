package ulai

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const (
	DefaultUri = "https://chatbot-api.userlocal.jp/api/chat"

	ResponseStatusSuccess = "success"
)

var (
	ErrNoKey = errors.New("no key")
)

type ChatResponse struct {
	Status string `json:"status"`
	Result string `json:"result"`
}

func (res *ChatResponse) isSuccess() bool {
	return res.Status == ResponseStatusSuccess
}

type Config struct {
	Uri string
	Key string
}

type Client struct {
	*http.Client
	config *Config
}

func NewClient() *Client {
	return &Client{
		Client: http.DefaultClient,
		config: &Config{
			Uri: DefaultUri,
		},
	}
}

func (client *Client) SetUri(uri string) {
	client.config.Uri = uri
}

func (client *Client) SetKey(key string) {
	client.config.Key = key
}

func (client *Client) Chat(ctx context.Context, message string) (string, error) {
	v := url.Values{}
	v.Add("key", client.config.Key)
	v.Add("message", message)

	if len(client.config.Key) == 0 {
		return "", ErrNoKey
	}

	u, err := url.Parse(client.config.Uri)
	if err != nil {
		return "", err
	}
	u.RawQuery = v.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return "", err
	}
	req.WithContext(ctx)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var res ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}
	if !res.isSuccess() {
		return "", fmt.Errorf(res.Status)
	}

	return res.Result, nil
}
