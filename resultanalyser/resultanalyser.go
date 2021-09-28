package resultanalyser

import (
	"github.com/anupsingh31/tictactoe/board"
)

type GameStatus string

const (
	GameProgress GameStatus = "Progress"
	GameWin      GameStatus = "Win"
	GameDraw     GameStatus = "Draw"
)

type ResultAnalyzer struct{}

func New() *ResultAnalyzer {
	return &ResultAnalyzer{}
}

func (ra *ResultAnalyzer) GetBoardStatus(b *board.Board, mark string, cellId uint8) GameStatus {
	if checkRows(b, mark, cellId) {
		return GameWin
	}
	if checkColumns(b, mark, cellId) {
		return GameWin
	}
	if checkReverseDiagonal(b, mark) {
		return GameWin
	}
	if checkDiagonal(b, mark) {
		return GameWin
	}
	if b.GetPositionFilled() == b.GetSizeBoard()*b.GetSizeBoard() {
		return GameDraw
	}
	return GameProgress
}

func checkRows(b *board.Board, mark string, cellId uint8) bool {
	var row uint8
	size := b.GetSizeBoard()
	if cellId%size == 0 {
		row = ((cellId/size)-1)*size + 1
	} else {
		row = (cellId/size)*size + 1
	}
	for j := 0; j < int(size); j++ {
		if !(b.GetPositionMark(row) == mark) {
			return false
		}
		row++
	}
	return true

}

func checkColumns(b *board.Board, mark string, cellId uint8) bool {
	var column uint8
	size := b.GetSizeBoard()
	if cellId%size == 0 {
		column = size
	} else {
		column = (cellId % size)
	}
	for j := 0; j < int(size); j++ {
		if !(b.GetPositionMark(column) == mark) {
			return false
		}
		column += size
	}
	return true

}

func checkDiagonal(b *board.Board, mark string) bool {
	var k, itr uint8 = 1, 1
	for itr <= b.GetSizeBoard() {
		if !(b.GetPositionMark(k) == mark) {
			return false
		}
		k += b.GetSizeBoard() + 1
		itr++
	}
	return true

}

func checkReverseDiagonal(b *board.Board, mark string) bool {
	var itr, j uint8 = 1, b.GetSizeBoard()
	for itr <= b.GetSizeBoard() {
		if !(b.GetPositionMark(j) == mark) {
			return false
		}
		j += (b.GetSizeBoard() - 1)
		itr++

	}
	return true
}
