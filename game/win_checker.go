package game

func NewWinChecker() WinChecker {
	setups := make(map[Position][]winningPlay)
	winningPlay := make(map[Position]*Player)
	return &mapWinChecker{setups: setups, winningPlay: winningPlay}
}

//WinChecker checks to see whether a game has a winner
type WinChecker interface {
	//WinChecker expects this method to be called everytime a turn has
	//occurred in a game
	TurnPlayed(p *Player, pos Position)
	//Winner returns the player who has won the game or nil if noone has
	//won yet
	Winner() *Player
}

type winningPlay struct {
	//TODO Stop passing around pointers to Players
	player *Player
	pos    Position
}

type mapWinChecker struct {
	//setups maps a position on the board to a slice of positions that
	//, if played, would put a player one move away from winning the game
	setups map[Position][]winningPlay

	winner *Player

	winningPlay map[Position]*Player
}

func (m *mapWinChecker) Winner() *Player {
	return m.winner
}

func (m *mapWinChecker) TurnPlayed(p *Player, pos Position) {
	if w, ok := m.winningPlay[pos]; ok {
		m.winner = w
		return
	}

	if wps, ok := m.setups[pos]; ok {
		for _, wp := range wps {
			if wp.player != p {
				break
			}
			m.winningPlay[wp.pos] = wp.player
		}
	}

	m.updateSetups(p, pos)
}

type setupPosition struct {
	setupPos   Position
	winningPos Position
}

func (m *mapWinChecker) updateSetups(p *Player, pos Position) {
	var setupPositions []setupPosition
	switch pos {
	case Position{0, 0}:
		setupPositions = []setupPosition{
			{Position{1, 0}, Position{2, 0}},
			{Position{2, 0}, Position{1, 0}},
			{Position{0, 1}, Position{0, 2}},
			{Position{0, 2}, Position{0, 1}},
			{Position{1, 1}, Position{2, 2}},
			{Position{2, 2}, Position{1, 1}},
		}
	case Position{0, 1}:
		setupPositions = []setupPosition{
			{Position{0, 0}, Position{0, 2}},
			{Position{0, 2}, Position{0, 0}},
			{Position{1, 1}, Position{2, 1}},
			{Position{2, 1}, Position{1, 1}},
		}
	case Position{0, 2}:
		setupPositions = []setupPosition{
			{Position{0, 1}, Position{0, 0}},
			{Position{1, 1}, Position{2, 0}},
			{Position{1, 2}, Position{2, 2}},
			{Position{0, 0}, Position{0, 1}},
			{Position{2, 0}, Position{1, 1}},
			{Position{2, 2}, Position{1, 2}},
		}
	case Position{1, 0}:
		setupPositions = []setupPosition{
			{Position{0, 0}, Position{2, 0}},
			{Position{2, 0}, Position{0, 0}},
			{Position{1, 1}, Position{1, 2}},
			{Position{1, 2}, Position{1, 1}},
		}
	case Position{1, 1}:
		setupPositions = []setupPosition{
			{Position{0, 0}, Position{2, 2}},
			{Position{2, 2}, Position{0, 0}},
			{Position{2, 0}, Position{0, 2}},
			{Position{0, 2}, Position{2, 0}},
			{Position{1, 0}, Position{1, 2}},
			{Position{1, 2}, Position{1, 0}},
			{Position{0, 1}, Position{2, 1}},
			{Position{2, 1}, Position{0, 1}},
		}
	case Position{1, 2}:
		setupPositions = []setupPosition{
			{Position{0, 2}, Position{2, 2}},
			{Position{2, 2}, Position{0, 2}},
			{Position{1, 0}, Position{1, 1}},
			{Position{1, 1}, Position{1, 0}},
		}
	case Position{2, 0}:
		setupPositions = []setupPosition{
			{Position{1, 1}, Position{0, 2}},
			{Position{0, 2}, Position{1, 1}},
			{Position{0, 0}, Position{1, 0}},
			{Position{1, 0}, Position{0, 0}},
			{Position{2, 1}, Position{2, 2}},
			{Position{2, 2}, Position{2, 1}},
		}
	case Position{2, 1}:
		setupPositions = []setupPosition{
			{Position{2, 0}, Position{2, 2}},
			{Position{2, 2}, Position{2, 0}},
			{Position{1, 1}, Position{0, 1}},
			{Position{0, 1}, Position{1, 1}},
		}
	case Position{2, 2}:
		setupPositions = []setupPosition{
			{Position{2, 1}, Position{2, 0}},
			{Position{2, 0}, Position{2, 1}},
			{Position{0, 2}, Position{1, 2}},
			{Position{1, 2}, Position{0, 2}},
			{Position{1, 1}, Position{0, 0}},
			{Position{0, 0}, Position{1, 1}},
		}
	}
	for _, sP := range setupPositions {
		m.setups[sP.setupPos] = append(m.setups[sP.setupPos],
			winningPlay{p, sP.winningPos})
	}
}
