package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
			visited[idx] = append(visited[idx], -2)
		}
	}
	var start Point
	// finding starting point
	for idx, line := range pipes {
		i := strings.Index(line, "S")
		if i > -1 {
			bfs = append(bfs, Point{i, idx})
			visited[idx][i] = 0
			start = Point{i, idx}
			break
		}
	}
	// start bfs
	dirs := []Point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} // up, down, left, right
	// can search the corresponding direction if the current pipe is the following
	validCurr := []string{"|LJS", "|7FS", "-J7S", "-LFS"}
	// can search the adjacent pipe if the next pipe is the following
	validAdj := []string{"|7F", "|LJ", "-LF", "-J7"} //
	var outPoint Point
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
					outPoint = Point{newX, newY}
					break search
				}
				continue
			}
			// visit the neighbor and set distance
			visited[newY][newX] = visited[top.y][top.x] + 1
			bfs = append(bfs, Point{newX, newY})
		}
	}
	backtrack := make([]Point, 0)
	copy := make([][]rune, len(pipes))
	for idx := range pipes {
		copy[idx] = make([]rune, 0, len(pipes[0]))
		for i := 0; i < len(pipes[0]); i++ {
			copy[idx] = append(copy[idx], '.')
		}
	}
	copy[outPoint.y][outPoint.x] = rune(pipes[outPoint.y][outPoint.x])
	backtrack = append(backtrack, outPoint)
	for len(backtrack) != 0 {
		top := backtrack[0]
		backtrack = backtrack[1:]
		for _, dir := range dirs {
			newX := top.x + dir.x
			newY := top.y + dir.y
			if newX < 0 ||
				newX >= len(pipes[0]) ||
				newY < 0 ||
				newY >= len(pipes) {
				continue
			}
			if visited[newY][newX] == visited[top.y][top.x]-1 {
				backtrack = append(backtrack, Point{newX, newY})
				copy[newY][newX] = rune(pipes[newY][newX])
			}
		}
	}
	// fix S
	u := start.y-1 > 0 && visited[start.y-1][start.x] == 1
	d := start.y+1 < len(pipes) && visited[start.y+1][start.x] == 1
	l := start.x-1 > 0 && visited[start.y][start.x-1] == 1
	r := start.x+1 < len(pipes[0]) && visited[start.y+1][start.x] == 1

	if u && d {
		copy[start.y][start.x] = '|'
	} else if l && r {
		copy[start.y][start.x] = '-'
	} else if u && l {
		copy[start.y][start.x] = 'J'
	} else if u && r {
		copy[start.y][start.x] = 'L'
	} else if d && l {
		copy[start.y][start.x] = '7'
	} else if d && r {
		copy[start.y][start.x] = 'F'
	}
	// checking for enclosed
	edge1 := regexp.MustCompile(`L-*J`)
	edge2 := regexp.MustCompile(`F-*7`)
	edge3 := regexp.MustCompile(`L-*7`)
	edge4 := regexp.MustCompile(`F-*J`)
	out := 0
	for y, line := range copy {
		for x, c := range line {
			if c != '.' {
				continue
			}
			// check all characters left of the current one
			left := string(line[:x])
			// remove edge cases
			left = edge1.ReplaceAllString(left, "")
			left = edge2.ReplaceAllString(left, "")
			left = edge3.ReplaceAllString(left, "|")
			left = edge4.ReplaceAllString(left, "|")
			count := make(map[rune]int)
			for _, tile := range left {
				count[tile]++
			}
			calc := count['|']
			if calc%2 != 0 {
				out++
				copy[y][x] = 'I'
			}
		}
	}
	for _, l := range copy {
		fmt.Println(string(l))
	}
	fmt.Println(out)
}
