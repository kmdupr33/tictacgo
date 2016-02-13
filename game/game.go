package game

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
	p1 := new(Player)
	p2 := new(Player)
	return &Game{board: NewBoard(),
		players:       []*Player{p1, p2},
		currentPlayer: p1}
}

//Game represents a game of tictacto
type Game struct {
	board         Board
	currentPlayer *Player
	players       []*Player
	winner        *Player
	turn          int
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
	return true
}

func (g *Game) IsCatsGame() bool {
	return g.board.IsFull()
}

func (g *Game) Winner() *Player {
	return g.winner
}

func (g *Game) playTurn() {
	g.turn++
}

func (g *Game) CurrentPlayer() *Player {
	return g.currentPlayer
}
