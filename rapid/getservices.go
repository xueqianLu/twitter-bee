package rapid

import (
	"errors"
	"github.com/xueqianLu/twitter-bee/types"
	"net/http"
)

var (
	ErrNotSupport = errors.New("not support")
)

func GetAllServices(keylist []string, client *http.Client) ([]types.RAPIUserGetter, []types.RAPIFollowerGetter) {
	user := make([]types.RAPIUserGetter, 0)
	follower := make([]types.RAPIFollowerGetter, 0)

	for _, key := range keylist {
		user = append(user, &t1{key: key})
		follower = append(follower, &t2{key: key})
		follower = append(follower, &t3{key: key})
		follower = append(follower, &t4{key: key})
	}

	return user, follower
}
