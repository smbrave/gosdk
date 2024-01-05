package adminapi

import (
	"encoding/json"
	"fmt"
	"github.com/smbrave/gosdk/util"
)

type Application struct {
	token   string
	address string
}

func NewApplication(addreess string, token string) *Application {
	if addreess == "" {
		addreess = "http://127.0.0.1:9281"
	}
	return &Application{
		address: addreess,
		token:   token,
	}
}

func (m *Application) Login(username, password, appid string) error {
	reqUrl := fmt.Sprintf("%s/admin/app/login", m.address)
	params := make(map[string]interface{})
	params["username"] = username
	params["password"] = password
	params["appid"] = appid
	reqBody, _ := json.Marshal(params)
	body, err := util.HttpPostJson(reqUrl, map[string]string{
		"x-token": m.token,
	}, reqBody)

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
