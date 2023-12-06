package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	// fmt.Println(sections)
	var seeds []int
	for idx, v := range strings.Split(sections[0][0], " ") {
		if idx == 0 {
			continue
		}
		seed, _ := strconv.Atoi(v)
		seeds = append(seeds, seed)
	}
	// fmt.Println(seeds)
	for i := 1; i < len(sections); i++ {
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
		for seedIdx, seed := range seeds {
			for idx, s := range src {
				if seed >= s && seed < s+rang[idx] {
					seeds[seedIdx] = dest[idx] + (seed - s)
					break
				}
			}
		}
		// fmt.Println(seeds)
	}
	minimumLoc := int(^uint(0) >> 1) // max
	for _, seed := range seeds {
		if seed < minimumLoc {
			minimumLoc = seed
		}
	}
	fmt.Println(minimumLoc)
}
