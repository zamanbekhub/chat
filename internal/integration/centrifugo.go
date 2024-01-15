package integration

import (
	"bytes"
	"chat/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type CentrifugoServer interface {
	Push(ctx context.Context, channel string, value string) error
}

type CentrifugoServerClient struct {
	client  *http.Client
	XApiKey string
	baseURL *url.URL
}

func NewCentrifugoServerClient(rawBaseURL string, XApiKey string) (*CentrifugoServerClient, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("NewTerminalClient - url.Parse(%s): %w", rawBaseURL, err)
	}

	return &CentrifugoServerClient{
		client: &http.Client{
			Timeout: time.Minute * 5,
		},
		XApiKey: XApiKey,
		baseURL: baseURL,
	}, nil
}

func (c *CentrifugoServerClient) Push(ctx context.Context, channel string, message string) error {
	requestURL := c.baseURL.JoinPath("/api/publish")
	req, err := c.newRequest(ctx, http.MethodPost, requestURL, model.CentrifugoPublish{
		Channel: channel,
		Data: model.CentrifugoPublishData{
			Value: message,
		},
	})
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	rawResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return err
	}

	var response model.CentrifugoResponse
	if err := json.Unmarshal(rawResponse, &response); err != nil {
		return err
	}

	return nil
}

func (c *CentrifugoServerClient) newRequest(
	ctx context.Context,
	method string,
	requestUrl *url.URL,
	body interface{},
) (*http.Request, error) {
	var bodyReader io.Reader
	if body != nil {
		rawBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		bodyReader = bytes.NewBuffer(rawBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, requestUrl.String(), bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-API-Key", c.XApiKey)

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	return req, nil
}
