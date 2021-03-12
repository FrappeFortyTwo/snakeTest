package main

import "math/rand"

type gameState struct {

	// a 2d array of signed integers
	board [][]uint
	// the current round (each move of the snake constitutes a round)
	round uint

	// score (number of food units eaten)
	score uint

	// food location
	xFood uint
	yFood uint

	// last input
	input string
}

func createArray(height uint, width uint) [][]uint {

	tmp := make([][]uint, height)
	for i := range tmp {
		tmp[i] = make([]uint, width)
	}

	return tmp
}

func newGameState() *gameState {

	return &gameState{
		board: createArray(5, 10),
		round: 0,
		score: 0,
		xFood: uint(rand.Intn(int(height))),
		yFood: uint(rand.Intn(int(width))),
		input: "",
	}
}
