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

func findGear(line string, y, l, r int) int {
	for idx, r := range line[l:r] {
		if r == '*' {
			return y*len(line) + l + idx
		}
	}
	return -1 // gear not found
}

func processGear(pos, num int, gears, gearCount map[int]int) {
	if pos == -1 {
		return
	}
	if gearCount[pos] == 0 {
		gears[pos] = num
	} else if gearCount[pos] == 1 {
		gears[pos] *= num
	}
	gearCount[pos]++
}

func main() {
	numMatcher := regexp.MustCompile("[0-9]+")
	gears := make(map[int]int)     // gear position -> gear ratio
	gearCount := make(map[int]int) // gear position -> number of numbers around it
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
			left := max(match[0]-1, 0)
			right := min(match[1]+1, len(line))
			if y > 0 {
				pos := findGear(schematic[y-1], y-1, left, right)
				processGear(pos, num, gears, gearCount)
			}
			pos := findGear(line, y, left, right)
			processGear(pos, num, gears, gearCount)
			if y < len(schematic)-1 {
				pos = findGear(schematic[y+1], y+1, left, right)
				processGear(pos, num, gears, gearCount)
			}
		}
	}
	for k, v := range gears {
		if gearCount[k] == 2 {
			sum += v
		}
	}
	fmt.Println(sum)
}
