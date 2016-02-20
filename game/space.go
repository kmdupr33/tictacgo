package game

import "fmt"

//-------------------------------------------------
// Space
//-------------------------------------------------

//Space represents a spot within the gameboard that can be filled with a player marker
type Space struct {
	marker Marker
	fmt.Stringer
}

func (s *Space) String() string {
	return s.marker.String()
}

//IsEmpty returns whether this space has a PlayerMarker in it
func (s *Space) IsEmpty() bool {
	return s.marker == 0
}
