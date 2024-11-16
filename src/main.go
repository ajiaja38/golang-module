package main

import (
	"fmt"
	"go-again/src/service"
)

func main() {
	fmt.Println("Hello Main Golang!")

	input := 1000000

	service.LoopBillion(input)
}
