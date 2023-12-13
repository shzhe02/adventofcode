package main

import (
	"bufio"
	"fmt"
	"os"
)

func diff(a, b string) int {
	diffs := 0
	for idx := range a {
		if a[idx] != b[idx] {
			diffs++
		}
	}
	return diffs
}

func findMirror(pattern []string) int { // lines above the mirror if found, 0 if not found.
	// mirror in the top half (first line is mirrored)
topSearch:
	for i := 1; i < len(pattern); i += 2 {
		totalDiffs := 0
		totalDiffs += diff(pattern[i], pattern[0])
		if pattern[i] == pattern[0] || totalDiffs == 1 {
			// mirror ends potentially found, check if correctly mirrored
			t := 1
			b := i - 1
			for t < b {
				totalDiffs += diff(pattern[t], pattern[b])
				if totalDiffs > 1 {
					continue topSearch
				}
				t++
				b--
			}
			// scan passed, mirror found.
			if totalDiffs == 1 {
				return (i / 2) + 1
			}
		}
	}
	// mirror in the bottom half (last line is mirrored)
bottomSearch:
	for i := len(pattern) - 2; i > 0; i -= 2 {
		totalDiffs := 0
		totalDiffs += diff(pattern[i], pattern[len(pattern)-1])
		if pattern[i] == pattern[len(pattern)-1] || totalDiffs == 1 {
			t := i + 1
			b := len(pattern) - 2
			for t < b {
				totalDiffs += diff(pattern[t], pattern[b])
				if totalDiffs > 1 {
					continue bottomSearch
				}
				t++
				b--
			}
			if totalDiffs == 1 {
				return len(pattern) - (((len(pattern) - 1 - i) / 2) + 1)
			}
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
