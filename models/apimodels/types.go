package apimodels

type UserInfoRequest struct {
	UserName string `json:"username"`
}

type UserInfoResponse struct {
	Follower int    `json:"follower"`
	Name     string `json:"name"`
	Id       string `json:"id"`
}

type FollowerObj struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
}

type FollowerListRequest struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Cursor string `json:"cursor"`
}

type FollowerListResponse struct {
	List []FollowerObj `json:"list"`
	Next string        `json:"next"`
}

type BaseResponse struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
