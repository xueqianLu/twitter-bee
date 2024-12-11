package types

import "github.com/xueqianLu/twitter-bee/models/apimodels"

type ServiceBackend interface {
	FollowerCount(req apimodels.FollowerCountRequest) (apimodels.FollowerCountResponse, error)
	FollowerList(req apimodels.FollowerListRequest) (apimodels.FollowerListResponse, error)
}

type TwitterAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	F2A      string `json:"f2a"`
}
