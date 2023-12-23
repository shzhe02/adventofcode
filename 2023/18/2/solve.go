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
		"3": {0, -1},
		"1": {0, 1},
		"2": {-1, 0},
		"0": {1, 0},
	}
	area := 0
	cursor := image.Point{0, 0}
	for _, line := range data {
		parts := strings.Split(line, " ")
		l, _ := strconv.ParseInt(parts[2][2:7], 16, strconv.IntSize)
		length := int(l)
		next := cursor.Add(dirs[parts[2][7:8]].Mul(length))
		area += (cursor.X*next.Y - cursor.Y*next.X) + length
		cursor = next
	}
	fmt.Println(area/2 + 1)
}
