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
	idSum := 0
games:
	for { // each game
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		gameParts := strings.Split(line, ": ")
		gameId, _ := strconv.Atoi(strings.Split(gameParts[0], " ")[1])
		tests := strings.Split(gameParts[1], "; ")
		for _, test := range tests {
			colors := strings.Split(test, ", ")
			for _, color := range colors {
				colorParts := strings.Split(color, " ")
				colorQuantity, _ := strconv.Atoi(colorParts[0])
				switch colorParts[1] {
				case "red":
					if colorQuantity > 12 {
						continue games
					}
				case "green":
					if colorQuantity > 13 {
						continue games
					}
				case "blue":
					if colorQuantity > 14 {
						continue games
					}
				}
			}
		}
		idSum += gameId
	}
	fmt.Println(idSum)
}
