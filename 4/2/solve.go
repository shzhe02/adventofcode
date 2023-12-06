package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	splitFn := func(c rune) bool { return c == ' ' }
	var input []string // input
	var cards []int    // wins per type of card
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		input = append(input, line)
	}
	totalTypes := len(input)
	numCards := make([]int, totalTypes) // total number of each type of card
	for i := range numCards {
		numCards[i] = 1
	}
	for _, line := range input {
		wins := 0
		winningNumbers := make(map[int]bool)
		numbers := strings.Split(strings.Split(line, ": ")[1], " | ")
		for _, v := range strings.FieldsFunc(numbers[0], splitFn) {
			val, _ := strconv.Atoi(v)
			winningNumbers[val] = true
		}
		for _, v := range strings.FieldsFunc(numbers[1], splitFn) {
			val, _ := strconv.Atoi(v)
			if winningNumbers[val] {
				wins++
			}
		}
		cards = append(cards, wins)
	}
	for idx, quantity := range numCards {
		for i := idx + 1; i < min(totalTypes, idx+1+cards[idx]); i++ {
			numCards[i] += quantity
		}
	}
	sum := 0
	for _, e := range numCards {
		sum += e
	}
	fmt.Println(sum)
}
