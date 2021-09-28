package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/anupsingh31/tictactoe/errorhandling"
	"github.com/anupsingh31/tictactoe/game"
	"github.com/anupsingh31/tictactoe/player"
	"github.com/anupsingh31/tictactoe/resultanalyser"
)

func main() {
	var size int64
	var num, player1, player2, pos, mark string
LOOP:
	for true {
		fmt.Println("Enter the size of Board between 3 to 7")
		fmt.Scanln(&num)
		size, _ = strconv.ParseInt(num, 10, 32)
		err := errorhandling.SizeOfBoardHandling(size)
		if err == nil {
			break
		} else {
			fmt.Println(err)
		}
	}
	fmt.Println("Enter the player1 name: ")
	fmt.Scanln(&player1)
	fmt.Println(player1, "choose you mark either \"X\" or \"O\"")
	fmt.Scanln(&mark)
	for true {
		mark = strings.ToUpper(mark)
		er := errorhandling.ValidMark(mark)
		if er == nil {
			break
		} else {
			fmt.Println(player1, er)
			fmt.Scanln(&mark)
		}
	}
	p1 := player.New(player1, mark)
	fmt.Println("Enter the player2 name: ")
	fmt.Scanln(&player2)
	for true {
		er := errorhandling.ValidNamePlayer(player1, player2)
		if er == nil {
			break
		} else {
			fmt.Println(er)
			fmt.Scanln(&player2)
		}
	}
	if strings.EqualFold(mark, "X") {
		fmt.Println(player2, " you mark should be \"O\" ")
		mark = "O"
	} else {
		fmt.Println(player2, " you mark should be \"X\" ")
		mark = "X"
	}
	p2 := player.New(player2, mark)
	p := [2]*player.Player{p1, p2}
	g := game.New(int32(size), p)
	fmt.Println("________________Game Started__________________")
	g.PrintBoard()
	gameRes := resultanalyser.GameProgress
	var cellNo int64
	user := g.CurrentUser()
	for gameRes == resultanalyser.GameProgress {
		user = g.CurrentUser()
		fmt.Println(user.GetName(), ", Enter the position between 1 to ", size*size)
		fmt.Scanln(&pos)
		cellNo, _ = strconv.ParseInt(pos, 10, 32)
		err := errorhandling.WrongNumberInsert(size, cellNo)
		if err != nil {
			fmt.Println(err)
		}
		res, err := g.Play(int32(cellNo))
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		gameRes = res
		g.PrintBoard()
	}
	switch gameRes {
	case resultanalyser.GameDraw:
		fmt.Println("Game was draw")
	case resultanalyser.GameWin:
		fmt.Println(user.GetName(), " Won")
	}
	var reStart string
	const check string = "Y"
	fmt.Println("Enter \"Y\" for Play and \"N\" for End the game")
	fmt.Scanln(&reStart)
	if strings.ToUpper(reStart) == check {
		fmt.Print("\033[H\033[2J")
		clearScreen()
		goto LOOP
	} else {
		fmt.Print("\033[H\033[2J")
		clearScreen()
		fmt.Println("\n\n\n\n\n          THANK YOU FOR PLAYING             ")
	}
}

func clearScreen() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}
