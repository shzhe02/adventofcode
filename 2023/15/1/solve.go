package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var data string
	for scanner.Scan() {
		in := scanner.Text()
		if len(in) == 0 {
			break
		}
		data = in
	}
	totalHash := 0
	curr := 0
	for _, c := range data {
		if c == ',' {
			totalHash += curr
			curr = 0
		} else {
			curr = ((curr + int(c)) * 17) % 256
		}
	}
	totalHash += curr
	fmt.Println(totalHash)
}
