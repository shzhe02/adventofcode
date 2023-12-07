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
		'J': 11,
		'Q': 12,
		'K': 13,
		'A': 14,
	}
	return cards[x]
}

// logic:
// We obtain the highest occurrences of a single character
// as well as the product of all the character occurrences multiplied together
// then multiply them together to get the score.
//
// Hand Type:
// | Hand Type | Max Count | Count Product | Score |
// | --------- | --------- | ------------- | ----- |
// | 5oAK      | 5         | 5             | 25    |
// | 4oAK      | 4         | 4             | 16    |
// | FH        | 3         | 6 (2 x 3)     | 18    |
// | 3oAK      | 3         | 3             | 9     |
// | 2P        | 2         | 4 (2 x 2)     | 8     |
// | 1P        | 2         | 2             | 4     |
// | HC        | 1         | 1             | 1     |
// Now we just have the exception for full house by artificially decreasing
// its score to 10, and the scores will be in the right order.
func typeOfHand(x string) int { // higher int, more valuable typeOfHand
	count := make(map[rune]int)
	for _, r := range x {
		count[r]++
	}
	max := 0
	mul := 1
	for _, v := range count {
		if v > max {
			max = v
		}
		mul *= v
	}
	if max*mul == 18 {
		return 10
	}
	return max * mul
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
