package main

import (
	"fmt"
	"time-manage-go/internal/config"
)

func main() {
	//	TODO: init config
	cfg := config.MustLoad()

	fmt.Println(cfg)

	//	TODO: init logger

	//	TODO: init storage

	//	TODO: init router

	//	TODO: run server
}
