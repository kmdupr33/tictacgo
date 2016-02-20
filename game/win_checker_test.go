package game

import (
	"reflect"
	"testing"
)

var setupTests = []struct {
	in  Position
	out map[Position][]Position
}{
	{
		in: Position{0, 0},
		out: map[Position][]Position{
			{0, 1}: {{0, 2}},
			{1, 1}: {{2, 0}},
			{1, 2}: {{2, 2}},
			{0, 0}: {{0, 1}},
			{2, 0}: {{1, 1}},
			{2, 2}: {{1, 2}},
		},
	},
}

func TestGame_updateSetups(t *testing.T) {
	for _, tt := range setupTests {
		m := mapWinChecker{}
		m.updateSetups(tt.in)
		if !reflect.DeepEqual(m.setups, tt.out) {
			t.Errorf("mapWinChecker.updateSetups(%v): %v expected: %v", tt.in, m.setups, tt.out)
		}
	}
}
