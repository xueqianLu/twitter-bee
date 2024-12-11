package client

import (
	"net"
	"net/http"
	"time"
)

type BeeClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func newHttpClient() *http.Client {
	hclient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     90 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
			DisableCompression:  true,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
		},
		Timeout: 30 * time.Second,
	}
	return hclient
}

func NewClient(baseURL string) *BeeClient {
	return &BeeClient{
		BaseURL:    baseURL,
		HTTPClient: newHttpClient(),
	}
}

//
