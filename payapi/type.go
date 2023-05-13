package payapi

type OrderUser struct {
	UserId      string `json:"userId"`
	UserName    string `json:"userName"`
	Source      string `json:"source"`
	MobileBrand string `json:"brand"`
	MobileModel string `json:"model"`
	Platform    string `json:"platform"`
	Channel     string `json:"channel"`
	Version     string `json:"version"`
}
type OrderGoods struct {
	GoodsId   string `json:"goodsId"`
	GoodsName string `json:"goodsName"`
	Source    string `json:"source"`
	Autopay   string `json:"autopay"`
}

type CreateOrderReq struct {
	OutTradeNo string      `json:"outTradeNo"`
	PayType    string      `json:"payType"`
	PayPrice   int64       `json:"payPrice"`
	NotifyUrl  string      `json:"notifyUrl"`
	User       *OrderUser  `json:"user,omitempty"`
	Goods      *OrderGoods `json:"goods,omitempty"`
}

type CommonResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
