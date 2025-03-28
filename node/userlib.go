package node

import (
	"encoding/json"
	"os"
)

func getKeyList(path string) []string {
	// read file from path and unmarshal json to slice.
	data, _ := os.ReadFile(path)
	var keylist []string
	_ = json.Unmarshal(data, &keylist)
	return keylist
}
