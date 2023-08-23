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
		[4]int{0, 0, 0, 0},
		[4]int{0, 0, 0, 0},
		[4]int{0, 0, 0, 0},
		[4]int{0, 0, 0, 0},
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

func (e *Engine) Next(_ Direction) error {
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
		for iCol, num := range row {
			if num == 0 {
				allBlanks = append(allBlanks, Coord{iRow, iCol})
			}
		}
	}

	if len(allBlanks) == 0 {
		return Coord{}, errors.New("no blank square found")
	}

	return allBlanks[e.random.Intn(len(allBlanks)-1)], nil
}
