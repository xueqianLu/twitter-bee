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
	_ types.RAPIUserGetter = &t2{}
)

type t2 struct {
	key    string
	client *http.Client
}

func (t *t2) GetUserInfo(name string) (types.UserProfile, error) {
	return types.UserProfile{}, ErrNotSupport
}

func (t *t2) GetFollowerIDs(param apimodels.FollowerListRequest) (apimodels.FollowerListResponse, error) {
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
		url = fmt.Sprintf("https://twitter135.p.rapidapi.com/v1.1/FollowersIds/?id=%s&count=%d",
			param.ID, 5000)
	} else {
		url = fmt.Sprintf("https://twitter135.p.rapidapi.com/v1.1/FollowersIds/?id=%s&count=%d&cursor=%s",
			param.ID, 5000, param.Cursor)
	}

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", t.key)
	req.Header.Add("x-rapidapi-host", "twitter135.p.rapidapi.com")

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

func (t *t2) Setup(client *http.Client) error {
	t.client = client
	return nil
}
