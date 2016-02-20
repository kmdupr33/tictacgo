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
	gamePlayer gamePlayer
	won        bool
}{
	{gamePlayer: func(g *Game) {},
		won: false},
	{gamePlayer: func(g *Game) {

		posToPlay := []Position{{0, 0},
			{1, 1}, //player 2
			{0, 1},
			{1, 2}, //player 2
			{0, 2}}

		for _, p := range posToPlay {
			err := g.PlayTurn(p)
			if err != nil {
				panic(err)
			}
		}
	},
		won: true},
	{gamePlayer: func(g *Game) {

		g.PlayTurn(Position{0, 0})

		// Player two plays
		g.PlayTurn(Position{1, 1})

		g.PlayTurn(Position{2, 2})

		// Player two plays
		g.PlayTurn(Position{1, 2})

		g.PlayTurn(Position{0, 2})
	},
		won: false},
}

func TestIsGameWon(t *testing.T) {
	t.SkipNow() //not implemented yet
	for _, tt := range gameWonTests {
		g := NewGame()
		tt.gamePlayer(g)
		w := g.IsGameWon()
		if w != tt.won {
			t.Logf("Supposed winner: %v", g.Winner())
			t.Logf("\n%v", g.board)
			t.Errorf("Expected: %v Got: %v", tt.won, w)
		}
	}
}
