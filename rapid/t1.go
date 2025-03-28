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
	_ types.RAPIUserGetter = &t1{}
)

type t1 struct {
	key    string
	client *http.Client
}

func (t *t1) GetUserInfo(name string) (types.UserProfile, error) {
	type response struct {
		Profile   string `json:"profile"`
		RestId    string `json:"rest_id"`
		Name      string `json:"name"`
		Friends   int    `json:"friends"`
		SubCount  int    `json:"sub_count"`
		CreatedAt string `json:"created_at"`
		Id        string `json:"id"`
	}
	url := fmt.Sprintf("https://twitter-api45.p.rapidapi.com/screenname.php?screenname=%s", name)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", t.key)
	req.Header.Add("x-rapidapi-host", "twitter-api45.p.rapidapi.com")

	res, err := t.client.Do(req)
	if err != nil {
		return types.UserProfile{}, err
	} else {
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)
		var decode response
		if err = json.Unmarshal(body, &decode); err != nil {
			return types.UserProfile{}, err
		}
		return types.UserProfile{
			Username:  decode.Name,
			Followers: decode.SubCount,
			Following: decode.Friends,
			UserId:    decode.Id,
		}, nil
	}

}

func (t *t1) GetFollowerIDs(req apimodels.FollowerListRequest) (apimodels.FollowerListResponse, error) {
	return apimodels.FollowerListResponse{}, ErrNotSupport
}

func (t *t1) Setup(client *http.Client) error {
	t.client = client
	return nil
}
