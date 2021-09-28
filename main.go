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
	var size uint64
	var num, player1, player2, pos, mark1, mark2 string
LOOP:
	for true {
		fmt.Println("Enter the size of Board between 3 to 7")
		fmt.Scanln(&num)
		size, _ = strconv.ParseUint(num, 10, 32)
		err := errorhandling.SizeOfBoardHandling(size)
		if err == nil {
			break
		} else {
			fmt.Println(err)
		}
	}
	fmt.Println("Enter the player1 name: ")
	fmt.Scanln(&player1)
	player1 = notNullHandling(player1)
	fmt.Println(player1, "choose any alphabet mark : ")
	fmt.Scanln(&mark1)
	mark1 = checkCorrectMark(mark1, player1)
	p1 := player.New(player1, mark1)
	fmt.Println("Enter the player2 name: ")
	fmt.Scanln(&player2)
	for {
		player21 := notNullHandling(player2)
		player22 := checkDuplicatePlayerData(player1, player21)
		if player21 == player22 {
			player2 = player21
			break
		}
		player2 = player22
	}
	fmt.Println(player2, "choose any alphabet mark : ")
	fmt.Scanln(&mark2)
	for {
		mark21 := checkCorrectMark(mark2, player2)
		mark22 := checkDuplicatePlayerData(mark1, mark21)
		if mark21 == mark22 {
			mark2 = mark21
			break
		}
		mark2 = mark22
	}

	p2 := player.New(player2, mark2)
	p := [2]*player.Player{p1, p2}
	g := game.New(uint8(size), p)
	fmt.Println("________________Game Started__________________")
	g.PrintBoard()
	gameRes := resultanalyser.GameProgress
	var cellNo uint64
	user := g.CurrentUser()
	for gameRes == resultanalyser.GameProgress {
		user = g.CurrentUser()
		fmt.Println(user.GetName(), ", Enter the position between 1 to ", size*size)
		fmt.Scanln(&pos)
		cellNo, _ = strconv.ParseUint(pos, 10, 32)
		err := errorhandling.WrongNumberInsert(size, cellNo)
		if err != nil {
			fmt.Println(err)
		}
		res, err := g.Play(uint8(cellNo))
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
	fmt.Println("Type \"Y\" for again Playing \n \"N\" for End the game")
	fmt.Scanln(&reStart)
	if strings.ToUpper(reStart) == check {
		clearScreen()
		goto LOOP
	} else {
		clearScreen()
		fmt.Println("\n\n\n\n\n          THANK YOU FOR PLAYING             ")
	}
}

func clearScreen() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

func checkCorrectMark(mark, player string) string {
	for {
		er := errorhandling.ValidMark(mark)
		if er == nil {
			break
		} else {
			fmt.Println(player, er)
			fmt.Scanln(&mark)
		}
	}
	return mark
}

func checkDuplicatePlayerData(dataP1, dataP2 string) string {
	for {
		er := errorhandling.ValidNamePlayerMark(dataP1, dataP2)
		if er == nil {
			break
		} else {
			fmt.Println(er)
			fmt.Scanln(&dataP2)
		}
	}
	return dataP2
}

func notNullHandling(player string) string {
	for {
		player = strings.TrimSpace(player)
		er := errorhandling.NotNullName(player)
		if er == nil {
			break
		} else {
			fmt.Println(er)
			fmt.Scanln(&player)
		}
	}
	return player
}
