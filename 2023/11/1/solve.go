package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	col, row int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func allDots(a string) bool {
	out := true
	for _, s := range a {
		if s != '.' {
			out = false
		}
	}
	return out
}

func main() {
	space := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		space = append(space, line)
		// expand rows
		if allDots(line) {
			space = append(space, line)
		}
	}
	// expand columns
colScan:
	for col := 0; col < len(space[0]); col++ {
		for _, row := range space {
			if row[col] != '.' {
				continue colScan
			}
		}
		// by now we are certain the column is empty, so we add the extra column
		for row := 0; row < len(space); row++ {
			space[row] = space[row][:col] + "." + space[row][col:]
		}
		// skip the next column, because we just added it
		col++
	}

	// finding galaxies
	galaxies := make([]Point, 0)
	for col := 0; col < len(space[0]); col++ {
		for row := 0; row < len(space); row++ {
			if space[row][col] == '#' {
				galaxies = append(galaxies, Point{col, row})
			}
		}
	}

	out := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			out += abs(galaxies[i].col-galaxies[j].col) + abs(galaxies[i].row-galaxies[j].row)
		}
	}

	fmt.Println(out)
}
