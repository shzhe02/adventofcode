package main

import (
	"bufio"
	"fmt"
	"os"
)

func findMirror(pattern []string) int { // lines above the mirror if found, 0 if not found.
	// mirror in the top half (first line is mirrored)
topSearch:
	for i := 1; i < len(pattern); i += 2 {
		if pattern[i] == pattern[0] {
			// mirror ends potentially found, check if correctly mirrored
			t := 1
			b := i - 1
			for t < b {
				if pattern[t] != pattern[b] {
					continue topSearch
				}
				t++
				b--
			}
			// scan passed, mirror found.
			return (i / 2) + 1
		}
	}
	// mirror in the bottom half (last line is mirrored)
bottomSearch:
	for i := len(pattern) - 2; i > 0; i -= 2 {
		if pattern[i] == pattern[len(pattern)-1] {
			t := i + 1
			b := len(pattern) - 2
			for t < b {
				if pattern[t] != pattern[b] {
					continue bottomSearch
				}
				t++
				b--
			}
			return len(pattern) - (((len(pattern) - 1 - i) / 2) + 1)
		}
	}
	return 0
}

func transpose(pattern []string) []string {
	output := make([]string, 0)
	for idx := range pattern[0] {
		line := ""
		for _, l := range pattern {
			line += string(l[idx])
		}
		output = append(output, line)
	}
	return output
}

func main() {
	patterns := make([][]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		pattern := make([]string, 0)
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) == 0 {
				break
			}
			pattern = append(pattern, line)
		}
		if len(pattern) == 0 {
			break
		}
		patterns = append(patterns, pattern)
	}
	out := 0
	for _, pattern := range patterns {
		// horizontal reflection sheck
		out += findMirror(pattern) * 100
		// vertial reflection check
		transposed := transpose(pattern)
		out += findMirror(transposed)
	}
	fmt.Println(out)
}
