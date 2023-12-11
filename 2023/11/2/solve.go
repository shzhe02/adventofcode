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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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

func countBetween(nums []int, a, b int) int {
	upper := max(a, b)
	lower := min(a, b)
	count := 0
	for _, v := range nums {
		if v < upper && v > lower {
			count++
		}
	}
	return count
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
	}
	// Store indexes of empty rows and cols
	emptyRows := make([]int, 0)
	emptyCols := make([]int, 0)

	// find empty rows
	for idx, row := range space {
		if allDots(row) {
			emptyRows = append(emptyRows, idx)
		}
	}

	// find empty cols
colScan:
	for col := 0; col < len(space[0]); col++ {
		for _, row := range space {
			if row[col] != '.' {
				continue colScan
			}
		}
		// by now we are certain the column is empty
		emptyCols = append(emptyCols, col)
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
			gA := galaxies[i]
			gB := galaxies[j]
			// figure out how many of the cols are between the col values of each galaxy
			blankCols := countBetween(emptyCols, gA.col, gB.col)
			colsBetween := abs(gA.col-gB.col) + blankCols*999999
			// figure out how many of the rows are between the row values of each galaxy
			blankRows := countBetween(emptyRows, gA.row, gB.row)
			rowsBetween := abs(gA.row-gB.row) + blankRows*999999
			out += colsBetween + rowsBetween
		}
	}

	fmt.Println(out)
}
