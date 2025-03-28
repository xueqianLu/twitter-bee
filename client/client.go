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

func (c *BeeClient) GetUserProfile(userName string) (*apimodels.UserInfoResponse, error) {
	url := fmt.Sprintf("%s/tbapi/v2/follower/profile", c.BaseURL)
	reqBody, err := json.Marshal(apimodels.UserInfoRequest{UserName: userName})
	if err != nil {
		fmt.Println("GetUserProfile json.Marshal error:", err)
		return nil, err
	}

	resp, err := c.HTTPClient.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("GetUserProfile HTTPClient.Post error:", err)
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

	var data apimodels.UserInfoResponse
	if err := json.Unmarshal(result.Data, &data); err != nil {
		fmt.Println("GetFollowerCount json.Unmarshal result error:", err)
		return nil, err
	}
	return &data, nil
}

func (c *BeeClient) GetFollowerList(name string, id string, cursor string) (*apimodels.FollowerListResponse, error) {
	url := fmt.Sprintf("%s/tbapi/v2/follower/list", c.BaseURL)
	reqBody, err := json.Marshal(apimodels.FollowerListRequest{Name: name, ID: id, Cursor: cursor})
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
