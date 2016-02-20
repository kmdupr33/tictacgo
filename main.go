package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kmdupr33/tictacgo/game"
)

func main() {

	g := game.New()

	printNewGameMessage()

	for !g.IsWon() || !g.IsCatsGame() {

		reader := bufio.NewReader(os.Stdin)
		fmt.Println(g.String())

		text, err := reader.ReadString('\n')
		if err != nil {
			printInstructions()
			continue
		}
		text = strings.TrimSuffix(text, "\n")
		if text == "help" {
			printInstructions()
			continue
		}

		cords := strings.Split(text, ",")
		x, err := strconv.Atoi(cords[0])
		y, err := strconv.Atoi(cords[1])
		if err != nil {
			fmt.Println("You've specified an invalid marker position")
			printInstructions()
			continue
		}
		err = g.PlayTurn(game.Position{x, y})
		if err != nil {
			fmt.Println(err.Error())
		}

	}
	fmt.Printf("%v's game!", g.Winner())
}

func printNewGameMessage() {
	fmt.Println("A new game has started! Type 'help' for instructions on how to play")
}

func printInstructions() {
	fmt.Println("To place a marker (X, or O), specify the desired location of the marker by typing a coordinate value. For example, entering '0,1' will place the marker in the top center square on the game board")
}
