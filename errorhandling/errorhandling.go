package errorhandling

import (
	"errors"
)

func SizeOfBoardHandling(size int) error {
	if size >= 3 && size <= 7 {
		return nil
	}
	return errors.New("Out of Range Size board")
}

func SlotAlreadyTaken() error {
	return errors.New("Slot not Avilable")
}

func WrongNumberInsert(size, limit int) error {
	if limit >= 1 && limit <= size*size {
		return nil
	} else {
		return errors.New("You have inserted wrong number")
	}
}
