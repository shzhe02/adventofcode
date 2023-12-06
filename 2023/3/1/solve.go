package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	numMatcher := regexp.MustCompile("[0-9]+")
	symMatcher := regexp.MustCompile("[^0-9.]")
	sum := 0
	var schematic []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		schematic = append(schematic, line)
	}
	for y, line := range schematic {
		matches := numMatcher.FindAllStringIndex(line, -1)
		for _, match := range matches {
			num, _ := strconv.Atoi(line[match[0]:match[1]])
			block := ""
			left := max(match[0]-1, 0)
			right := min(match[1]+1, len(line))
			if y > 0 {
				block += schematic[y-1][left:right]
			}
			block += line[left:right]
			if y < len(schematic)-1 {
				block += schematic[y+1][left:right]
			}
			if a := symMatcher.Find([]byte(block)); a != nil {
				// symbol found within surroundings
				sum += num
			}
		}
	}
	fmt.Println(sum)
}
