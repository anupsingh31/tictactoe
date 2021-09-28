package resultanalyser

import (
	"testing"

	"github.com/anupsingh31/tictactoe/board"
)

var ra = New()
var b = board.New(3)

func TestGetBoardStatus(t *testing.T) {
	b.SetMarkCell(1, "X")
	actual := ra.GetBoardStatus(b, "X", 1)
	expected := GameProgress
	if actual != expected {
		t.Error("Actual is ", actual, "but excepted is : ", expected)
	}

	b.SetMarkCell(4, "O")

	actual = ra.GetBoardStatus(b, "O", 4)
	expected = GameProgress
	if actual != expected {
		t.Error("Actual is ", actual, "but excepted is : ", expected)
	}

	b.SetMarkCell(2, "X")

	actual = ra.GetBoardStatus(b, "X", 2)
	expected = GameProgress
	if actual != expected {
		t.Error("Actual is ", actual, "but excepted is : ", expected)
	}

	b.SetMarkCell(5, "O")

	actual = ra.GetBoardStatus(b, "O", 5)
	expected = GameProgress
	if actual != expected {
		t.Error("Actual is ", actual, "but excepted is : ", expected)
	}

	b.SetMarkCell(3, "X")

	actual = ra.GetBoardStatus(b, "X", 3)
	expected = GameWin
	if actual != expected {
		t.Error("Actual is ", actual, "but excepted is : ", expected)
	}

}
