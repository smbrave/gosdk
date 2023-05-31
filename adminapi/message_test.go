package message

import (
	"fmt"
	"testing"
)

func TestMessage_Send(t *testing.T) {

	s := NewMessage("http://u.batiao8.com", "", "bid")
	err := s.Send("jiangyong", "tes")
	fmt.Println(err)
}
