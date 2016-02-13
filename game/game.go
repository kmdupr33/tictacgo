package game

import (
	"bufio"
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
