package adapi

import (
	"errors"
	"strings"
)

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Request struct {
	Channel string            //安装包的渠道
	Version string            //安装包版本
	Os      string            //手机系统类别 android、ioss
	Ip      string            //客户端的外网ip
	Ua      string            //客户端的user-agent
	Model   string            //客户端的手机型号，如：NOH-AN00
	Idfa    string            //客户端的广告id，ios时候有效
	Oaid    string            //客户端的广告id，android时有效
	Imei    string            //设备唯一识别码
	Extra   map[string]string //其他额外数据
	Active  bool              // 是否直接激活
}

type Result struct {
	AdId   int64             `json:"adId"`
	Source string            `json:"source"`
	Extra  map[string]string `json:"extra"`
}

func (c *Request) Check() error {
	if strings.ToLower(c.Os) != "ios" && c.Channel == "" {
		return errors.New("channel must set")
	}
	if c.Version == "" {
		return errors.New("version must set")
	}
	return nil
}

type AccountReport struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Day       string  `json:"day"`
	Cost      float64 `json:"cost"`
	Show      int64   `json:"show"`
	Click     int64   `json:"click"`
	Download  int64   `json:"download"`
	Active    int64   `json:"active"`
	Pay       int64   `json:"pay"`
	PayAmount float64 `json:"payAmount"`
}
