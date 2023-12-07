package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getCardVal(x rune) int {
	val, err := strconv.Atoi(string(x))
	if err == nil {
		return val
	}
	cards := map[rune]int{
		'T': 10,
		'J': 1,
		'Q': 11,
		'K': 12,
		'A': 13,
	}
	return cards[x]
}

// logic:
// add number of Js to the character that has the max number of occurrences
// then calculate hand tier just like before
func typeOfHand(x string) int { // higher int, more valuable typeOfHand
	count := make(map[rune]int)
	for _, r := range x {
		count[r]++
	}
	max := 0
	mul := 1
	for k, v := range count {
		if k == 'J' {
			continue
		}
		if v > max {
			max = v
		}
		mul *= v
	}
	if max == 0 {
		return 25
	}
	mul = (mul / max) * (max + count['J'])
	max = max + count['J']
	value := max * mul
	if value == 18 {
		return 10
	}
	return value
}

func main() {
	input := make(map[string]int)
	hands := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		parts := strings.Split(line, " ")
		bid, _ := strconv.Atoi(parts[1])
		input[parts[0]] = bid
		hands = append(hands, parts[0])
	}
	sort.Slice(hands, func(a, b int) bool {
		sA := hands[a]
		sB := hands[b]
		eval := typeOfHand(sA) - typeOfHand(sB)
		if eval != 0 { // not the same type of hand
			return eval < 0
		}
		// second section, compare values
		for idx, c := range sA {
			d := rune(sB[idx])
			if c != d {
				return getCardVal(c) < getCardVal(d)
			}
		}
		return true
	})
	out := 0
	for idx, hand := range hands {
		out += (idx + 1) * input[hand]
		fmt.Println(hand, input[hand], out)
	}
	fmt.Println(out)
}
