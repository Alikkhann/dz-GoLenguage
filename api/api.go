package api

import (
	"fmt"
	"myproject/bins"
	"myproject/config"
)

func GetKey() {
	key := config.NewConfig()
	fmt.Println(key.Key)
}

func Print(*config.Config, bins.BinList) {
	
}

