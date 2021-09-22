package main

import (
	"fmt"

	"github.com/anupsingh31/tictactoe/board"
	"github.com/anupsingh31/tictactoe/errorhandling"
	"github.com/anupsingh31/tictactoe/resultanalyser"
)

func main() {
	var size, position, row, column int
	for true {
		fmt.Println("Enter the size of Board between 3 to 7")
		fmt.Scanln(&size)
		err := errorhandling.SizeOfBoardHandling(size)
		if err == nil {
			break
		} else {
			fmt.Println(err)
		}
	}
	board.BoardFrame(size)
	var result string = "progress"
	var err error
	var turnPlayer1 bool = true
	player1 := []string{}
	var name, mark string
	for result == "progress" {
		if len(player1) == 0 {
			fmt.Println("Enter the player 1 name ")
			fmt.Scanln(&name)
		} else if len(player1) == 1 {
			fmt.Println("Enter the player 2 name ")
			fmt.Scanln(&name)
		}
		player1 = append(player1, name)
		for true {
			if turnPlayer1 {
				fmt.Println(player1[0], "enter the position of board between 1 to", size*size)
			} else {
				fmt.Println(player1[1], "enter the position of board between 1 to", size*size)
			}
			fmt.Scanln(&position)
			err = errorhandling.WrongNumberInsert(size, position)
			if err == nil {
				break
			} else {
				fmt.Println(err)
			}
		}
		if turnPlayer1 {
			mark = "X "
			turnPlayer1 = false
		} else {
			mark = "O "
			turnPlayer1 = true
		}
		row, column, err = board.SetGrid(mark, position)
		if err == nil {
			result = resultanalyser.ResultAnalyser(mark, row, column)
		} else {
			fmt.Println(err)
			turnPlayer1 = !turnPlayer1
			continue
		}
		printBoard()
		if result == "win" {
			if !turnPlayer1 {
				fmt.Println(player1[0], " Win")
			} else {
				fmt.Println(player1[1], " Win")

			}
		} else if result == "draw" {
			fmt.Println("It's Draw ")

		}
	}

}

func printBoard() {
	for _, row := range board.Boardgrid {
		for _, val := range row {
			fmt.Print(val)
		}
		fmt.Println()
	}
}
