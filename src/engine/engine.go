package engine

import (
	"errors"

	"github.com/LeMikaelF/2048/src/grid"
)

type Engine struct {
	Grid grid.Grid
}

func New() *Engine {
	return NewFromLiteral(grid.Grid{
		[4]int{0, 0, 0, 0},
		[4]int{0, 0, 0, 0},
		[4]int{0, 0, 0, 0},
		[4]int{0, 0, 0, 0},
	})
}

func NewFromLiteral(theGrid grid.Grid) *Engine {
	return &Engine{Grid: theGrid}
}

type Direction string

const (
	Up    Direction = "up"
	Down  Direction = "down"
	Left  Direction = "left"
	Right Direction = "right"
)

type lostError struct{}

func (l lostError) Error() string {
	return "player lost the game"
}

func (l lostError) Lost() {
	// marker method
}

func (e *Engine) Next(direction Direction) error {
	coord, err := e.findBlank()
	if err != nil {
		return lostError{}
	}

	e.Grid[coord.row][coord.col] = 2

	return nil
}

type Coord struct {
	row int
	col int
}

func (e *Engine) findBlank() (Coord, error) {
	for iRow, row := range e.Grid {
		for iCol, num := range row {
			if num == 0 {
				return Coord{iRow, iCol}, nil
			}
		}
	}

	return Coord{}, errors.New("no blank square found")
}
