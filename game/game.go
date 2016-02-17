package game

import "fmt"

//Player is a tictacto player
type Player struct {
	marker                   Marker
	playedSpaces             []Position
	possibleWinningPositions []winningPlacement
}

func (p *Player) NewPlayer(m Marker) *Player {
	return &Player{marker: m,
		possibleWinningPositions: winningPositions[:]}
}

//-------------------------------------------------
// Marker
//-------------------------------------------------

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
	p1 := &Player{marker: X}
	p2 := &Player{marker: O}
	return &Game{board: NewBoard(),
		players: []*Player{p1, p2}}
}

//Game represents a game of tictacto
type Game struct {
	board              Board
	currentPlayerIndex int
	players            []*Player
	winner             *Player
	turn               int
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

func (g *Game) IsGameWon() bool {
	if g.turn < 4 {
		return false
	}

	for _, p := range g.players {
		for _, ps := range p.playedSpaces {
			sp := g.board.SpaceAt(ps)
			neighbors := sp.check()
			fmt.Println(neighbors)
			for _, n := range neighbors {
				for _, nn := range n.check() {
					if nn != n {
						g.winner = p
						return true
					}
				}
			}
		}
	}

	return false
}

func (g *Game) IsCatsGame() bool {
	return g.board.IsFull()
}

func (g *Game) Winner() *Player {
	return g.winner
}

//PlayTurn places a marker for the current player at the
//position passed in. It returns an error if this method is called when the it is
//not this player's turn to play.
func (g *Game) PlayTurn(p Position) error {
	cp := g.CurrentPlayer()
	err := g.board.PlaceMarker(p, cp.marker)
	if err != nil {
		return err
	}
	g.turn++
	cp.playedSpaces = append(cp.playedSpaces, p)
	//g.currentPlayerIndex should only ever be 0 or 1
	g.currentPlayerIndex = -g.currentPlayerIndex + 1
	return nil
}

func (g *Game) CurrentPlayer() *Player {
	return g.players[g.currentPlayerIndex]
}
