package game

//-------------------------------------------------
// Game
//-------------------------------------------------

//NewGame returns a new game of tictacto
func NewGame() *Game {
	p1 := &Player{marker: X}
	p2 := &Player{marker: O}
	return &Game{board: NewBoard(),
		players: []*Player{p1, p2}}
}

//Game represents a game of tictacto
type Game struct {
	board              Board
	currentPlayerIndex int
	players            []*Player
	winner             *Player
	turn               int
	//setups maps a position on the board to a slice of positions that
	//, if played, would put a player one move away from winning the game
	setups map[Position][]Position
}

//Position represents a positon on a tictacto grid
type Position struct {
	x, y int
}

func (g *Game) IsGameWon() bool {
	if g.turn < 4 {
		return false
	}

	g.isGameWon(g.board.SpaceAt(Position{0, 0}))

	return false
}

func (g *Game) isGameWon(s *Space) bool {
	return false
}

func (g *Game) IsCatsGame() bool {
	return g.board.IsFull()
}

func (g *Game) Winner() *Player {
	return g.winner
}

//PlayTurn places a marker for the current player at the
//position passed in. It returns an error if this method is called when the it is
//not this player's turn to play.
func (g *Game) PlayTurn(p Position) error {
	cp := g.CurrentPlayer()
	err := g.board.PlaceMarker(p, cp.marker)
	if err != nil {
		return err
	}
	g.turn++

	//Update current player
	//g.currentPlayerIndex should only ever be 0 or 1
	g.currentPlayerIndex = -g.currentPlayerIndex + 1

	return nil
}

func (g *Game) updateSetups(p Position) {

}

func (g *Game) CurrentPlayer() *Player {
	return g.players[g.currentPlayerIndex]
}
