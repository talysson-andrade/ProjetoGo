package main

import (
	"fmt"
	"math"
)

const (
	LENGTH = 20
	WIDTH  = 20
	raio   = 5
)

func main() {
	var screen [LENGTH][WIDTH]rune

	for theta := 0.0; theta < 360; theta += 0.1 {
		x := raio*math.Cos(theta*math.Pi/180) + (LENGTH / 2)
		y := raio*math.Sin(theta*math.Pi/180) + (WIDTH / 2)
		if int(x) >= 0 && int(x) < LENGTH && int(y) >= 0 && int(y) < WIDTH {
			screen[int(x)][int(y)] = '*'
		}
	}

	printScreen(screen)
}

func printScreen(screen [LENGTH][WIDTH]rune) {
	for i := 0; i < LENGTH; i++ {
		for j := 0; j < WIDTH; j++ {
			if screen[i][j] != 0 {
				fmt.Printf("%c", screen[i][j])
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
