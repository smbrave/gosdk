package adminapi

import (
	"encoding/json"
	"fmt"
	"github.com/smbrave/gosdk/util"
	"net/url"
)

type Message struct {
	token   string
	sender  string
	address string
}

func NewMessage(addreess string, token string, sender string) *Message {
	if addreess == "" {
		addreess = "http://127.0.0.1:9281"
	}
	return &Message{
		address: addreess,
		token:   token,
		sender:  sender,
	}
}

func (m *Message) Send(receiver, content string) error {
	reqUrl := fmt.Sprintf("%s/admin/message/send?sender=%s&receiver=%s&content=%s", m.address, m.sender, receiver, url.QueryEscape(content))
	body, err := util.HttpGet(reqUrl, map[string]string{
		"x-token": m.token,
	})
	if err != nil {
		return err
	}
	var rsp util.Response
	if err := json.Unmarshal(body, &rsp); err != nil {
		return err
	}
	if rsp.Code != 0 {
		return fmt.Errorf("%d:%s", rsp.Code, rsp.Message)
	}
	return nil
}
