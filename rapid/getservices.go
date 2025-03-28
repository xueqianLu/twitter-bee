package rapid

import (
	"errors"
	"github.com/xueqianLu/twitter-bee/types"
	"math/rand"
	"net/http"
)

var (
	ErrNotSupport = errors.New("not support")
)

type ServiceBalancer struct {
	// userGetter is a list of user getter
	userGetter []types.RAPIUserGetter
	// followerGetter is a list of follower getter
	followerGetter []types.RAPIFollowerGetter
}

func (b ServiceBalancer) GetRandomUserGetter() []types.RAPIUserGetter {
	idx := rand.Intn(100) % len(b.userGetter)
	slice := make([]types.RAPIUserGetter, 0)
	slice = append(slice, b.userGetter[idx:]...)
	slice = append(slice, b.userGetter[:idx]...)

	return slice
}

func (b ServiceBalancer) GetRandomFollowerGetter() []types.RAPIFollowerGetter {
	idx := rand.Intn(100) % len(b.followerGetter)
	slice := make([]types.RAPIFollowerGetter, 0)
	slice = append(slice, b.followerGetter[idx:]...)
	slice = append(slice, b.followerGetter[:idx]...)

	return slice
}

func GetAllServices(keylist []string, client *http.Client) ServiceBalancer {
	user := make([]types.RAPIUserGetter, 0)
	follower := make([]types.RAPIFollowerGetter, 0)

	for _, key := range keylist {
		user = append(user, &t1{key: key, client: client})
		follower = append(follower, &t2{key: key, client: client})
		follower = append(follower, &t3{key: key, client: client})
		follower = append(follower, &t4{key: key, client: client})
	}

	return ServiceBalancer{
		userGetter:     user,
		followerGetter: follower,
	}
}
