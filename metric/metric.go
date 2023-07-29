package metric

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Metric struct {
	Timestamp int64             `json:"timestamp"`
	Metric    string            `json:"metric"`
	Value     float64           `json:"value"`
	Tags      map[string]string `json:"tags"`
}

type MetricService struct {
	address string
}

func NewMetricService(address string) *MetricService {
	return &MetricService{
		address: address,
	}
}

func (s *MetricService) Report(m *Metric) error {
	reqUrl := fmt.Sprintf("%s/opentsdb/put", s.address)

	reqBody, _ := json.Marshal(m)
	resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	result := make(map[string]interface{})
	if err := json.Unmarshal(rspBody, &result); err != nil {
		return err
	}
	fail := result["fail"].(int)
	if fail != 0 {
		return errors.New(string(rspBody))
	}
	return nil
}

func (s *MetricService) ReportBatch(m []*Metric) error {
	return nil
}
