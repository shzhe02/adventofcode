package main

import (
	"bufio"
	"fmt"
	"os"
)

func calculateLoad(platform []string) int {
	totalLoad := 0
	for idx, row := range platform {
		load := len(platform) - idx
		for _, square := range row {
			if square == 'O' {
				totalLoad += load
			}
		}
	}
	return totalLoad
}

func tiltUp(platform []string) []string {
	out := make([]string, len(platform))
	for col := range platform[0] {
		cursor := 0
		for row := range platform {
			if platform[row][col] == 'O' {
				out[cursor] += "O"
				cursor++
			} else if platform[row][col] == '#' {
				// bring cursor up until row
				for ; cursor < row; cursor++ {
					out[cursor] += "."
				}
				out[cursor] += "#"
				cursor++
			}
		}
		for ; cursor < len(platform); cursor++ {
			out[cursor] += "."
		}
	}
	return out
}

// rotate clockwise
func rotateC(platform []string) []string {
	out := make([]string, len(platform[0]))
	for col := range platform[0] {
		for row := range platform {
			out[col] += string(platform[len(platform)-row-1][col])
		}
	}
	return out
}

// after doing enough rotations, the load will cycle across 7 values.
func findCycle(platform []string, count int) map[int]int {
	cycles := make(map[int]int) // cycle % 7 -> load
	streak := 0
	out := platform
	for i := 1; ; i++ {
		for j := 0; j < 4; j++ {
			out = tiltUp(out)
			out = rotateC(out)
		}
		load := calculateLoad(out)
		if cycles[i%7] != load {
			cycles[i%7] = load
			streak = 0
		} else {
			streak++
		}
		if streak > 17 {
			break
		}
	}
	return cycles
}

func main() {
	platform := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		platform = append(platform, line)
	}
	out := findCycle(platform, 10000)
	fmt.Println(out[1000000000%7])
}
