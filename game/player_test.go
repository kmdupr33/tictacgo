package game

import "testing"

func TestPlayer_RemovePossibleWins(t *testing.T) {
	p := NewPlayer(X)
	t.Logf("Iniital winning positions: %v", p.possibWinningPos)
	posToRemove := Position{0, 0}
	p.removePossibleWins(posToRemove)
	for _, wp := range p.possibWinningPos {
		for _, pos := range wp {
			if pos == posToRemove {
				t.Errorf("expected: %v to be removed but found it in winning play %v", posToRemove, wp)
				t.Logf("winning positions: %v", p.possibWinningPos)
			}
		}
	}
}
