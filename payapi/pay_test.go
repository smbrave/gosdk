package payapi

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewPay(t *testing.T) {
	pay := NewPay("http://127.0.0.1:9281", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ4Mzc1Njc2MjEsImlhdCI6MTY4Mzk2NzYyMSwiaXNzIjoiZGVsbC1qeSIsIm5iZiI6MTY4Mzk2NzU2MSwiQXBwSWQiOjEwMDIwfQ.P_JKbNjTvVDON901I_lT7ydJCaFwSnC-CrDO9eiEGRg")
	data, err := pay.CreateOrder(&CreateOrderReq{
		PayPrice:  1000,
		PayType:   PayTypeWeixin,
		NotifyUrl: "https://bid.yic8.cn/api/batiao",
		User: &OrderUser{
			UserId:      "1234",
			UserName:    "test",
			Source:      "baidu",
			MobileModel: "AT1000",
			MobileBrand: "Huawei",
			Platform:    "android",
			Version:     "1.2.3",
			Channel:     "bd_tg_01",
		},
		Goods: &OrderGoods{
			GoodsName: "测试",
			GoodsId:   "1234",
			Source:    "bootpage",
			Autopay:   true,
		},
	})
	if err != nil {
		panic(err)
	}
	res, _ := json.Marshal(data)
	fmt.Println(string(res))
}
