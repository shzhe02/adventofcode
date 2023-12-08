package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(curr []string, match string) bool {
	state := true
	for _, v := range curr {
		state = state && strings.HasSuffix(v, match)
	}
	return state
}

func gcd(a, b uint64) uint64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b uint64) uint64 {
	return (a * b) / gcd(a, b)
}

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
	guideL := make(map[string]string)
	guideR := make(map[string]string)
	for _, node := range sections[1] {
		guideL[node[:3]] = node[7:10]
		guideR[node[:3]] = node[12:15]
	}
	curr := make([]string, 0)
	for k := range guideL {
		if strings.HasSuffix(k, "A") {
			curr = append(curr, k)
		}
	}
	loopIdx := make([]int, 0, len(curr))
	loopLen := make([]int, 0, len(curr))
nodes:
	for _, node := range curr {
		steps := 0
		cursor := node
		visited := make(map[string]int)
		for {
			for idx, dir := range path {
				if dir == 'L' {
					cursor = guideL[cursor]
				} else {
					cursor = guideR[cursor]
				}
				if visited[cursor+string(idx)] != 0 {
					loopIdx = append(loopIdx, steps)
					loopLen = append(loopLen, steps-visited[cursor+string(idx)])
					continue nodes
				} else {
					visited[cursor+string(idx)] = steps
				}
				steps++
			}
		}
	}
	out := uint64(loopLen[0])
	for i := 1; i < len(loopLen); i++ {
		out = lcm(out, uint64(loopLen[i]))
		fmt.Println(out)
	}
}
