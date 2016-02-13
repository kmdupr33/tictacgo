package game

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

//-------------------------------------------------
// PlayerMarker
//-------------------------------------------------

//Player is a tictacto player
type Player struct {
	marker       Marker
	playedSpaces []Position
}

//Marker is an "X" or an "O" within the tictacto grid
type Marker int

//Player Markers
const (
	_        = iota
	X Marker = iota
	O Marker = iota
)

func (m Marker) String() string {
	switch m {
	case X:
		return "X"
	case O:
		return "O"
	default:
		return " "
	}
}

//-------------------------------------------------
// Board
//-------------------------------------------------

//Board represents the "grid" where Player's markers can be placed
type Board interface {
	IsFull() bool
	MarkerAt(p Position) Marker
	PlaceMarker(pos Position, m Marker) error
	fmt.Stringer
}

type arrayBoard [3][3]*Space

func (b *arrayBoard) IsFull() bool {
	for i := 0; i < 3; i++ {
		for _, s := range b[i] {
			if s.marker == Marker(0) {
				return false
			}
		}
	}
	return true
}

func (b *arrayBoard) MarkerAt(p Position) Marker {
	return b[p.x][p.y].marker
}

func (b *arrayBoard) PlaceMarker(pos Position, m Marker) error {
	s := b[pos.x][pos.y]
	if !s.IsEmpty() {
		return errors.New("game: space already occupied")
	}
	s.marker = m
	return nil
}

func (b arrayBoard) String() string {
	return fmt.Sprintf("|---|---|---|\n| %v | %v | %v |\n|---|---|---|\n| %v | %v | %v |\n|---|---|---|\n| %v | %v | %v |\n|---|---|---|\n",
		b[0][0], b[1][0], b[2][0],
		b[1][0], b[1][1], b[1][2],
		b[2][0], b[2][1], b[2][2])
}

//NewBoard creates a new empty game board
func NewBoard() Board {

	tls := new(Space)
	tms := new(Space)
	trs := new(Space)

	mls := new(Space)
	mms := new(Space)
	mrs := new(Space)

	bls := new(Space)
	bms := new(Space)
	brs := new(Space)

	a := &arrayBoard{
		{tls, tms, trs},
		{mls, mms, mrs},
		{bls, bms, brs},
	}

	tls.neighbors = append(tls.neighbors, tms, mls)
	tms.neighbors = append(tms.neighbors, tls, trs, mms)
	trs.neighbors = append(trs.neighbors, tms, mrs)

	mls.neighbors = append(mls.neighbors, tls, bls, mms)
	mms.neighbors = append(mms.neighbors, tms, bms, mls, mrs)
	mrs.neighbors = append(mrs.neighbors, mms, trs, brs)

	bls.neighbors = append(bls.neighbors, bms, mls)
	bms.neighbors = append(mls.neighbors, bls, brs, mms)
	brs.neighbors = append(brs.neighbors, mrs, bms)

	return a
}

//-------------------------------------------------
// Game
//-------------------------------------------------

//NewGame returns a new game of tictacto
func NewGame() *Game {
	return &Game{board: NewBoard(), currentPlayer: Player{}}
}

//Game represents a game of tictacto
type Game struct {
	board         Board
	currentPlayer Player
	winner        Player
	turn          int
}

//Play starts a game of tictacto
func (g Game) Play() {
	fmt.Println("A new game has started! Type 'help' for instructions on how to play")
	for !g.isGameWon() || !g.isCatsGame() {
		g.playTurn()
	}
	fmt.Printf("%v's game!", g.winner)
}

//Position represents a positon on a tictacto grid
type Position struct {
	x, y int
}

type winningPlacement [3]Position

var winningPositions = [8]winningPlacement{
	{{0, 0}, {0, 1}, {0, 2}},
	{{0, 0}, {1, 0}, {2, 0}},
	{{0, 0}, {1, 1}, {2, 2}},
	{{1, 0}, {1, 1}, {1, 2}},
	{{2, 0}, {2, 1}, {2, 2}},
	{{2, 0}, {1, 1}, {0, 2}},
	{{0, 1}, {1, 1}, {2, 1}},
	{{0, 2}, {1, 2}, {2, 2}},
}

func (g Game) isGameWon() bool {
	if g.turn < 4 {
		return false
	}
	return true
}

func (g Game) isCatsGame() bool {
	return g.board.IsFull()
}

func (g Game) playTurn() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%v's Turn: ", g.currentPlayer)
	text, _ := reader.ReadString('\n')
	if text == "help" {
		printInstructions()
	}
	fmt.Println(text)
}

func printInstructions() {

}
