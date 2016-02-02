package game

import "fmt"

//-------------------------------------------------
// Space
//-------------------------------------------------

//NeighborChecker checks to see whether there are adjacent spaces that have the same
//player marker as the player marker within the NeighborChecker's space
type NeighborChecker interface {
	Check() []Space
}

type topLeftSpace struct {
	Space
	southSpace Space
	eastSpace  Space
}

type topMidSpace struct {
	Space
	northSpace Space
	eastSpace  Space
	westSpace  Space
}

type topRightSpace struct {
	Space
	southSpace Space
	westSpace  Space
}

type midRightSpace struct {
	Space
	northSpace Space
	southSpace Space
	westSpace  Space
}

type botRightSpace struct {
	Space
	northSpace Space
	westSpace  Space
}

type botMidSpace struct {
	Space
	northSpace Space
	eastSpace  Space
	westSpace  Space
}

type botLeftSpace struct {
	Space
	northSpace Space
	eastSpace  Space
}

type midLeftSpace struct {
	Space
	northSpace Space
	southSpace Space
	eastSpace  Space
}

type midMidSpace struct {
	Space
	northSpace Space
	southSpace Space
	eastSpace  Space
	westSpace  Space
}

//Space represents a spot within the gameboard that can be filled with a player marker
type Space interface {
	fmt.Stringer
	IsEmpty() bool
	Check() []Space
	Marker() Marker
	SetMarker(m Marker)
}

type baseSpace struct {
	Marker Marker
}

func (s baseSpace) String() string {
	return s.Marker.String()
}

//IsEmpty returns whether this space has a PlayerMarker in it
func (s baseSpace) IsEmpty() bool {
	return s.Marker == 0
}
