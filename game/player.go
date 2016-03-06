package game

import "errors"

//-------------------------------------------------
// Marker
//-------------------------------------------------

//Marker is an "X" or an "O" within the tictacto grid
type Marker int

//Player Markers
const (
	_ Marker = iota
	X
	O
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

//Player is a tictacto player
type Player struct {
	marker Marker
	brain  ComputerPlayerBrain
}

func NewPlayer(m Marker) *Player {
	return &Player{marker: m}
}

//NextMove returns the next move that computer player should play
func (p *Player) NextMove() (Position, error) {
	if !p.IsComputer() {
		return Position{}, errors.New("Player is not a computer player")
	}
	return p.brain.getComputerPlayerMove(), nil
}

//IsComputer returns whether p is a computer player
func (p *Player) IsComputer() bool {
	return p.brain != nil
}

func (p *Player) String() string {
	return p.marker.String()
}
