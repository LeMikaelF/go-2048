package engine

import (
	"math/rand"
	"testing"

	"github.com/LeMikaelF/2048/src/grid"
)

func Test(t *testing.T) {
	type test struct {
		name         string
		grid         grid.Grid
		randomSource rand.Source
		move         Direction
		assertions   func(t *testing.T, actualGrid grid.Grid, err error)
	}
	tests := []test{
		{
			name: "given a grid with one tile and no slides, it generates a number on an empty square (random seed 1)",
			grid: grid.Grid{
				[]int{0, 0, 0, 0},
				[]int{2, 0, 0, 0},
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
			},
			randomSource: rand.NewSource(1),
			move:         Left,
			assertions: func(t *testing.T, actualGrid grid.Grid, err error) {
				assertNil(t, err)
				assertEquals(t, actualGrid, grid.Grid{
					[]int{0, 0, 0, 0},
					[]int{2, 0, 0, 0},
					[]int{0, 0, 0, 0},
					[]int{0, 0, 2, 0},
				})
			},
		},
		{
			name: "given a grid with one tile and no slides, it generates a number on an empty square (random seed 2)",
			grid: grid.Grid{
				[]int{0, 0, 0, 0},
				[]int{2, 0, 0, 0},
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
			},
			randomSource: rand.NewSource(2),
			move:         Left,
			assertions: func(t *testing.T, actualGrid grid.Grid, err error) {
				assertNil(t, err)
				assertEquals(t, actualGrid, grid.Grid{
					[]int{0, 0, 0, 0},
					[]int{2, 2, 0, 0},
					[]int{0, 0, 0, 0},
					[]int{0, 0, 0, 0},
				})
			},
		},
		{
			name: "given a grid with one tile and one left slide, it slides the tile",
			grid: grid.Grid{
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 2},
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
			},
			randomSource: rand.NewSource(1),
			move:         Left,
			assertions: func(t *testing.T, actualGrid grid.Grid, err error) {
				assertNil(t, err)
				assertEquals(t, actualGrid, grid.Grid{
					[]int{0, 0, 0, 0},
					[]int{2, 0, 0, 0},
					[]int{0, 0, 0, 0},
					[]int{0, 0, 2, 0},
				})
			},
		},
		{
			name: "given a grid with one tile and one right slide, it slides the tile",
			grid: grid.Grid{
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
				[]int{2, 0, 0, 0},
				[]int{0, 0, 0, 0},
			},
			randomSource: rand.NewSource(1),
			move:         Right,
			assertions: func(t *testing.T, actualGrid grid.Grid, err error) {
				assertNil(t, err)
				assertEquals(t, actualGrid, grid.Grid{
					[]int{0, 0, 0, 0},
					[]int{0, 0, 0, 0},
					[]int{0, 0, 0, 2},
					[]int{0, 0, 2, 0},
				})
			},
		},
		{
			name: "given a grid with one tile and one down slide, it slides the tile",
			grid: grid.Grid{
				[]int{0, 0, 4, 0},
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
			},
			randomSource: rand.NewSource(1),
			move:         Down,
			assertions: func(t *testing.T, actualGrid grid.Grid, err error) {
				assertNil(t, err)
				assertEquals(t, actualGrid, grid.Grid{
					[]int{0, 0, 0, 0},
					[]int{0, 0, 0, 0},
					[]int{0, 0, 0, 0},
					[]int{0, 2, 4, 0},
				})
			},
		},
		{
			name: "given a grid with one tile and one up slide, it slides the tile",
			grid: grid.Grid{
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 4},
			},
			randomSource: rand.NewSource(1),
			move:         Up,
			assertions: func(t *testing.T, actualGrid grid.Grid, err error) {
				assertNil(t, err)
				assertEquals(t, actualGrid, grid.Grid{
					[]int{0, 0, 0, 4},
					[]int{0, 0, 0, 0},
					[]int{0, 0, 0, 0},
					[]int{0, 0, 2, 0},
				})
			},
		},

		{
			name: "given a grid with only one blank square, it fills it up",
			grid: grid.Grid{
				[]int{2, 2, 2, 2},
				[]int{2, 2, 2, 2},
				[]int{2, 2, 0, 2},
				[]int{2, 2, 2, 2},
			},
			randomSource: rand.NewSource(1),
			move:         Up,
			assertions: func(t *testing.T, actualGrid grid.Grid, err error) {
				assertNil(t, err)
				assertEquals(t, actualGrid, grid.Grid{
					[]int{2, 2, 2, 2},
					[]int{2, 2, 2, 2},
					[]int{2, 2, 2, 2},
					[]int{2, 2, 2, 2},
				})
			},
		},
		{
			name: "given a full grid, it returns an error",
			grid: grid.Grid{
				[]int{2, 2, 2, 2},
				[]int{2, 2, 2, 2},
				[]int{2, 2, 2, 2},
				[]int{2, 2, 2, 2},
			},
			randomSource: rand.NewSource(1),
			move:         Up,
			assertions: func(t *testing.T, actualGrid grid.Grid, err error) {
				type lostError interface {
					Lost()
				}
				if err, ok := err.(lostError); !ok {
					t.Errorf("wanted error with Lost() method, was %v (type %t)", err, err)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			engine := NewFromLiteral(tt.grid, withRandomSource(tt.randomSource))
			err := engine.Next(tt.move)

			tt.assertions(t, engine.Grid, err)
		})
	}
}

func assertNil(t *testing.T, err error) {
	if err != nil {
		t.Errorf("expected nil, was %v", err)
	}
}

type Equalable[T any] interface {
	Equals(T) bool
}

func assertEquals[LEFT Equalable[RIGHT], RIGHT any](t *testing.T, left LEFT, right RIGHT) {
	equals := left.Equals(right)
	if !equals {
		t.Errorf("wanted equality, but left was\n%v\nand right was\n%v\n", left, right)
	}
}
