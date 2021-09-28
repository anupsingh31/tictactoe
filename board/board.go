package board

import (
	"errors"
	"strconv"

	"github.com/anupsingh31/tictactoe/cell"
)

type Board struct {
	cells  []cell.Cell
	size   uint8
	marked uint8
}

func New(size uint8) *Board {
	c := make([]cell.Cell, size*size+1)
	for i := 0; i < len(c); i++ {
		c[i] = *cell.New()
	}
	return &Board{
		cells: c,
		size:  size,
	}
}

func (b *Board) SetMarkCell(cellId uint8, mark string) error {
	if cellId > b.size*b.size || cellId < 1 {
		return errors.New("index out of bounds")
	}
	err := b.cells[cellId-1].SetMark(mark)
	if err != nil {
		return err
	}
	b.marked++
	return nil
}

func (b *Board) GetSizeBoard() uint8 {
	return b.size
}

func (b *Board) GetPositionMark(cellId uint8) string {
	return b.cells[cellId-1].Getmark()
}

func (b *Board) String() string {
	var cellNo uint8 = 0
	str := "    "
	for i := uint8(0); i < b.size; i++ {
		for j := uint8(0); j < b.size; j++ {
			str += "_____"
			if b.cells[cellNo].Getmark() == cell.Nomark {
				str += strconv.Itoa(int(cellNo) + 1)
			} else {
				str += b.cells[cellNo].Getmark()
			}
			str += "_____"
			if j < b.size-1 {
				str += "|"
			}
			cellNo += 1
		}
		str += "\n    "

	}
	return str
}

func (b *Board) GetPositionFilled() uint8 {
	return b.marked
}
