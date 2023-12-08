package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sections := make([][]string, 0)
scanSections:
	for {
		count := 0
		section := make([]string, 0)
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) == 0 {
				if count == 0 {
					break scanSections
				}
				break
			}
			section = append(section, line)
			count++
		}
		sections = append(sections, section)
	}
	path := sections[0][0]
	guide := make(map[string][]string)
	for _, node := range sections[1] {
		dirs := []string{node[7:10], node[12:15]}
		guide[node[:3]] = dirs
	}
	steps := 0
	curr := "AAA"
	for curr != "ZZZ" {
		for _, dir := range path {
			if dir == 'L' {
				curr = guide[curr][0]
			} else if dir == 'R' {
				curr = guide[curr][1]
			}
			steps++
			if curr == "ZZZ" {
				break
			}
		}
	}
	fmt.Println(steps)
}
