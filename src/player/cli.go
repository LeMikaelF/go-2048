package player

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/LeMikaelF/2048/src/engine"
	"github.com/LeMikaelF/2048/src/grid"
)

type Cli struct {
	engine *engine.Engine
}

func New() *Cli {
	return &Cli{
		engine: engine.New(),
	}
}

type lostError interface {
	Lost()
}

func (c *Cli) Run() {
	for {
		fmt.Println("Current grid.")
		fmt.Println(prettyPrint(c.engine.Grid))

		direction, err := readDirection()
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}

		fmt.Println("")
		err = c.engine.Next(direction)
		if _, ok := err.(lostError); ok {
			fmt.Println("******You lost the game!********")
			return
		}
	}
}

func readDirection() (engine.Direction, error) {
	bytes := make([]byte, 4)
	for {
		fmt.Println("Press any arrow, then Enter.")
		_, err := os.Stdin.Read(bytes)
		if err != nil {
			return "", errors.New("error reading from stdin")
		}

		switch fmt.Sprintf("%v", bytes) {
		case "[27 91 67 10]":
			return engine.Right, nil
		case "[27 91 68 10]":
			return engine.Left, nil
		case "[27 91 66 10]":
			return engine.Down, nil
		case "[27 91 65 10]":
			return engine.Up, nil
		default:
			fmt.Printf("unknown input %v, try again.\n", bytes)
		}
	}
}

func prettyPrint(grid grid.Grid) string {
	var sb strings.Builder
	for iRow, row := range grid {
		for _, val := range row {
			sb.WriteString(strconv.Itoa(val))
			sb.WriteString(" ")
		}
		if iRow != len(grid)-1 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}
