package payapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/smbrave/gosdk/util"
)

var (
	PayTypeWeixin = "weixin"
	PayTypeAlipay = "alipay"
)

type Pay struct {
	address string
	token   string
}

func NewPay(address string, token string) *Pay {
	if address == "" {
		address = "http://127.0.0.1:9281"
	}
	return &Pay{
		address: address,
		token:   token,
	}
}

func (p *Pay) CreateOrder(order *CreateOrderReq) (map[string]interface{}, error) {
	if order.PayType == "" {
		errors.New("payType is nil")
	}
	if order.PayPrice <= 0 {
		return nil, errors.New("payPrice is nil")
	}
	reqBody, _ := json.Marshal(order)

	result, err := util.HttpPostJson(p.address+"/api/pay/order", map[string]string{
		"x-token": p.token,
	}, reqBody)

	if err != nil {
		return nil, err
	}
	var rsp CommonResponse
	if err := json.Unmarshal([]byte(result), &rsp); err != nil {
		return nil, err
	}
	if rsp.Code != 0 {
		return nil, fmt.Errorf("%d:%s", rsp.Code, rsp.Message)
	}

	return rsp.Data, nil
}

func (p *Pay) GetOrder(outTradeNo string) (map[string]interface{}, error) {
	if outTradeNo == "" {
		errors.New("outTradeNo is nil")
	}

	reqUrl := fmt.Sprintf("%s/api/pay/order?outTradeNo=%s", p.address, outTradeNo)
	result, err := util.HttpGet(reqUrl, map[string]string{
		"x-token": p.token,
	})

	if err != nil {
		return nil, err
	}
	var rsp CommonResponse
	if err := json.Unmarshal([]byte(result), &rsp); err != nil {
		return nil, err
	}
	if rsp.Code != 0 {
		return nil, fmt.Errorf("%d:%s", rsp.Code, rsp.Message)
	}

	return rsp.Data, nil
}
