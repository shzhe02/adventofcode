package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		data = append(data, line)
	}
	dirs := map[string]image.Point{
		"U": {0, -1},
		"D": {0, 1},
		"L": {-1, 0},
		"R": {1, 0},
	}
	area := 0
	cursor := image.Point{0, 0}
	for _, line := range data {
		parts := strings.Split(line, " ")
		length, _ := strconv.Atoi(parts[1])
		next := cursor.Add(dirs[parts[0]].Mul(length))
		area += (cursor.X*next.Y - cursor.Y*next.X) + length
		cursor = next
	}
	fmt.Println(area/2 + 1)
}
