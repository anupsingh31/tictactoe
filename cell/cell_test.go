package cell

import "testing"

func TestNewCell(t *testing.T) {
	expectedMark := Nomark
	actual := New()

	if actual.mark != expectedMark {
		t.Error("Actual is ", actual.mark, "but excepted is : ", expectedMark)
	}
}
