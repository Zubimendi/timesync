package config

import "sync"

var once sync.Once
var instance *Config

type Config struct {
	DefaultTimezone string
}

func LoadConfig() * Config{
	once.Do(func(){
		instance = &Config{
			DefaultTimezone: "UTC",
		}
	})
	return instance
}
