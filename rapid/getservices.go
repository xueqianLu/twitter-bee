package rapid

import (
	"errors"
	"github.com/xueqianLu/twitter-bee/config"
	"github.com/xueqianLu/twitter-bee/types"
	"math/rand"
	"net/http"
	"strings"
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

func (b ServiceBalancer) GetFollowerGetterByStart(start types.RAPIFollowerGetter) []types.RAPIFollowerGetter {
	var idx = 0
	for i, v := range b.followerGetter {
		if v == start {
			idx = i
			break
		}
	}
	slice := make([]types.RAPIFollowerGetter, 0)
	slice = append(slice, b.followerGetter[idx:]...)
	slice = append(slice, b.followerGetter[:idx]...)
	return slice
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

func GetAllServices(keylist []string, client *http.Client, conf *config.Config) ServiceBalancer {
	user := make([]types.RAPIUserGetter, 0)
	follower := make([]types.RAPIFollowerGetter, 0)
	enableList := strings.Split(conf.EnableRAPI, ",")
	var enableMap = make(map[string]bool)
	for _, v := range enableList {
		enableMap[v] = true
	}

	for _, key := range keylist {
		{
			m := &t1{key: key, client: client}
			if enableMap[m.Name()] {
				user = append(user, m)
			}
		}
		{
			m := &t2{key: key, client: client}
			if enableMap[m.Name()] {
				follower = append(follower, m)
			}
		}
		{
			m := &t3{key: key, client: client}
			if enableMap[m.Name()] {
				follower = append(follower, m)
			}
		}
		{
			m := &t4{key: key, client: client}
			if enableMap[m.Name()] {
				follower = append(follower, m)
			}
		}
	}

	return ServiceBalancer{
		userGetter:     user,
		followerGetter: follower,
	}
}
