package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/kmdupr33/tictacgo/game"
)

var compPlayerNum = flag.Int("computerPlayer", 0, "specify which player is the computer. possible values: 1 or 2")

func main() {
	flag.Parse()
	f, err := os.Create("./log.txt")
	if err != nil {
		fmt.Println("Couldnt create log file")
	}
	log.SetOutput(f)

	if *compPlayerNum > 2 || *compPlayerNum < 0 {
		printUsage()
		return
	}

	g := game.New(*compPlayerNum)

	printNewGameMessage()

	var posToPlay game.Position

	for !g.IsWon() && !g.IsCatsGame() {

		fmt.Println(g)

		if cp := g.CurrentPlayer(); !cp.IsComputer() {
			log.Println("Non-computer player. Prompting user")
			reader := bufio.NewReader(os.Stdin)

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

			posToPlay = game.Position{X: x, Y: y}

		} else {
			log.Println("Computer player. Getting position from player")
			posToPlay, err = cp.NextMove()
			//This only happens if we're calling NextMove() on a player that
			//isn't a computer player. If this happens, there's no sensible
			//way to recover.
			if err != nil {
				panic(err)
			}
		}

		err = g.PlayTurn(posToPlay)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Println(g)
}

func printUsage() {
	fmt.Println("tictacgo [-computerPlayer=<1|2>]")
}

func printNewGameMessage() {
	fmt.Println("A new game has started! Type 'help' for instructions on how to play")
}

func printInstructions() {
	fmt.Println("To place a marker (X, or O), specify the desired location of the marker by typing a coordinate value. For example, entering '0,1' will place the marker in the top center square on the game board")
}
