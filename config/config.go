package config

import (
	"os"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	key := os.Getenv("KEY")
	//fmt.Println(key) ошибка была в ключе надо было добавить ковычки - ''
	return &Config{
		Key: key,
	}
}