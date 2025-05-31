package http_client

import (
	"net/http"
	"net/url"
	"time"
)

type HttpClient struct {
	client *http.Client
}

func NewHttpClient() *HttpClient {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	return &HttpClient{
		client: &client,
	}
}

func (c *HttpClient) Get(newURL string, params map[string]string) (*http.Response, error) {
	baseUrl, err := url.Parse(newURL)
	if err != nil {
		return nil, err
	}

	query := baseUrl.Query()
	for key, value := range params {
		query.Add(key, value)
	}

	baseUrl.RawQuery = query.Encode()

	return c.client.Get(baseUrl.String())
}
