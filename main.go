package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
)

// height & width of the board
var height uint
var width uint

func main() {

	// fetch board height & width
	println("\nEnter board height :")
	_, err := fmt.Scanln(&height)
	if err != nil {
		log.Fatalln("Invalid input for board height. Please try again.")
	}

	println("\nEnter board width :")
	_, err = fmt.Scanln(&width)
	if err != nil {
		log.Fatalln("Invalid input for board height. Please try again.")
	}

	// spawn gameState
	g := newGameState()

	// spawn snake
	s := newSnake()

	// render snake ( on start )
	g.board[s.xHead][s.yHead] = 1

	// keep the game running ~ game loop
	for {
		// print gameState
		printState(g, s)

		// Take user input
		takeInput(g, s)

		// make changes to game state

	}
}

func printState(g *gameState, s *snake) {

	// check for respawn
	if s.xHead == g.xFood && s.yHead == g.yFood {

		// respawn food coordinates
		g.xFood = uint(rand.Intn(int(height)))
		g.yFood = uint(rand.Intn(int(width)))

		// increase score
		g.score++
		// increase it's length
		s.length++
	}

	fmt.Printf("\nROUND : %d\tSCORE : %d\tLENGTH : %d\tHEAD : [%d,%d]\n", g.round, g.score, s.length, s.xHead, s.yHead)
	println()

	for i := 0; i < int(height); i++ {
		for j := 0; j < int(width); j++ {

			if uint(i) == g.xFood && uint(j) == g.yFood {
				g.board[i][j] = 5
			}

			fmt.Printf("%d", g.board[i][j])

		}
		println()
	}
}

func takeInput(g *gameState, s *snake) {

	println("\nEnter Input : W (top), S (bottom), A (left), D (right):\n")
	fmt.Scanln(&g.input)
	g.input = strings.ToLower(g.input)

	switch g.input {

	case "w":
		// incease round
		g.round++

		// clear head current position
		g.board[s.xHead][s.yHead] = 0
		// move head up
		s.xHead--

		// if head reaches top edge
		if s.xHead+1 == 0 {
			s.xHead = height - 1
		}

		// render changes
		g.board[s.xHead][s.yHead] = 1

	case "s":
		// incease round
		g.round++

		// clear head current position
		g.board[s.xHead][s.yHead] = 0
		// move head down
		s.xHead++

		// if head reaches top edge
		if s.xHead == height {
			s.xHead = 0
		}

		// render changes
		g.board[s.xHead][s.yHead] = 1

	case "a":
		// incease round
		g.round++

		// clear head current position
		g.board[s.xHead][s.yHead] = 0
		// move head left
		s.yHead--

		// if head reaches top edge
		if s.yHead+1 == 0 {
			s.yHead = width - 1
		}

		// render changes
		g.board[s.xHead][s.yHead] = 1

	case "d":
		// incease round
		g.round++

		// clear head current position
		g.board[s.xHead][s.yHead] = 0
		// move head right
		s.yHead++

		// if head reaches top edge
		if s.yHead == width {
			s.yHead = 0
		}

		// render changes
		g.board[s.xHead][s.yHead] = 1

	default:
		println("Invalid Input. Please try again !:\n")

	}

}
