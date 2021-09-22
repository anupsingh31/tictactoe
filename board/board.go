package board

import (
	"github.com/anupsingh31/tictactoe/errorhandling"
)

var Size int
var Boardgrid = [][]string{}

func BoardFrame(size int) {
	Size = size

	for i := 0; i < size; i++ {
		row := []string{}
		for j := 0; j < size; j++ {
			row = append(row, "_ ")
		}
		Boardgrid = append(Boardgrid, row)
	}
}

func SetGrid(mark string, positionBoard int) (int, int, error) {
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

func getRow(grid int) int {
	if grid%Size == 0 {
		return (grid / Size) - 1
	}
	return grid / Size
}

func getColumn(grid int) int {
	if grid%Size == 0 {
		return Size - 1
	}
	return (grid % Size) - 1
}
