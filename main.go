package main

import (
	"fmt"
	"strconv"

	"github.com/anupsingh31/tictactoe/errorhandling"
	"github.com/anupsingh31/tictactoe/game"
	"github.com/anupsingh31/tictactoe/player"
	"github.com/anupsingh31/tictactoe/resultanalyser"
)

func main() {
	var size int64
	var num, player1, player2, pos string
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
	p1 := player.NewPlayer(player1, "X")
	fmt.Println("Enter the player2 name: ")
	fmt.Scanln(&player2)
	p2 := player.NewPlayer(player2, "O")
	p := [2]*player.Player{p1, p2}
	g := game.NewGame(int32(size), p)
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
		fmt.Println(user.GetName(), " won")
	}
}
