package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string

	ref := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

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
	preSearch:
		for i := 0; i < len(line); i++ {
			if num, err := strconv.Atoi(string(line[i])); err == nil {
				sum += num * 10
				break preSearch
			}
			slice := line[i:]
			for num, val := range ref {
				if strings.HasPrefix(slice, num) {
					sum += val * 10
					break preSearch
				}
			}
		}
	postSearch:
		for i := len(line) - 1; i >= 0; i-- {
			if num, err := strconv.Atoi(string(line[i])); err == nil {
				sum += num
				break postSearch
			}
			slice := line[:i+1]
			for num, val := range ref {
				if strings.HasSuffix(slice, num) {
					sum += val
					break postSearch
				}
			}
		}
	}
	fmt.Println(sum)
}
