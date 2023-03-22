package adapi

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSdk(t1 *testing.T) {
	sdk := NewSdk("http://ad.yic8.cn", "10020")
	res, err := sdk.Relation(1519)
	//res, err := sdk.GetBaiduAccountReport("", "")
	if err != nil {
		panic(err)
	}
	r, _ := json.Marshal(res)
	fmt.Println(string(r))
}
