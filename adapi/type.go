package adapi

import (
	"errors"
	"strings"
)

type Client struct {
	AppId   int64             //分配一个appid
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

func (c *Client) Check() error {
	if c.AppId == 0 {
		return errors.New("appid must set")
	}
	if strings.ToLower(c.Os) != "ios" && c.Channel == "" {
		return errors.New("channel must set")
	}
	if c.Version == "" {
		return errors.New("version must set")
	}
	return nil
}
