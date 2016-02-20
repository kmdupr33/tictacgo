package game

//-------------------------------------------------
// Marker
//-------------------------------------------------

//Marker is an "X" or an "O" within the tictacto grid
type Marker int

//Player Markers
const (
	_ = iota
	//TODO Remove unnecessary iota
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

//Player is a tictacto player
type Player struct {
	marker Marker
}

func NewPlayer(m Marker) *Player {
	return &Player{marker: m}
}

func (p *Player) String() string {
	return p.marker.String()
}
