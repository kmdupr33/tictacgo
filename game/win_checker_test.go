package game

import (
	"reflect"
	"testing"
)

var pl1 = &Player{X}

var setupTests = []struct {
	pl  *Player
	pos Position
	out map[Position][]winningPlay
}{
	{
		pl:  pl1,
		pos: Position{0, 0},
		out: map[Position][]winningPlay{
			{0, 1}: {{pl1, Position{0, 2}}},
			{1, 1}: {{pl1, Position{2, 2}}},
			{2, 0}: {{pl1, Position{1, 0}}},
			{2, 2}: {{pl1, Position{1, 1}}},
			{0, 2}: {{pl1, Position{0, 1}}},
			{1, 0}: {{pl1, Position{2, 0}}},
		},
	},
}

func TestGame_updateSetups(t *testing.T) {
	for _, tt := range setupTests {
		setups := make(map[Position][]winningPlay)
		winningPlay := make(map[Position]*Player)
		m := mapWinChecker{setups: setups, winningPlay: winningPlay}
		m.updateSetups(tt.pl, tt.pos)
		if !reflect.DeepEqual(m.setups, tt.out) {
			t.Errorf("mapWinChecker.updateSetups(%v, %v):\n %v \nexpected:\n %v\n",
				tt.pl, tt.pos, m.setups, tt.out)
		}
	}
}
