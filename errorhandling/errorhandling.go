package errorhandling

import (
	"errors"
)

func SizeOfBoardHandling(size int64) error {
	if size >= 3 && size <= 7 {
		return nil
	}
	return errors.New("Out of Range Size board")
}

func WrongNumberInsert(size, limit int64) error {
	if limit >= 1 && limit <= size*size {
		return nil
	} else {
		return errors.New("You have inserted wrong number")
	}
}

func ValidNamePlayer(player1, player2 string) error {
	if player1 == player2 {
		return errors.New("Enter valid player2 name")
	}
	return nil
}

func ValidMark(mark string) error {
	if mark == "X" || mark == "O" {
		return nil
	}
	return errors.New("Enter \"X\" or \"O\" only")
}
