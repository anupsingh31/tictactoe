package player

import "testing"

var actual = New("anupam", "X")

func TestNewPlayer(t *testing.T) {
	expectedPlayerName := "anupam"
	expectedPlayerMark := "X"
	if actual.name != expectedPlayerName {
		t.Error("Actual is ", actual.name, "but excepted is : ", expectedPlayerName)
	}
	if actual.mark != expectedPlayerMark {
		t.Error("Actual is ", actual.mark, "but excepted is : ", expectedPlayerMark)

	}
}

func TestGetName(t *testing.T) {
	expectedPlayerName := "anupam"
	actualVal := actual.GetName()
	if actualVal != expectedPlayerName {
		t.Error("Actual is ", actualVal, "but excepted is : ", expectedPlayerName)
	}
}

func TestGetPlayerMark(t *testing.T) {
	expectedPlayerMark := "X"
	actualVal := actual.GetPlayerMark()
	if actualVal != expectedPlayerMark {
		t.Error("Actual is ", actualVal, "but excepted is : ", expectedPlayerMark)
	}
}
