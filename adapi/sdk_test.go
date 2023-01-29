package adapi

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSdk(t1 *testing.T) {
	sdk := NewSdk("https://ad.xinzhibid.com", "10000")
	res, err := sdk.GetOceanAccountReport("", "")
	//res, err := sdk.GetBaiduAccountReport("", "")
	if err != nil {
		panic(err)
	}
	r, _ := json.Marshal(res)
	fmt.Println(string(r))
}
