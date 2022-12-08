package adapi

import (
	"fmt"
	"testing"
)

func TestSdk(t *testing.T) {
	s := NewSdk("http://localhost:9280")
	res, err := s.Active(&Client{
		AppId:   10000,
		Version: "1.1.1",
		Channel: "test",
		Idfa:    "b306c7260f30c1af5d9b74e3e70c279e", // ios有效
		Os:      "ios",
		Ip:      "39.144.154.163",
		Model:   "iphone13,4",
		Oaid:    "826907af-09ff-40d0-af5d-2c32390a539f", //android有效
		Imei:    "e8776385f699ba294fc82cad5791a10d",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	if err := s.Register(res.AdId); err != nil {
		panic(err)
	}
	if err := s.Pay(res.AdId); err != nil {
		panic(err)
	}
}
