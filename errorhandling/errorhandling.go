package errorhandling

import (
	"errors"
)

func SizeOfBoardHandling(size uint64) error {
	if size >= 3 && size <= 7 {
		return nil
	}
	return errors.New("Out of Range Size board")
}

func WrongNumberInsert(size, limit uint64) error {
	if limit >= 1 && limit <= size*size {
		return nil
	} else {
		return errors.New("You have inserted wrong number")
	}
}

func ValidNamePlayerMark(player1, player2 string) error {
	if player1 == player2 {
		return errors.New("You have inserted player1 data.\nEnter again...")
	}
	return nil
}

func ValidMark(mark string) error {
	if (mark >= "a" && mark <= "z" && len(mark) == 1) || (mark >= "A" && mark <= "Z" && len(mark) == 1) {
		return nil
	}
	return errors.New("Please, Enter single Alphabet only")
}

func NotNullName(name string) error {
	if len(name) != 0 {
		return nil
	}
	return errors.New("Please, Enter you name. Space not supported")
}
