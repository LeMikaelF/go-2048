package engine

import (
	"errors"
	"math/rand"

	"github.com/LeMikaelF/2048/src/grid"
)

type Engine struct {
	Grid   grid.Grid
	random *rand.Rand
}

func New(options ...EngineOption) *Engine {
	return NewFromLiteral(grid.Grid{
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
	}, options...)
}

func NewFromLiteral(theGrid grid.Grid, options ...EngineOption) *Engine {
	e := &Engine{Grid: theGrid}

	e.random = rand.New(rand.NewSource(int64(rand.Uint64())))

	for _, option := range options {
		option(e)
	}

	return e
}

type EngineOption func(*Engine)

func withRandomSource(source rand.Source) EngineOption {
	return func(engine *Engine) {
		engine.random = rand.New(source)
	}
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
	//TODO slide everything
	e.slideAll(direction)

	coord, err := e.findRandomBlank()
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

func (e *Engine) findRandomBlank() (Coord, error) {
	allBlanks := make([]Coord, 0)

	for iRow, row := range e.Grid {
		for iCol, val := range row {
			if val == 0 {
				allBlanks = append(allBlanks, Coord{iRow, iCol})
			}
		}
	}

	if len(allBlanks) == 0 {
		return Coord{}, errors.New("no blank square found")
	}

	return allBlanks[e.random.Intn(len(allBlanks)-1)], nil
}

func (e *Engine) slideAll(direction Direction) {
	//directionToSlideCoords := map[Direction]Coord{
	//Up:    Coord{-1, 0},
	//Down:  Coord{1, 0},
	//Left:  Coord{0, -1},
	//Right: Coord{0, 1},
	//}
	//
	//slideCoords, ok := directionToSlideCoords[direction]
	//if !ok {
	//	panic(fmt.Sprintf("unknown direction: %v", direction))
	//}

	for _, row := range e.Grid {
		for iCol := 0; iCol < len(row); {
			canSlide := iCol != 0 && row[iCol-1] == 0 && row[iCol] != 0
			if canSlide {
				row[iCol-1] = row[iCol]
				row[iCol] = 0
				iCol--
			} else {
				iCol++
			}
		}
	}
}
