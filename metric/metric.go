package metric

import (
	"time"
)

func Metric(name string, value float64, tag map[string]string) {
	if serv == nil {
		return
	}
	serv.add(&metric{
		Metric:    name,
		Value:     value,
		Tags:      tag,
		Timestamp: time.Now().Unix(),
	})
}
