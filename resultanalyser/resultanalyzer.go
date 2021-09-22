package resultanalyser

import (
	"github.com/anupsingh31/tictactoe/board"
)

var count int64 = 0

func ResultAnalyser(mark string, row, column int64) string {
	count += 1
	if checkRows(mark, row) {
		return "win"
	}
	if checkColumns(mark, column) {
		return "win"
	}
	if checkReverseDiagonal(mark) {
		return "win"
	}
	if checkDiagonal(mark) {
		return "win"
	}
	if count == (board.Size * board.Size) {
		return "draw"
	}
	return "progress"
}

func checkRows(mark string, row int64) bool {
	for j := 0; j < int(board.Size); j++ {
		if !(board.Boardgrid[row][j] == mark) {
			return false
		}
	}
	return true

}

func checkColumns(mark string, column int64) bool {
	for j := 0; j < int(board.Size); j++ {
		if !(board.Boardgrid[j][column] == mark) {
			return false
		}
	}
	return true

}

func checkDiagonal(mark string) bool {
	k := 0
	for k < int(board.Size) {
		if !(board.Boardgrid[k][k] == mark) {
			return false
		}
		k++
	}
	return true

}

func checkReverseDiagonal(mark string) bool {
	k := 0
	j := board.Size - 1
	for k < int(board.Size) {
		if !(board.Boardgrid[k][j] == mark) {
			return false
		}
		k++
		j--
	}
	return true
}
