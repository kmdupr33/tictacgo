package game

import "testing"

type markerPlacer func(Board)

var printBoardTests = []struct {
	markerPlacer markerPlacer
	output       string
}{
	{markerPlacer: func(b Board) {},
		output: "|---|---|---|\n|   |   |   |\n|---|---|---|\n|   |   |   |\n|---|---|---|\n|   |   |   |\n|---|---|---|\n",
	},
	{markerPlacer: func(b Board) {
		b.PlaceMarker(Position{0, 0}, X)
	},
		output: "|---|---|---|\n| X |   |   |\n|---|---|---|\n|   |   |   |\n|---|---|---|\n|   |   |   |\n|---|---|---|\n",
	},
	{markerPlacer: func(b Board) {
		b.PlaceMarker(Position{0, 0}, X)
		b.PlaceMarker(Position{0, 1}, X)
	},
		output: "|---|---|---|\n| X |   |   |\n|---|---|---|\n| X |   |   |\n|---|---|---|\n|   |   |   |\n|---|---|---|\n",
	},
	{markerPlacer: func(b Board) {
		b.PlaceMarker(Position{0, 1}, X)
	},
		output: "|---|---|---|\n|   |   |   |\n|---|---|---|\n| X |   |   |\n|---|---|---|\n|   |   |   |\n|---|---|---|\n",
	},
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

type gamePlayer func(g *Game)

var gameWonTests = []struct {
	name      string
	posToPlay []Position
	won       bool
}{
	{name: "Board empty",
		posToPlay: nil,
		won:       false},
	{name: "Win Should be True",
		posToPlay: []Position{{0, 0},
			{1, 1}, //player 2
			{0, 1},
			{1, 2}, //player 2
			{0, 2}},
		won: true},
	{name: "Win should be false",
		posToPlay: []Position{{0, 0},
			{1, 1}, // Player two plays
			{2, 2},
			{1, 2},
			{0, 2}}, // Player two plays
		won: false},
}

func TestGame_IsWon(t *testing.T) {
	for _, tt := range gameWonTests {
		t.Logf("Starting test called: %s", tt.name)
		g := New()
		for _, p := range tt.posToPlay {
			err := g.PlayTurn(p)
			if err != nil {
				panic(err)
			}
		}
		w := g.IsWon()
		if w != tt.won {
			t.Logf("Supposed winner: %v", g.Winner())
			t.Logf("\n%v", g.board)
			t.Errorf("Expected: %v Got: %v", tt.won, w)
		}
	}
}
