package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type config struct {
	App   string `yaml:"app"`
	Pg    string `yaml:"pg"`
	Kafka string `yaml:"kafka"`
	Debug bool   `yaml:"debug"`
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
