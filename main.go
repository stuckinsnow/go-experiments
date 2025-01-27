package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 20; i++ {
		fmt.Println(i)
		if i < 10 {
			fmt.Println("Hello")
		}
	}

	i := 0
	for i < 10 {
		fmt.Println("hello", i)
		i++
	}

	nums := []int{1, 2, 3}
	fmt.Println(nums)

outer:
	for i := 0; i < 3; i++ {

		fmt.Println("Hellooo")

		for j := 0; j < 3; j++ {
			if i == j {
				continue outer
			}
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}
}
