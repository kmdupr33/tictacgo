package game

import "fmt"

//-------------------------------------------------
// Space
//-------------------------------------------------

//Space represents a spot within the gameboard that can be filled with a player marker
type Space struct {
	marker    Marker
	neighbors []*Space
	fmt.Stringer
}

//check() checks whether there are adjacent spaces that have the same
//player marker as the player marker within this space
func (s *Space) check() []*Space {
	var matchingNeighbors []*Space
	for _, n := range s.neighbors {
		if s.marker == n.marker {
			matchingNeighbors = append(matchingNeighbors, n)
		}
	}
	return matchingNeighbors
}

func (s *Space) String() string {
	return s.marker.String()
}

//IsEmpty returns whether this space has a PlayerMarker in it
func (s *Space) IsEmpty() bool {
	return s.marker == 0
}
