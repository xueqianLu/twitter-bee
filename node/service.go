package node

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/xueqianLu/twitter-bee/config"
	"github.com/xueqianLu/twitter-bee/models/apimodels"
)

type handleService struct {
	n    *Node
	conf *config.Config
}

func (h handleService) FollowerCount(req apimodels.FollowerCountRequest) (apimodels.FollowerCountResponse, error) {
	res := apimodels.FollowerCountResponse{
		Count: 0,
	}
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

			res.Count = info.Followers
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
	for _, getter := range h.n.getBalancer.GetRandomFollowerGetter() {
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
