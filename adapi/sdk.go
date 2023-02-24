package adapi

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Sdk struct {
	address string
	appId   string
}

func NewSdk(address string, appId string) *Sdk {
	if address == "" {
		address = "http://127.0.0.1:9280"
	}

	return &Sdk{
		appId:   appId,
		address: address,
	}
}

func (s *Sdk) httpGet(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	type rsp_t struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	var rsp rsp_t
	if err := json.Unmarshal(body, &rsp); err != nil {
		return err
	}
	if rsp.Code != 0 {
		return fmt.Errorf("%d:%s", rsp.Code, rsp.Message)
	}
	return nil

}

func (s *Sdk) httpMatchGet(url string) (*Result, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type rsp_t struct {
		Code    int     `json:"code"`
		Message string  `json:"message"`
		Data    *Result `json:"data"`
	}

	var rsp rsp_t
	if err := json.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	if rsp.Code != 0 {
		return nil, fmt.Errorf("%d:%s", rsp.Code, rsp.Message)
	}
	return rsp.Data, nil
}

func (s *Sdk) Match(c *Request) (*Result, error) {
	if err := c.Check(); err != nil {
		return nil, err
	}

	params := url.Values{}

	params.Add("appId", s.appId)
	params.Add("os", strings.ToLower(c.Os))
	params.Add("ip", c.Ip)
	params.Add("ua", c.Ua)
	params.Add("model", c.Model)
	params.Add("idfa", c.Idfa)
	params.Add("oaid", c.Oaid)
	params.Add("imei", c.Imei)
	params.Add("channel", c.Channel)
	params.Add("version", c.Version)
	params.Add("active", strconv.FormatBool(c.Active))
	if c.Extra != nil {
		extra, _ := json.Marshal(c.Extra)
		params.Add("extra", string(extra))
	}

	u := fmt.Sprintf("%s/api/client/match?%s", s.address, params.Encode())
	return s.httpMatchGet(u)
}

func (s *Sdk) Active(adId int64, extra map[string]string) error {
	params := url.Values{}
	params.Add("adId", strconv.FormatInt(adId, 10))
	if extra != nil {
		ex, _ := json.Marshal(extra)
		params.Add("extra", string(ex))
	}

	url := fmt.Sprintf("%s/api/client/active?%s",
		s.address, params.Encode())

	return s.httpGet(url)
}

/*
	extra存放扩展数据，注册相关信息都填上
	unionid、phone、openid、nickname
*/

func (s *Sdk) Register(adId int64, extra map[string]string) error {
	params := url.Values{}
	params.Add("adId", strconv.FormatInt(adId, 10))
	if extra != nil {
		ex, _ := json.Marshal(extra)
		params.Add("extra", string(ex))
	}

	url := fmt.Sprintf("%s/api/client/register?%s",
		s.address, params.Encode())

	return s.httpGet(url)
}

/*
	extra存放扩展数据，支付相关信息都填上
	支付金额：payFee （单位分）
	支付方式：payType（weixin、alipay、apple）
	支付位置：payLocation （内容自定义）
	支付商品：goodsId、goodsName
*/

func (s *Sdk) Pay(adId int64, extra map[string]string) error {
	params := url.Values{}
	params.Add("adId", strconv.FormatInt(adId, 10))
	if extra != nil {
		ex, _ := json.Marshal(extra)
		params.Add("extra", string(ex))
	}

	url := fmt.Sprintf("%s/api/client/pay?%s",
		s.address, params.Encode())

	return s.httpGet(url)
}

func (s *Sdk) GetOceanAccountReport(startDay, endDay string) ([]*AccountReport, error) {
	if startDay == "" && endDay == "" {
		startDay = time.Now().Format("2006-01-02")
		endDay = startDay
	}
	url := fmt.Sprintf("%s/admin/ocean/account/report?appId=%s&startDay=%s&endDay=%s", s.address, s.appId, startDay, endDay)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type rsp_t struct {
		BaseResponse
		Data []map[string]interface{} `json:"data"`
	}

	var rsp rsp_t
	if err := json.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	if rsp.Code != 0 {
		return nil, fmt.Errorf("%d:%s", rsp.Code, rsp.Message)
	}
	results := make([]*AccountReport, 0)
	for _, data := range rsp.Data {
		r := new(AccountReport)
		r.Id = cast.ToString(data["id"])
		r.Name = cast.ToString(data["name"])
		r.Day = cast.ToString(data["day"])
		r.Cost = cast.ToFloat64(data["cost"])
		r.Show = cast.ToInt64(data["show"])
		r.Click = cast.ToInt64(data["click"])
		r.Download = cast.ToInt64(data["download"])
		r.Active = cast.ToInt64(data["active"])
		r.Pay = cast.ToInt64(data["pay"])
		r.PayAmount = cast.ToFloat64(data["payAmount"])
		results = append(results, r)
	}
	return results, nil
}

func (s *Sdk) GetBaiduAccountReport(startDay, endDay string) ([]*AccountReport, error) {
	if startDay == "" && endDay == "" {
		startDay = time.Now().Format("2006-01-02")
		endDay = startDay
	}
	url := fmt.Sprintf("%s/admin/baidu/account/report?appId=%s&startDay=%s&endDay=%s", s.address, s.appId, startDay, endDay)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type rsp_t struct {
		BaseResponse
		Data []map[string]interface{} `json:"data"`
	}

	var rsp rsp_t
	if err := json.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	if rsp.Code != 0 {
		return nil, fmt.Errorf("%d:%s", rsp.Code, rsp.Message)
	}
	results := make([]*AccountReport, 0)
	for _, data := range rsp.Data {
		r := new(AccountReport)
		r.Id = cast.ToString(data["userId"])
		r.Name = cast.ToString(data["userName"])
		r.Day = cast.ToString(data["day"])
		r.Cost = cast.ToFloat64(data["cost"])
		r.Show = cast.ToInt64(data["impression"])
		r.Click = cast.ToInt64(data["click"])
		r.Download = cast.ToInt64(data["download"])
		r.Active = cast.ToInt64(data["active"])
		r.Pay = cast.ToInt64(data["pay"])
		r.PayAmount = cast.ToFloat64(data["payAmount"])
		results = append(results, r)
	}
	return results, nil
}
