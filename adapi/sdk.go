package adapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

	u := fmt.Sprintf("%s/api/client/active?appId=%d&os=%s&ip=%s&model=%s&idfa=%s&oaid=%s&imei=%s",
		s.address, c.AppId, strings.ToLower(c.Os), c.Ip, url.QueryEscape(c.Model), c.Idfa, c.Oaid, c.Imei)
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
