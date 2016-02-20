package game

//WinChecker checks to see whether a game has a winner
type WinChecker interface {
	//WinChecker expects this method to be called everytime a turn has
	//occurred in a game
	TurnPlayed(p Player, pos Position)
	//Winner returns the player who has won the game or nil if noone has
	//won yet
	Winner() *Player
}

type mapWinChecker struct {
	//setups maps a position on the board to a slice of positions that
	//, if played, would put a player one move away from winning the game
	setups map[Position][]Position
}

func (m *mapWinChecker) TurnPlayed(p Player, pos Position) {

}

func (m *mapWinChecker) updateSetups(pos Position) {

}
