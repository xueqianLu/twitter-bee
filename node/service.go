package node

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru"
	log "github.com/sirupsen/logrus"
	"github.com/xueqianLu/twitter-bee/config"
	"github.com/xueqianLu/twitter-bee/models/apimodels"
	"github.com/xueqianLu/twitter-bee/types"
)

var (
	cache, _ = lru.New(100)
)

type handleService struct {
	n    *Node
	conf *config.Config
}

func (h handleService) UserProfile(req apimodels.UserInfoRequest) (apimodels.UserInfoResponse, error) {
	res := apimodels.UserInfoResponse{}
	for _, getter := range h.n.getBalancer.GetRandomUserGetter() {
		info, err := getter.GetUserInfo(req.UserName)
		if err != nil {
			continue
		} else {
			log.WithFields(log.Fields{
				"user":      req.UserName,
				"userid":    info.UserId,
				"followers": info.Followers,
			}).Info("GetProfile success")

			res.Name = info.Username
			res.Follower = info.Followers
			res.Id = info.UserId
			return res, nil
		}
	}
	return res, fmt.Errorf("can not get the follower count")

}

func (h handleService) FollowerList(req apimodels.FollowerListRequest) (apimodels.FollowerListResponse, error) {
	res := apimodels.FollowerListResponse{
		List: make([]apimodels.FollowerObj, 0),
		Next: "",
	}
	var slice []types.RAPIFollowerGetter
	if req.Cursor == "" {
		slice = h.n.getBalancer.GetRandomFollowerGetter()
	} else {
		if v, ok := cache.Get(req.Cursor); ok {
			idx := v.(int)
			slice = h.n.getBalancer.GetFollowerGetterByStart(idx)
		} else {
			return res, fmt.Errorf("invalid cursor")
		}
	}
	for _, getter := range slice {
		data, err := getter.GetFollowerIDs(req)
		if err != nil {
			continue
		} else {
			return data, nil
		}
	}
	return res, fmt.Errorf("can not get the follower list")
}

func newService(n *Node, conf *config.Config) *handleService {
	return &handleService{
		n:    n,
		conf: conf,
	}
}
