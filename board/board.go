package board

import (
	"errors"

	"github.com/anupsingh31/tictactoe/cell"
)

type Board struct {
	cells  []cell.Cell
	size   int32
	marked int32
}

func New(size int32) *Board {
	c := make([]cell.Cell, size*size+1)
	for i := 0; i < len(c); i++ {
		c[i] = *cell.New()
	}
	return &Board{
		cells: c,
		size:  size,
	}
}

func (b *Board) SetMarkCell(cellId int32, mark string) error {
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

func (b *Board) GetSizeBoard() int32 {
	return b.size
}

func (b *Board) GetPositionMark(cellId int32) string {
	return b.cells[cellId-1].Getmark()
}

func (b *Board) String() string {
	var cellNo int32 = 0
	str := "    "
	for i := int32(0); i < b.size; i++ {
		for j := int32(0); j < b.size; j++ {
			str += "_____" + b.cells[cellNo].Getmark() + "_____"
			if j < b.size-1 {
				str += "|"
			}
			cellNo += 1
		}
		str += "\n    "

	}
	return str
}

func (b *Board) GetPositionFilled() int32 {
	return b.marked
}
