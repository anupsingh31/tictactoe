package game

import (
	"fmt"

	"github.com/anupsingh31/tictactoe/board"
	"github.com/anupsingh31/tictactoe/player"
	"github.com/anupsingh31/tictactoe/resultanalyser"
)

type Game struct {
	b       *board.Board
	ra      *resultanalyser.ResultAnalyzer
	players [2]*player.Player
	turn    int32
}

func NewGame(boardSize int32, players [2]*player.Player) *Game {
	return &Game{
		b:       board.NewBoard(boardSize),
		ra:      resultanalyser.NewResultAnalyzer(),
		players: players,
		turn:    0,
	}
}

func (g *Game) Play(cellId int32) (resultanalyser.GameStatus, error) {
	currentMark := g.players[g.turn].GetPlayerMark()
	err := g.b.SetMarkCell(cellId, currentMark)
	if err != nil {
		return resultanalyser.GameProgress, err
	}
	g.changeTurn()
	sta, err := g.ra.GetBoardStatus(g.b, currentMark, cellId), nil
	fmt.Println(sta)
	return sta, err
}

func (g *Game) changeTurn() {
	g.turn = (g.turn + 1) % int32(len(g.players))
}

func (g *Game) CurrentUser() *player.Player {
	return g.players[g.turn]
}

func (g *Game) PrintBoard() {
	fmt.Println(g.b)
}
