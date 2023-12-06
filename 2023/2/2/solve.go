package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	powSum := 0
	for { // each game
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		gameParts := strings.Split(line, ": ")
		tests := strings.Split(gameParts[1], "; ")
		rMin := 0
		gMin := 0
		bMin := 0
		for _, test := range tests {
			colors := strings.Split(test, ", ")
			for _, color := range colors {
				colorParts := strings.Split(color, " ")
				colorQuantity, _ := strconv.Atoi(colorParts[0])
				switch colorParts[1] {
				case "red":
					rMin = max(rMin, colorQuantity)
				case "green":
					gMin = max(gMin, colorQuantity)
				case "blue":
					bMin = max(bMin, colorQuantity)
				}
			}
		}
		powSum += rMin * gMin * bMin
	}
	fmt.Println(powSum)
}
