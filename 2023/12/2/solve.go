package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// if in consists of only '?' and '.', it is clear
func checkClear(in string) bool {
	out := true
	for _, v := range in {
		if v == '#' {
			out = false
		}
	}
	return out
}

// glossary:
// - block: continuous sequence of ? or #
func calculate(remaining string, groups []int, mem map[string]int) int {
	// Base cases + optimizations
	if len(groups) == 0 {
		// all groups have been accounted for, therefore this is a valid solution.
		if checkClear(remaining) {
			return 1
		} else {
			return 0
		}
	} else if len(remaining) < sum(groups)+len(groups)-1 {
		// otherwise, unable to add the group of springs
		return 0
	}
	// preproccessing
	for len(remaining) >= groups[0] {
		// remove all dots in the beginning of remaining
		// .....??? will have the same result as ??? no matter the groups
		i := 0
		for i < len(remaining) {
			if rune(remaining[i]) != '.' {
				break
			}
			i++
		}
		remaining = remaining[i:]
		// figure out block size
		i = 0
		for i < len(remaining) {
			if rune(remaining[i]) == '.' {
				break
			}
			i++
		}
		if i == groups[0] || (i > groups[0] && rune(remaining[groups[0]]) == '?') {
			break
		} else if len(remaining) > 0 && rune(remaining[0]) == '?' {
			remaining = remaining[1:]
		} else {
			return 0
		}
	}
	if len(remaining) < sum(groups)+len(groups)-1 {
		return 0
	}
	count, exists := mem[remaining+string(len(groups))]
	if exists {
		return count
	}

	out := 0
	// two options
	// A - remaining string starts with a group of springs
	if len(remaining) == groups[0] {
		// A1 - group uses up the rest of remaining
		out += calculate(remaining[groups[0]:], groups[1:], mem)
	} else {
		// A2 - group does not use up the rest of remaining
		out += calculate(remaining[groups[0]+1:], groups[1:], mem)
	}
	// B - remaining string starts with a ? that becomes a .
	if remaining[0] == '?' {
		out += calculate(remaining[1:], groups, mem)
	}
	mem[remaining+string(len(groups))] = out
	return out
}

func main() {
	records := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		records = append(records, line)
	}
	out := 0
	progress := 1
	for _, record := range records {
		parts := strings.Split(record, " ")
		field := strings.Repeat(parts[0]+"?", 5)
		field = field[:len(field)-1]
		groupsString := strings.Repeat(parts[1]+",", 5)
		groupsString = groupsString[:len(groupsString)-1]
		groups := make([]int, 0)
		for _, v := range strings.Split(groupsString, ",") {
			num, _ := strconv.Atoi(v)
			groups = append(groups, num)
		}
		mem := make(map[string]int)
		out += calculate(field, groups, mem)
		progress++
	}
	fmt.Println(out)
}
