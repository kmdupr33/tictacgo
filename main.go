package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kmdupr33/tictacgo/game"
)

func main() {

	g := new(game.Game)
	fmt.Println("A new game has started! Type 'help' for instructions on how to play")
	for !g.IsGameWon() || !g.IsCatsGame() {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("%v's Turn: ", g.CurrentPlayer())
		text, _ := reader.ReadString('\n')
		if text == "help" {
			printInstructions()
		}
		fmt.Println(text)
	}
	fmt.Printf("%v's game!", g.Winner())
}

func printInstructions() {

}
