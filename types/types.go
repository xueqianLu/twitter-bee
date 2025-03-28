package types

import (
	"github.com/xueqianLu/twitter-bee/models/apimodels"
	"net/http"
)

type ServiceBackend interface {
	FollowerCount(req apimodels.FollowerCountRequest) (apimodels.FollowerCountResponse, error)
	FollowerList(req apimodels.FollowerListRequest) (apimodels.FollowerListResponse, error)
}

type TwitterAccount struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	F2A       string `json:"f2a"`
	Token     string `json:"token"`
	CSRFToken string `json:"csrf"`
}

type UserProfile struct {
	Username  string `json:"username"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
	UserId    string `json:"id"`
}

type RAPIUserGetter interface {
	Setup(client *http.Client) error
	GetUserInfo(name string) (UserProfile, error)
}

type RAPIFollowerGetter interface {
	Setup(client *http.Client) error
	GetFollowerIDs(req apimodels.FollowerListRequest) (apimodels.FollowerListResponse, error)
}
