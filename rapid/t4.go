package rapid

import (
	"encoding/json"
	"fmt"
	"github.com/xueqianLu/twitter-bee/models/apimodels"
	"github.com/xueqianLu/twitter-bee/types"
	"io"
	"net/http"
)

var (
	_ types.RAPIUserGetter = &t4{}
)

type t4 struct {
	key    string
	client *http.Client
}

func (t *t4) Name() string {
	// 300/month
	return "t4"
}

func (t *t4) GetUserInfo(name string) (types.UserProfile, error) {
	return types.UserProfile{}, ErrNotSupport
}

func (t *t4) GetFollowerIDs(param apimodels.FollowerListRequest) (apimodels.FollowerListResponse, error) {
	type response struct {
		Ids               []string    `json:"ids"`
		NextCursor        int64       `json:"next_cursor"`
		NextCursorStr     string      `json:"next_cursor_str"`
		PreviousCursor    int         `json:"previous_cursor"`
		PreviousCursorStr string      `json:"previous_cursor_str"`
		TotalCount        interface{} `json:"total_count"`
	}
	var url string
	if param.Cursor == "" {
		url = fmt.Sprintf("https://twitter293.p.rapidapi.com/user/%s/followers/ids?count=%d",
			param.ID, 5000)
	} else {
		url = fmt.Sprintf("https://twitter293.p.rapidapi.com/user/%s/followers/ids?count=%d&cursor=%s",
			param.ID, 5000, param.Cursor)
	}

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", t.key)
	req.Header.Add("x-rapidapi-host", "twitter293.p.rapidapi.com")

	res, err := t.client.Do(req)
	if err != nil {
		return apimodels.FollowerListResponse{}, err
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	var decode response
	if err = json.Unmarshal(body, &decode); err != nil {
		return apimodels.FollowerListResponse{}, err
	}
	list := make([]apimodels.FollowerObj, len(decode.Ids))
	for i, id := range decode.Ids {
		list[i] = apimodels.FollowerObj{
			ID: id,
		}
	}
	return apimodels.FollowerListResponse{
		List: list,
		Next: decode.NextCursorStr,
	}, nil
}

func (t *t4) Setup(client *http.Client) error {
	t.client = client
	return nil
}
