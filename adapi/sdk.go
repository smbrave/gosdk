package adapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Sdk struct {
	address string
}

func NewSdk(address string) *Sdk {
	if address == "" {
		address = "http://127.0.0.1:9280"
	}

	return &Sdk{
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

func (s *Sdk) Active(c *Client) (*Result, error) {
	if err := c.Check(); err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("appId", strconv.FormatInt(c.AppId, 10))
	params.Add("os", strings.ToLower(c.Os))
	params.Add("ip", c.Ip)
	params.Add("ua", c.Ua)
	params.Add("model", c.Model)
	params.Add("idfa", c.Idfa)
	params.Add("oaid", c.Oaid)
	params.Add("imei", c.Imei)
	params.Add("channel", c.Channel)
	params.Add("version", c.Version)
	if c.Extra != nil {
		extra, _ := json.Marshal(c.Extra)
		params.Add("extra", string(extra))
	}

	u := fmt.Sprintf("%s/api/client/active?%s", s.address, params.Encode())
	return s.httpMatchGet(u)
}

func (s *Sdk) Register(adId int64) error {
	url := fmt.Sprintf("%s/api/client/register?adId=%d",
		s.address, adId)
	return s.httpGet(url)
}

func (s *Sdk) Pay(adId int64) error {
	url := fmt.Sprintf("%s/api/client/pay?adId=%d",
		s.address, adId)
	return s.httpGet(url)
}
