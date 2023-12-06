package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f := func(c rune) bool {
		return c == ' '
	}
	scanner := bufio.NewScanner(os.Stdin)
	var data [][]int // first row = duration, 2nd row = distance threshold
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		var nums []int
		for idx, val := range strings.FieldsFunc(line, f) {
			if idx == 0 {
				continue
			}
			num, _ := strconv.Atoi(val)
			nums = append(nums, num)
		}
		data = append(data, nums)
	}
	var waysToWin []int
	for idx, dur := range data[0] {
		dist := data[1][idx]
		// first, check if distance is maximal
		if dist >= (dur>>1)*(dur-dur>>1) {
			continue
		}
		l := 0
		for {
			if l*(dur-l) > dist {
				break
			}
			l++
		}
		waysToWin = append(waysToWin, dur-2*l+1)
	}
	// binary search for both ends
	out := 1
	if len(waysToWin) == 0 {
		fmt.Println(0)
	} else {
		for _, v := range waysToWin {
			out *= v
		}
		fmt.Println(out)
	}
}
