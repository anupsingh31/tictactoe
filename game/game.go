package game

import (
	"fmt"

	"github.com/anupsingh31/tictactoe/board"
	"github.com/anupsingh31/tictactoe/player"
	"github.com/anupsingh31/tictactoe/resultanalyser"
)

type Game struct {
	b                *board.Board
	ra               *resultanalyser.ResultAnalyzer
	players          [2]*player.Player
	turn, limitmoved uint8
}

func New(boardSize uint8, players [2]*player.Player) *Game {
	return &Game{
		b:          board.New(boardSize),
		ra:         resultanalyser.New(),
		players:    players,
		turn:       0,
		limitmoved: 0,
	}
}

func (g *Game) Play(cellId uint8) (resultanalyser.GameStatus, error) {
	currentMark := g.players[g.turn].GetPlayerMark()
	err := g.b.SetMarkCell(cellId, currentMark)
	if err != nil {
		return resultanalyser.GameProgress, err
	}
	g.changeTurn()
	g.limitmoved += 1
	if g.limitmoved >= 2*g.b.GetSizeBoard()-1 {
		status := g.ra.GetBoardStatus(g.b, currentMark, cellId)
		return status, err
	}
	return resultanalyser.GameProgress, err
}

func (g *Game) changeTurn() {
	g.turn = (g.turn + 1) % uint8(len(g.players))
}

func (g *Game) CurrentUser() *player.Player {
	return g.players[g.turn]
}

func (g *Game) PrintBoard() {
	fmt.Println(g.b)
}
