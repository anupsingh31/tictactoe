package board

import (
	"testing"

	"github.com/anupsingh31/tictactoe/cell"
)

func TestNewBoard(t *testing.T) {
	b := NewBoard(3)
	for _, val := range b.cells {
		if val.Getmark() != cell.Nomark {
			t.Error("Actual mark is ", val.Getmark(), "but excepted is : ", cell.Nomark)
		}
	}
}
