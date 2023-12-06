package main

import (
	"fmt"
)

func main() {
	test := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(test); i++ {
		fmt.Println(test[i])
		test = append(test, 1)
	}
}
