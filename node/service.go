package node

import (
	"github.com/xueqianLu/twitter-bee/config"
	"github.com/xueqianLu/twitter-bee/models/apimodels"
)

type handleService struct {
	n    *Node
	conf *config.Config
}

func (h handleService) FollowerCount(req apimodels.FollowerCountRequest) (apimodels.FollowerCountResponse, error) {
	h.n.spider.GetProfile(req.UserId)
	//TODO implement me
	panic("implement me")
}

func (h handleService) FollowerList(req apimodels.FollowerListRequest) (apimodels.FollowerListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func newService(n *Node, conf *config.Config) *handleService {
	return &handleService{
		n:    n,
		conf: conf,
	}
}
