package adminapi

import (
	"fmt"
	"strings"
	"testing"
)

func TestMessage_Send(t *testing.T) {

	s := NewMessage("http://u.batiao8.com", "", "alarm")

	mess := make([]string, 0)
	mess = append(mess, "加发到")
	mess = append(mess, "asdfasdfadfa")
	mess = append(mess, "@3q452345~!@@#$$%^&*())")
	err := s.Send("jiangyong", strings.Join(mess, "\n"))
	fmt.Println(err)
}
