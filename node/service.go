package node

import (
	"context"
	"encoding/base64"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/xueqianLu/twitter-bee/config"
	"github.com/xueqianLu/twitter-bee/models/apimodels"
	"strings"
	"time"
)

type handleService struct {
	n    *Node
	conf *config.Config
}

func (h handleService) FollowerCount(req apimodels.FollowerCountRequest) (apimodels.FollowerCountResponse, error) {
	tryCount := 0
	res := apimodels.FollowerCountResponse{
		Count: 0,
	}
	spider := h.n.spider
	if h.n.available == false {
		return res, fmt.Errorf("current service is not available")
	}

	for {
		if tryCount >= 2 {
			return res, fmt.Errorf("can not get the %v follower count", req.UserId)
		}

		fc, err := spider.GetProfile(req.UserId)

		if err != nil {
			spider.GetGuestToken()
			tryCount++
			fmt.Printf("GetProfile %s failed, err:%s\n", req.UserId, err.Error())
			time.Sleep(time.Second * time.Duration(tryCount))
			continue
		}

		res.Count = fc.FollowersCount

		return res, nil
	}
}

func (h handleService) FollowerList(req apimodels.FollowerListRequest) (apimodels.FollowerListResponse, error) {
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	res := apimodels.FollowerListResponse{
		List: make([]apimodels.FollowerObj, 0),
		Next: "",
	}
	if err := h.n.RateLimit.Wait(ctx); err != nil {
		return res, fmt.Errorf("wait rate limit failed with error:%s", err.Error())
	}

	var (
		success bool = false
		try          = 0
		spider       = h.n.spider
	)

	for !success && try < 1 {
		followers, next, err := spider.FetchFollowers(req.User, 1000, req.Cursor)
		if err != nil {
			try++
			log.WithError(err).Error("FetchFollowers failed")
			continue
		}

		success = true
		res.Next = next
		for _, v := range followers {
			sDec, _ := base64.StdEncoding.DecodeString(v.UserID)
			userId, _ := strings.CutPrefix(string(sDec), "User:")
			res.List = append(res.List, apimodels.FollowerObj{
				ID:       userId,
				Name:     v.Name,
				UserName: v.Username,
			})
		}
	}
	if !success {
		return res, fmt.Errorf("can not get the follower list")
	}
	return res, nil
}

func newService(n *Node, conf *config.Config) *handleService {
	return &handleService{
		n:    n,
		conf: conf,
	}
}
