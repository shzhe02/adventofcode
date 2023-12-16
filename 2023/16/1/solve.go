package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Direction int64

const (
	Up    Direction = 0
	Down  Direction = 1
	Left  Direction = 2
	Right Direction = 3
)

type Beam struct {
	row int
	col int
	dir Direction
}

func main() {
	nextDir := make(map[rune]map[Direction]Direction)
	nextDir['/'] = map[Direction]Direction{
		Up:    Right,
		Down:  Left,
		Left:  Down,
		Right: Up,
	}
	nextDir['\\'] = map[Direction]Direction{
		Up:    Left,
		Down:  Right,
		Left:  Up,
		Right: Down,
	}
	grid := make([]string, 0)
	energized := make([][]bool, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		grid = append(grid, line)
		row := make([]bool, len(line))
		energized = append(energized, row)
	}
	beams := make([]Beam, 0) // keep track of all active beams.
	// each beam represents its direction upon entering the tile located at row, col.
	beams = append(beams, Beam{0, 0, Right})
	for i := 0; i < len(beams); i++ {
		beam := beams[i]
		currRow := beam.row
		currCol := beam.col
		currDir := beam.dir
		found := false
		for _, b := range beams[:i] {
			if b.row == currRow && b.col == currCol && b.dir == currDir {
				found = true
			}
		}
		if found {
			continue
		}
		// keep the beam moving
		for {
			// check if beam is in bounds
			if currRow < 0 || currRow >= len(grid) || currCol < 0 || currCol >= len(grid[0]) {
				break
			}
			energized[currRow][currCol] = true
			currTile := rune(grid[currRow][currCol])
			if strings.Contains("/\\", string(currTile)) {
				// Beam does not split
				currDir = nextDir[currTile][currDir]
			} else if strings.Contains("|-", string(currTile)) {
				// Beam splits
				if currTile == '|' && (currDir == Left || currDir == Right) {
					beams = append(beams, Beam{currRow - 1, currCol, Up})
					beams = append(beams, Beam{currRow + 1, currCol, Down})
					break
				} else if currTile == '-' && (currDir == Up || currDir == Down) {
					beams = append(beams, Beam{currRow, currCol - 1, Left})
					beams = append(beams, Beam{currRow, currCol + 1, Right})
					break
				}
			}
			// advance the beam
			switch currDir {
			case Up:
				currRow -= 1
			case Down:
				currRow += 1
			case Left:
				currCol -= 1
			case Right:
				currCol += 1
			}
		}
	}
	out := 0
	for _, l := range energized {
		for _, e := range l {
			if e {
				out++
			}
		}
	}
	fmt.Println(out)
}
