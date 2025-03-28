package client

import (
	"testing"
	"time"
)

func TestBeeClient_GetFollowerList(t *testing.T) {
	url := "http://localhost:8088"
	cli := NewBeeClient(url)
	info, err := cli.GetUserProfile("caduceus_cad")
	if err != nil {
		t.Error(err)
	} else {
		println("got user name:", info.Name, " id:", info.Id, " followers:", info.Follower)
	}
	var cursor = ""
	for i := 0; i < 5; i++ {
		res, err := cli.GetFollowerList(info.Name, info.Id, cursor)
		if err != nil {
			t.Error(err)
		}
		println("got follower list", len(res.List), "next", res.Next)
		cursor = res.Next
		time.Sleep(1 * time.Second)
	}
}
