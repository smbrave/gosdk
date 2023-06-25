package adapi

import (
	"fmt"
	"testing"
)

func TestSdk(t1 *testing.T) {
	fmt.Println("ok1")
	sdk := NewSdk("http://localhost:9281", "10020", "")
	res, err := sdk.GetAccountReport("", "", "")
	if err != nil {
		panic(err)
	}
	fmt.Println("ok2")
	fmt.Println(len(res))
	fmt.Printf("%+v", res)
	for i, r := range res {
		fmt.Println(i, r.Name)
	}

	fmt.Println("ok3")
}
