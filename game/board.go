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
