package cell

import (
	"errors"
)

const (
	Xmark  = "X"
	Omark  = "O"
	Nomark = "-"
)

type Cell struct {
	mark string
}

func NewCell() *Cell {
	return &Cell{mark: Nomark}
}

func (c *Cell) SetMark(mark string) error {
	if c.mark == Nomark {
		c.mark = mark
		return nil
	} else {
		return errors.New("Cell is already marked")
	}
}

func (c *Cell) Getmark() string {
	return c.mark
}
