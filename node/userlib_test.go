package node

import "testing"

func TestGetUserLib(t *testing.T) {
	path := "../userlib.json"
	res := getUserLib(path)
	if len(res) == 0 {
		t.Errorf("getUserLib() = %v, want %v", res, "not empty")
	} else {
		t.Logf("get user count %v ", len(res))
	}
}
