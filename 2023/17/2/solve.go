package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Direction int64

const (
	Up    Direction = 1
	Down  Direction = -1
	Left  Direction = 2
	Right Direction = -2
	None  Direction = 0
)

type Block struct {
	row int
	col int
	dir Direction
}

func main() {
	city := make([][]int, 0)
	best := make(map[Block]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		cityLine := make([]int, 0)
		for _, block := range line {
			loss, _ := strconv.Atoi(string(block))
			cityLine = append(cityLine, loss)
		}
		city = append(city, cityLine)
	}
	dirs := []Block{
		{-1, 0, Up},
		{1, 0, Down},
		{0, -1, Left},
		{0, 1, Right},
	}
	// Using priority queue would be much more ideal
	// as I can stop early and guarantee the correct result,
	// but implementing it is a pain.
	bfs := make([]Block, 0)
	bfs = append(bfs, Block{0, 0, None})
	for len(bfs) > 0 {
		top := bfs[0]
		bfs = bfs[1:]
		for _, dir := range dirs {
			// must turn 90 degrees
			if top.dir == dir.dir || top.dir == -dir.dir {
				continue
			}
			accumulate := best[top]
			for i := 1; i <= 10; i++ {
				next := Block{
					top.row + dir.row*i,
					top.col + dir.col*i,
					dir.dir,
				}
				// bounds checking
				if next.row < 0 || next.row >= len(city) ||
					next.col < 0 || next.col >= len(city[0]) {
					continue
				}
				accumulate += city[next.row][next.col]
				if i < 4 {
					continue
				}
				// check if there is actually an improvement
				if accumulate < best[next] ||
					best[next] == 0 {
					best[next] = accumulate
					bfs = append(bfs, next)
				}
			}
		}
	}
	result := 0
	for _, dir := range dirs {
		out := best[Block{len(city) - 1, len(city[0]) - 1, dir.dir}]
		if (out > 0 && out < result) || result == 0 {
			result = out
		}
	}
	fmt.Println(result)
}
