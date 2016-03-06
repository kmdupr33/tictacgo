package game

import "math/rand"

//ComputerPlayerBrain determines which moves the computer player should make
type ComputerPlayerBrain interface {
	getComputerPlayerMove() Position
}

//Randomly chooses where the computer player should play
type randomComputerPlayerBrain struct {
	board Board
}

func (r *randomComputerPlayerBrain) getComputerPlayerMove() Position {

	x := randPos()
	y := randPos()

	pos := Position{x, y}
	sp := r.board.SpaceAt(pos)

	if sp.IsEmpty() {
		return pos
	}

	return r.getComputerPlayerMove()
}

func randPos() int {
	return rand.Intn(3)
}
