package cell

import (
	"testing"
)

func TestNewCell(t *testing.T) {
	expectedMark := Nomark
	actual := New()

	if actual.mark != expectedMark {
		t.Error("Actual is ", actual.mark, "but excepted is : ", expectedMark)
	}
}

func TestSetMark(t *testing.T) {
	var expectedVal error = nil
	actualVal := New().SetMark("X")
	if actualVal != expectedVal {
		t.Error("Actual is ", actualVal, "but excepted is : ", expectedVal)
	}
}

func TestGetMark(t *testing.T) {
	expected := New()
	expected.SetMark("X")
	expectedVal := "X"
	actualVal := expected.Getmark()
	if actualVal != expectedVal {
		t.Error("Actual is ", actualVal, "but excepted is : ", expectedVal)
	}
}
