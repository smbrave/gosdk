package adapi

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSdk(t1 *testing.T) {
	type tt struct {
		A string            `json:"a"`
		B map[string]string `json:"b"`
	}
	var t tt
	s := `{"a":"1", "b":{"1":"2","3":"4"}}`
	if err := json.Unmarshal([]byte(s), &t); err != nil {
		panic(err)
	}
	res, _ := json.Marshal(t)
	fmt.Println(string(res))
}
