package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkZero(arr []int) bool {
	out := true
	for _, v := range arr {
		out = out && v == 0
	}
	return out
}

func main() {
	histories := make([][]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		parts := strings.Split(line, " ")
		history := make([]int, 0, len(parts))
		for _, v := range parts {
			num, _ := strconv.Atoi(v)
			history = append(history, num)
		}
		histories = append(histories, history)
	}
	out := 0
	for _, history := range histories {
		cascade := [][]int{history}
		nextSize := len(history)
		zeroed := checkZero(history)
		row := 0
		for !zeroed {
			next := make([]int, 0, nextSize-1)
			for idx, v := range cascade[row] {
				if idx == 0 {
					continue
				}
				nextVal := v - cascade[row][idx-1]
				next = append(next, nextVal)
			}
			zeroed = checkZero(next)
			cascade = append(cascade, next)
			nextSize--
			row++
		}
		// last row should be 0 by now
		val := 0
		for i := len(cascade) - 1; i >= 0; i-- {
			val = cascade[i][0] - val
		}
		out += val
	}
	fmt.Println(out)
}
