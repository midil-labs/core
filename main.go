package main

import (
	"fmt"
	"github.com/midil-labs/core/pkg/config"
)


func main() {
	cfg, err := config.LoadConfig("./pkg/config")
	if err != nil {
		fmt.Printf("error loading config: %v\n", err)
	}
	fmt.Printf("config: % #v\n", cfg)
}