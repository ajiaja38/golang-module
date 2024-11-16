package service

import (
	"fmt"
	"time"
)

type Student struct {
	Name, Address string
	Age, PostCode int
}

func LoopStudent(students []Student) {
	for _, student := range students {
		fmt.Println(student)
	}
}

func FizzBuzz(input int) {
	for i := 1; i <= input; i++ {
		if i%3 == 0 || i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func SumTwo(input []int, target int) [2]int {
	var result [2]int

	for i := 0; i < len(input); i++ {
		for j := i + 1; i < len(input); j++ {
			if input[i]+input[j] == target {
				result[0] = i
				result[1] = j

				return result
			}
		}
	}

	return result
}

func LoopBillion(input int) {
	value := time.Now()

	for i := 1; i <= input; i++ {
		fmt.Println(i)
	}

	elapse := time.Since(value)
	fmt.Println(elapse)
}
