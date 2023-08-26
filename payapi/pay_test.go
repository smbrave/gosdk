package payapi

import (
	"fmt"
	"testing"
)

func TestNewPay(t *testing.T) {
	payApi := NewPay("http://u.batiao8.com", "")
	order, err := payApi.GetOrder("ZB_20230826122127_2YAxfg")
	if err != nil {
		panic(err)
	}
	fmt.Println(order)
}
