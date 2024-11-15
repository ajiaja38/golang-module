package main

import (
	"fmt"
	"go-again/src/service"
)

func main() {
	fmt.Println("Hello Main Golang!")

	input := []int{7, 8, 5, 2}

	resultSumTwo := service.SumTwo(input, 9)
	fmt.Println(resultSumTwo)
}
