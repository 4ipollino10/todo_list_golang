package main

import (
	"fmt"
	"go_demo/cmd/bootstrap/config"
)

func main() {
	cfg := config.MustLoadConfig()
	fmt.Printf("%#v\n", cfg)
}
