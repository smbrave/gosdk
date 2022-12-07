package adapi

import "errors"

type Client struct {
	AppId int64  `form:"appId"` //每一个产品唯一分配一个appid
	Os    string `form:"os"`    //手机系统类别 android、ioss
	Ip    string `form:"ip"`    //客户端的外网ip
	Ua    string `form:"ua"`    //客户端的user-agent
	Model string `form:"model"` //客户端的手机型号，如：NOH-AN00
	Idfa  string `form:"idfa"`  //客户端的广告id，ios时候有效
	Oaid  string `form:"oaid"`  //客户端的广告id，android时有效
	Imei  string `form:"imei"`  //设备唯一识别码
}

type Result struct {
	AdId   int64  `json:"adId"`
	Source string `json:"source"`
}

func (c *Client) Check() error {
	if c.AppId == 0 {
		return errors.New("appid must set")
	}
	return nil
}
