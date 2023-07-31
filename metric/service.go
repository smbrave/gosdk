package metric

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

type metric struct {
	Timestamp int64             `json:"timestamp"`
	Metric    string            `json:"metric"`
	Value     float64           `json:"value"`
	Tags      map[string]string `json:"tags"`
}

type service struct {
	metrics chan *metric
	megers  map[string]*metric
	config  *Config
}

func (s *service) run() {
	timer := time.NewTicker(time.Duration(s.config.Interval) * time.Second)
	for {
		select {
		case m := <-s.metrics:
			s.process(m)

		case <-timer.C:
			s.report()
		}
	}
}

func (s *service) sortMap(tag map[string]string) string {
	arr := make([]string, 0)
	for k, v := range tag {
		arr = append(arr, fmt.Sprintf("%s=%s", k, v))
	}
	sort.Strings(arr)
	return strings.Join(arr, ":")
}

func (s *service) process(m *metric) {
	if s.megers == nil {
		s.megers = make(map[string]*metric)
	}
	key := m.Metric
	if m.Tags != nil {
		key += "_" + s.sortMap(m.Tags)
	}

	if v, ok := s.megers[key]; ok {
		v.Value += m.Value
		v.Timestamp = m.Timestamp
		return
	}
	s.megers[key] = m
}

func (s *service) report() {
	if s.megers == nil {
		return
	}
	metrics := make([]*metric, 0)
	for _, v := range s.megers {
		metrics = append(metrics, v)
	}
	reqUrl := fmt.Sprintf("%s/opentsdb/put", serv.config.Address)

	reqBody, _ := json.Marshal(metrics)
	resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Printf("http.Post error :%s", err.Error())
		return
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf(" ioutil.ReadAll error :%s", err.Error())
		return
	}
	result := make(map[string]interface{})
	if err := json.Unmarshal(rspBody, &result); err != nil {
		log.Printf("json result : %s", string(rspBody))
		return
	}

	fail := cast.ToInt(result["fail"])
	if fail != 0 {
		log.Printf("http result : %s", string(rspBody))
		return
	}
	s.megers = nil
}

func (s *service) add(m *metric) {
	select {
	case s.metrics <- m:
	default:
		fmt.Println("chan is full")
	}
}
