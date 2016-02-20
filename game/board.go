package game

import (
	"errors"
	"fmt"
)

//-------------------------------------------------
// Board
//-------------------------------------------------

//Board represents the "grid" where Player's markers can be placed
type Board interface {
	IsFull() bool
	SpaceAt(p Position) *Space
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

func (b *arrayBoard) SpaceAt(p Position) *Space {
	return b[p.X][p.Y]
}

func (b *arrayBoard) PlaceMarker(pos Position, m Marker) error {
	if !isValCoord(pos.X) || !isValCoord(pos.Y) {
		return errors.New("invalid coordinate")
	}
	s := b[pos.X][pos.Y]
	if !s.IsEmpty() {
		return errors.New("game: space already occupied")
	}
	s.marker = m
	return nil
}

func isValCoord(x int) bool {
	return x > -1 && x < 3
}

func (b arrayBoard) String() string {
	return fmt.Sprintf("|---|---|---|\n| %v | %v | %v |\n|---|---|---|\n| %v | %v | %v |\n|---|---|---|\n| %v | %v | %v |\n|---|---|---|\n",
		b[0][0], b[1][0], b[2][0],
		b[0][1], b[1][1], b[2][1],
		b[0][2], b[1][2], b[2][2])
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

	return a
}
