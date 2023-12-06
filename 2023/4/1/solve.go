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
	splitFn := func(c rune) bool { return c == ' ' }
	sum := 0
	for scanner.Scan() {
		cardSum := 0
		winningNumbers := make(map[int]bool)
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		numbers := strings.Split(strings.Split(line, ": ")[1], " | ")
		for _, v := range strings.FieldsFunc(numbers[0], splitFn) {
			val, _ := strconv.Atoi(v)
			winningNumbers[val] = true
		}
		for _, v := range strings.FieldsFunc(numbers[1], splitFn) {
			val, _ := strconv.Atoi(v)
			if winningNumbers[val] {
				if cardSum == 0 {
					cardSum++
				} else {
					cardSum *= 2
				}
			}
		}
		sum += cardSum
	}
	fmt.Println(sum)
}
