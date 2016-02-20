package game

import (
	"fmt"
	"log"
)

//-------------------------------------------------
// Game
//-------------------------------------------------

//New returns a new game of tictacto
func New() *Game {
	log.Print("New Game started")
	p1 := &Player{marker: X}
	p2 := &Player{marker: O}
	return &Game{board: NewBoard(),
		players:    []*Player{p1, p2},
		winChecker: NewWinChecker()}
}

//Game represents a game of tictacto
type Game struct {
	board              Board
	currentPlayerIndex int
	players            []*Player
	winner             *Player
	turn               int
	winChecker         WinChecker
}

func (g *Game) String() string {
	string := g.board.String()
	if !g.IsWon() {
		return string +
			fmt.Sprintf("%v's Turn: ", g.CurrentPlayer())
	} else if !g.IsCatsGame() {
		return string + fmt.Sprintf("%v's game!\n", g.Winner())
	}
	return string + fmt.Sprintf("Cat's game!")
}

//Position represents a positon on a tictacto grid valid x and y values
//are [0,2]
type Position struct {
	X, Y int
}

//IsWon returns whether the game has a winner already
func (g *Game) IsWon() bool {
	if g.turn < 4 {
		return false
	}

	return g.winChecker.Winner() != nil
}

//IsCatsGame returns whether the game has a winner
func (g *Game) IsCatsGame() bool {
	return g.board.IsFull()
}

//Winner returns the winner of the game or nil if there is no
//winner
func (g *Game) Winner() *Player {
	return g.winChecker.Winner()
}

//PlayTurn places a marker for the current player at the
//position passed in. It returns an error if
//the space at the position passed in has already been filled
//or if the position passed in has invalid x or y coordinates
func (g *Game) PlayTurn(p Position) error {
	cp := g.CurrentPlayer()
	err := g.board.PlaceMarker(p, cp.marker)
	if err != nil {
		return err
	}
	g.turn++

	g.winChecker.TurnPlayed(cp, p)

	//Update current player
	//g.currentPlayerIndex should only ever be 0 or 1
	g.currentPlayerIndex = -g.currentPlayerIndex + 1

	return nil
}

//CurrentPlayer returns the Player whose turn is currently active
func (g *Game) CurrentPlayer() *Player {
	return g.players[g.currentPlayerIndex]
}
