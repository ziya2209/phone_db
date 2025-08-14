package main

import "github.com/ziya2209/goproject/phonedb/internal/controller"

func main() {
	if err := controller.Start(); err != nil {
		panic(err)
	}
}
