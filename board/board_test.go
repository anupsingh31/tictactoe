package board

import (
	"testing"

	"github.com/anupsingh31/tictactoe/cell"
)

var b = New(3)

func TestNewBoard(t *testing.T) {
	for _, val := range b.cells {
		if val.Getmark() != cell.Nomark {
			t.Error("Actual mark is ", val.Getmark(), "but excepted is : ", cell.Nomark)
		}
	}
}

func TestSetMarkCell(t *testing.T) {
	actual := b.SetMarkCell(3, "X")
	var expected error = nil
	if actual != expected {
		t.Error("Actual is ", actual, "but excepted is : ", expected)
	}
}

func TestGetSizeBoard(t *testing.T) {
	actual := b.GetSizeBoard()
	expected := 3
	if int(actual) != expected {
		t.Error("Actual is ", actual, "but excepted is : ", expected)
	}
}

func TestGetPositionMark(t *testing.T) {
	actual := b.GetPositionMark(3)
	expected := "X"
	if actual != expected {
		t.Error("Actual is ", actual, "but excepted is : ", expected)
	}
}
