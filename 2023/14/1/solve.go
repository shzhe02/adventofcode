package main

import (
	"bufio"
	"fmt"
	"os"
)

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
	totalLoad := 0
	for col := 0; col < len(platform[0]); col++ {
		load := len(platform)
		for row := 0; row < len(platform); row++ {
			// load at a certain square = len(platform) - row
			if platform[row][col] == '#' {
				load = len(platform) - row - 1
			} else if platform[row][col] == 'O' {
				totalLoad += load
				load--
			}
		}
	}
	fmt.Println(totalLoad)
}
