package game

import "fmt"

//-------------------------------------------------
// Game
//-------------------------------------------------

//NewGame returns a new game of tictacto
func New() *Game {
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
	return fmt.Sprintln(g.board) +
		fmt.Sprintf("%v's Turn: ", g.CurrentPlayer())
}

//Position represents a positon on a tictacto grid
type Position struct {
	X, Y int
}

func (g *Game) IsWon() bool {
	if g.turn < 4 {
		return false
	}

	return g.winChecker.Winner() != nil
}

func (g *Game) IsCatsGame() bool {
	return g.board.IsFull()
}

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
func (g *Game) CurrentPlayer() *Player {
	return g.players[g.currentPlayerIndex]
}
