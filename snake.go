package main

import "math/rand"

type snake struct {

	// coordiantes of snake head
	xHead uint
	yHead uint

	// lenght of the snake
	length uint

	// turn position
	turn uint
}

func newSnake() *snake {

	return &snake{
		xHead:  uint(rand.Intn(int(height))),
		yHead:  uint(rand.Intn(int(width))),
		length: 2,
		turn:   0,
	}
}
