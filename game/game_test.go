package game

import "testing"

type markerPlacer func(Board)

var printBoardTests = []struct {
	markerPlacer markerPlacer
	output       string
}{
	{markerPlacer: func(b Board) {},
		output: "|---|---|---|\n|   |   |   |\n|---|---|---|\n|   |   |   |\n|---|---|---|\n|   |   |   |\n|---|---|---|\n"},
	{markerPlacer: func(b Board) {
		b.PlaceMarker(Position{0, 0}, X)
	},
		output: "|---|---|---|\n| X |   |   |\n|---|---|---|\n|   |   |   |\n|---|---|---|\n|   |   |   |\n|---|---|---|\n"},
}

func TestBoardStringer(t *testing.T) {
	for _, tt := range printBoardTests {
		b := NewBoard()
		tt.markerPlacer(b)
		o := b.String()
		if tt.output != o {
			t.Errorf("\nExpected:\n %v \nGot:\n %v", tt.output, o)
		}
	}

}

var gameWonTests = []struct {
	markerPlacer markerPlacer
	won          bool
}{
	{markerPlacer: func(b Board) {},
		won: false},
	{markerPlacer: func(b Board) {
		b.PlaceMarker(Position{0, 0}, X)
		b.PlaceMarker(Position{0, 1}, X)
		b.PlaceMarker(Position{0, 2}, X)
	},
		won: true},
	{markerPlacer: func(b Board) {
		b.PlaceMarker(Position{0, 0}, X)
		b.PlaceMarker(Position{1, 0}, X)
		b.PlaceMarker(Position{2, 0}, X)
	},
		won: true},
}

func TestIsGameWon(t *testing.T) {
	for _, tt := range gameWonTests {
		g := NewGame()
		tt.markerPlacer(g.board)
		w := g.isGameWon()
		if w != tt.won {
			t.Errorf("Expected: %v Got: %v", tt.won, w)
		}
	}

}
