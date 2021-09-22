package board

import (
	"github.com/anupsingh31/tictactoe/errorhandling"
)

var Size int64
var Boardgrid = [][]string{}

func BoardFrame(size int64) {
	Size = size

	for i := 0; i < int(size); i++ {
		row := []string{}
		for j := 0; j < int(size); j++ {
			row = append(row, "_ ")
		}
		Boardgrid = append(Boardgrid, row)
	}
}

func SetGrid(mark string, positionBoard int64) (int64, int64, error) {
	row := getRow(positionBoard)
	column := getColumn(positionBoard)
	if Boardgrid[row][column] == "_ " {
		Boardgrid[row][column] = mark
		return row, column, nil
	} else {
		err := errorhandling.SlotAlreadyTaken()
		return 0, 0, err
	}
}

func getRow(grid int64) int64 {
	if grid%Size == 0 {
		return (grid / Size) - 1
	}
	return grid / Size
}

func getColumn(grid int64) int64 {
	if grid%Size == 0 {
		return Size - 1
	}
	return (grid % Size) - 1
}
