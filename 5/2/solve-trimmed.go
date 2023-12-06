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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var sections [][]string
	for {
		var section []string
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) == 0 {
				break
			}
			section = append(section, line)
		}
		sections = append(sections, section)
		scanner.Scan()
		if len(scanner.Text()) == 0 {
			break
		}
	}
	var seeds [][]int // contains pairs in format (start, range)
	seedsSplit := strings.Split(sections[0][0], " ")
	for i := 1; i < len(seedsSplit); i += 2 {
		start, _ := strconv.Atoi(seedsSplit[i])
		q, _ := strconv.Atoi(seedsSplit[i+1])
		seeds = append(seeds, []int{start, q})
	}
	for i := 1; i < len(sections); i++ {
		var newSeeds [][]int
		section := sections[i]
		var dest []int
		var src []int
		var rang []int
		for _, conv := range section { // converting section into map
			nums := strings.Split(conv, " ") // each mapping
			d, _ := strconv.Atoi(nums[0])
			dest = append(dest, d)
			s, _ := strconv.Atoi(nums[1])
			src = append(src, s)
			r, _ := strconv.Atoi(nums[2])
			rang = append(rang, r)
		}
		// begin matching seeds
		// for _, seedRange := range seeds {
		for i := 0; i < len(seeds); i++ {
			seedRange := seeds[i]
			// seedRange is a pair
			// check overlap between the seed pair and each map section
			overlaps := false
			for idx, s := range src {
				r := rang[idx]
				d := dest[idx]
				overlapStart := max(seedRange[0], s)
				overlapEnd := min(seedRange[0]+seedRange[1], s+r)
				overlapLen := overlapEnd - overlapStart
				if overlapLen < 1 { // no overlap; try with next mapping
					continue
				}
				offset := overlapStart - s
				left := overlapStart - seedRange[0]
				right := seedRange[1] - overlapLen - left
				if left > 0 {
					seeds = append(seeds, []int{seedRange[0], left})
				}
				if right > 0 {
					seeds = append(seeds, []int{seedRange[0] + left + overlapLen, right})
				}
				newSeeds = append(newSeeds, []int{d + offset, overlapLen})
				overlaps = true
			}
			if !overlaps {
				newSeeds = append(newSeeds, seedRange)
			}
		}
		seeds = newSeeds
	}
	minimumLoc := int(^uint(0) >> 1) // max
	for _, seed := range seeds {
		if seed[0] < minimumLoc {
			minimumLoc = seed[0]
		}
	}
	fmt.Println(minimumLoc)
}
