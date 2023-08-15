package metric

import (
	"math/rand"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	c := NewConfg()
	c.Address = "https://monitor.batiao8.com"
	Init(c)
	for i := 0; i < 100; i++ {
		Metric("test", float64(rand.Int()%100), map[string]string{
			"a": "b",
			"c": "d",
		})
		time.Sleep(time.Second)
	}
}

func TestMetric(t *testing.T) {
	c := NewConfg()
	c.Address = "https://monitor.batiao8.com"
	Init(c)

	Metric("test.test1.test2", 12, map[string]string{
		"a": "b",
		"c": "d",
	})

}
