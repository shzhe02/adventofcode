package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string

	var sum = 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, scanner.Text())
	}
	for _, line := range lines {
		// get first digit
		for i := 0; i < len(line); i++ {
			if num, err := strconv.Atoi(string(line[i])); err == nil {
				sum += num * 10
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if num, err := strconv.Atoi(string(line[i])); err == nil {
				sum += num
				break
			}
		}
	}
	fmt.Println(sum)
}
