package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	pipes := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		pipes = append(pipes, line)
	}
	bfs := make([]Point, 0)
	visited := make([][]int, len(pipes))
	for idx := range visited {
		for i := 0; i < len(pipes[0]); i++ {
			visited[idx] = append(visited[idx], -1)
		}
	}
	// finding starting point
	for idx, line := range pipes {
		i := strings.Index(line, "S")
		if i > -1 {
			bfs = append(bfs, Point{i, idx})
			visited[idx][i] = 0
			break
		}
	}
	// start bfs
	dirs := []Point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} // up, down, left, right
	// can search the corresponding direction if the current pipe is the following
	validCurr := []string{"|LJS", "|7FS", "-J7S", "-LFS"}
	// can search the adjacent pipe if the next pipe is the following
	validAdj := []string{"|7F", "|LJ", "-LF", "-J7"} //
	out := 0
search:
	for len(bfs) != 0 {
		top := bfs[0] // top
		bfs = bfs[1:] // pop
		for idx, dir := range dirs {
			// check if the current pipe can go in this direction
			currPipe := pipes[top.y][top.x]
			if strings.ContainsRune(validCurr[idx], rune(currPipe)) == false {
				continue
			}
			newX := top.x + dir.x
			newY := top.y + dir.y
			// continue if new coord is invalid or has been visited already
			if newX < 0 ||
				newX >= len(pipes[0]) ||
				newY < 0 ||
				newY >= len(pipes) {
				continue
			}
			adjPipe := pipes[newY][newX]
			// check if adjacent pipe is valid
			if strings.ContainsRune(validAdj[idx], rune(adjPipe)) == false {
				continue
			}
			// check if pipe has been visited
			if visited[newY][newX] >= 0 {
				// loop found? (loop midpoint found)
				if visited[newY][newX] == visited[top.y][top.x]+1 {
					out = visited[top.y][top.x] + 1
					break search
				}
				continue
			}
			// visit the neighbor and set distance
			visited[newY][newX] = visited[top.y][top.x] + 1
			bfs = append(bfs, Point{newX, newY})
		}
	}
	fmt.Println(out)
}
