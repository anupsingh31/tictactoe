package game

import (
	"testing"

	"github.com/anupsingh31/tictactoe/board"
	"github.com/anupsingh31/tictactoe/player"
)

var b = board.New(3)
var p1 = player.New("anupam", "X")
var p2 = player.New("abhi", "O")
var p = [2]*player.Player{p1, p2}
var g = New(3, p)

func TestNewGame(t *testing.T) {
	expectedBoardSize := 3
	actualBoardSize := g.b.GetSizeBoard()
	if int(actualBoardSize) != expectedBoardSize {
		t.Error("Actual is ", actualBoardSize, "but excepted is : ", expectedBoardSize)
	}
	var expectedTurn uint8 = 0
	actualTurn := g.turn
	if actualTurn != expectedTurn {
		t.Error("Actual is ", actualTurn, "but excepted is : ", expectedTurn)
	}
	var expectedLimitMove uint8 = 0
	actualLimitMove := g.limitmoved
	if actualLimitMove != expectedLimitMove {
		t.Error("Actual is ", actualTurn, "but excepted is : ", expectedLimitMove)
	}
}
