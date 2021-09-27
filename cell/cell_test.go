package cell

import "testing"

func TestNewCell(t *testing.T) {
	expectedMark := Nomark
	actual := NewCell()

	if actual.mark != expectedMark {
		t.Error("Actual is ", actual.mark, "but excepted is : ", expectedMark)
	}
}
