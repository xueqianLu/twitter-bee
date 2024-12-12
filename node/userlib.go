package node

import (
	"encoding/json"
	"github.com/xueqianLu/twitter-bee/types"
	"io/ioutil"
)

func getUserLib(path string) map[string]types.TwitterAccount {
	// read file from path and unmarshal json to map.
	res := make(map[string]types.TwitterAccount)
	userlist := make([]types.TwitterAccount, 0)
	data, _ := ioutil.ReadFile(path)
	_ = json.Unmarshal(data, &userlist)
	for _, user := range userlist {
		res[user.Username] = user
	}
	return res
}
