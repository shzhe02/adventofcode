package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func generateCombinations(totalLen int, groups []int, minOffset int) []string {
	// fmt.Println(totalLen)
	if totalLen < 0 {
		return []string{""}
	} else if len(groups) == 0 {
		return []string{strings.Repeat(".", totalLen)}
	}
	top := groups[0]
	maxOffset := totalLen
	combinations := make([]string, 0)
	for i := 1; i < len(groups); i++ {
		maxOffset -= groups[i] + 1
	}
	maxOffset -= groups[0]
	for offset := minOffset; offset <= maxOffset; offset++ {
		prefix := strings.Repeat(".", offset) + strings.Repeat("#", top)
		suffixes := generateCombinations(totalLen-offset-top, groups[1:], 1)
		for _, suffix := range suffixes {
			combinations = append(combinations, prefix+suffix)
		}
	}
	return combinations
}

func checkValid(pattern, target string) bool {
	matches := true
	for i := 0; i < len(pattern); i++ {
		if rune(pattern[i]) != '?' {
			matches = matches && pattern[i] == target[i]
		}
	}
	return matches
}

func main() {
	records := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		records = append(records, line)
	}
	out := 0
	for _, record := range records {
		parts := strings.Split(record, " ")
		field := parts[0]
		groups := make([]int, 0)
		for _, v := range strings.Split(parts[1], ",") {
			num, _ := strconv.Atoi(v)
			groups = append(groups, num)
		}
		// crappy brute force solution, since input seems short
		combinations := generateCombinations(len(field), groups, 0)
		for _, combination := range combinations {
			if checkValid(field, combination) {
				out++
			}
		}
	}
	fmt.Println(out)
}
