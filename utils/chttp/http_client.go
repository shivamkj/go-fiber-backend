package chttp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/qnify/api-server/utils/consts"
)

// Client is a simple HTTP client wrapper
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	RetryCount int
}

func NewClient(baseURL string, reuse bool) *Client {
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	if reuse {
		client.Transport = &http.Transport{
			// configure connection reuse
			MaxIdleConns:        10,
			MaxIdleConnsPerHost: 2,
		}
	}

	return &Client{
		BaseURL:    baseURL,
		HTTPClient: client,
		RetryCount: 3,
	}
}

func (c *Client) Get(endpoint string, headers map[string]string) (*http.Response, []byte, error) {
	resp, err := c.doRequest("GET", endpoint, headers, nil)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return resp, respBody, nil
}

func (c *Client) Post(endpoint string, headers map[string]string, body io.Reader) (*http.Response, []byte, error) {
	resp, err := c.doRequest("POST", endpoint, headers, body)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return resp, respBody, nil
}

func (c *Client) PostJson(endpoint string, body any) (*http.Response, []byte, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.doRequest("POST", endpoint, map[string]string{
		consts.ContentType: consts.JsonType,
	}, bytes.NewBuffer(jsonData))

	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return resp, respBody, nil
}

func (c *Client) doRequest(method, endpoint string, headers map[string]string, body io.Reader) (*http.Response, error) {

	var resp *http.Response

	for retry := 0; retry <= c.RetryCount; retry++ {
		req, err := http.NewRequest(method, c.BaseURL+endpoint, body)
		if err != nil {
			return nil, err
		}

		if len(headers) != 0 {
			for key, value := range headers {
				req.Header.Set(key, value)
			}
		}

		resp, err = c.HTTPClient.Do(req)

		if !(err != nil || resp.StatusCode >= 500 || resp.StatusCode == 0) {
			break
		}

		if resp != nil {
			// Drain the response body, to reuse connection
			io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
		}

		if retry < c.RetryCount {
			// Apply backoff before the next retry
			time.Sleep(time.Duration(1<<uint(retry)) * time.Second)
		}
	}

	return resp, nil
}
