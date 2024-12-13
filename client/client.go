package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xueqianLu/twitter-bee/models/apimodels"
	"net/http"
	"time"
)

type BeeClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewBeeClient(baseURL string) *BeeClient {
	return &BeeClient{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *BeeClient) GetFollowerCount(userID string) (*apimodels.FollowerCountResponse, error) {
	url := fmt.Sprintf("%s/tbapi/v1/follower/count", c.BaseURL)
	reqBody, err := json.Marshal(apimodels.FollowerCountRequest{UserId: userID})
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result apimodels.FollowerCountResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *BeeClient) GetFollowerList(user string, cursor string) (*apimodels.FollowerListResponse, error) {
	url := fmt.Sprintf("%s/tbapi/v1/follower/list", c.BaseURL)
	reqBody, err := json.Marshal(apimodels.FollowerListRequest{User: user, Cursor: cursor})
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result apimodels.FollowerListResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
