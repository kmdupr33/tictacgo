package game

import "fmt"

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

//Player is a tictacto player
type Player struct {
	marker           Marker
	playedPositions  []Position
	possibWinningPos []winningPlacement
}

func NewPlayer(m Marker) *Player {
	return &Player{marker: m,
		possibWinningPos: winningPositions[:]}
}

func (p *Player) removePossibleWins(pos Position) {
	for i, wp := range p.possibWinningPos {
		fmt.Printf("winning position: %v at index: %d\n", wp, i)
		for _, wpos := range wp {
			fmt.Printf("examining winning pos: %v", p.possibWinningPos[i])
			if wpos == pos {
				fmt.Println("removePossibleWins(): Found matching winning position")
				p.possibWinningPos = append(p.possibWinningPos[:i],
					p.possibWinningPos[i+1:]...)
			}
		}
	}

}
