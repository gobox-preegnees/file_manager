package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	Debug bool `yaml:"debug"`
	App   struct {
		Addr string `yaml:"addr"`
	}
	Kafka struct {
		Addr     []string `yaml:"addr"`
		Producer struct {
			ErrTopic  string `yaml:"err_topic"`
			Attempts  int    `yaml:"attempts"`
			Timeout   int    `yaml:"timeout"`
			Sleeptime int    `yaml:"sleeptime"`
		}
		Consumer struct {
			StateTopic string `yaml:"state_topic"`
			GroupId    string `yaml:"group_id"`
			Partition  int    `yaml:"partition"`
		}
	}
	Pg struct {
		Url string `yaml:"url"`
	}
}

var once sync.Once
var cnf config

func GetConfig(path string) *config {

	once.Do(func() {
		if err := cleanenv.ReadConfig(path, &cnf); err != nil {
			panic(err)
		}
	})
	return &cnf
}
