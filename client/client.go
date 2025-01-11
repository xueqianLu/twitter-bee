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

type Response struct {
	Code int             `json:"code"`
	Data json.RawMessage `json:"data"`
	Msg  string          `json:"message"`
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
		fmt.Println("GetFollowerCount json.Marshal error:", err)
		return nil, err
	}

	resp, err := c.HTTPClient.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("GetFollowerCount HTTPClient.Post error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("GetFollowerCount unexpected status code:", resp.StatusCode)
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("GetFollowerCount decode http result error:", err)
		return nil, err
	}

	var data apimodels.FollowerCountResponse
	if err := json.Unmarshal(result.Data, &data); err != nil {
		fmt.Println("GetFollowerCount json.Unmarshal result error:", err)
		return nil, err
	}
	return &data, nil
}

func (c *BeeClient) GetFollowerList(user string, cursor string) (*apimodels.FollowerListResponse, error) {
	url := fmt.Sprintf("%s/tbapi/v1/follower/list", c.BaseURL)
	reqBody, err := json.Marshal(apimodels.FollowerListRequest{User: user, Cursor: cursor})
	if err != nil {
		fmt.Println("GetFollowerList json.Marshal error:", err)
		return nil, err
	}

	resp, err := c.HTTPClient.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("GetFollowerList HTTPClient.Post error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("GetFollowerList unexpected status code:", resp.StatusCode)
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("GetFollowerList decode http result error:", err)
		return nil, err
	}
	var data apimodels.FollowerListResponse
	if err := json.Unmarshal(result.Data, &data); err != nil {
		fmt.Println("GetFollowerList json.Unmarshal result error:", err)
		return nil, err
	}
	return &data, nil
}
