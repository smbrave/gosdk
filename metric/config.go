package metric

var (
	serv *service
)

type Config struct {
	Address  string
	Interval int
}

func NewConfg() *Config {
	return &Config{
		Address:  "10.0.1.15:17000",
		Interval: 10,
	}
}

func Init(c *Config) error {
	if serv != nil {
		return nil
	}

	serv = &service{
		config:  c,
		metrics: make(chan *metric, 100000),
	}
	go serv.run()
	return nil
}
