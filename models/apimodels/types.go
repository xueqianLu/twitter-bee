package apimodels

type FollowerCountRequest struct {
	UserId string `json:"user_id"`
}

type FollowerCountResponse struct {
	Count int `json:"count"`
}

type FollowerObj struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
}

type FollowerListRequest struct {
	User   string `json:"user"`
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
