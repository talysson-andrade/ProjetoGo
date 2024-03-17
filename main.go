package main

import (
	"fmt"
)

const (
	LENGTH = 50
	WIDTH  = 50
	raio   = 5
)

func main() {
	var screen [LENGTH][WIDTH]rune
	p1 := [2]float64{6, 3}
	p2 := [2]float64{7, 6}
	center := [2]float64{(LENGTH / 2) - 1, (WIDTH / 2) - 1}
	if verifyPoint(&p1) {
		screen[int(p1[0])][int(p1[1])] = '*'
	}
	screen[int(center[0])][int(center[1])] = '#'

	connectPoints(&screen, &p1, &p2)
	printScreen(screen)
}

func verifyPoint(point *[2]float64) bool {
	if point[0] < LENGTH && point[1] < WIDTH {
		return true
	}
	return false
}

func connectPoints(screen *[LENGTH][WIDTH]rune, p1 *[2]float64, p2 *[2]float64) {
	m := (p2[1] - p1[1]) / (p2[0] - p1[0])
	n := p1[1] - (m * p1[0])

	for i := 0.0; i < LENGTH; i++ {
		for j := 0.0; j < WIDTH; j++ {
			if i == m*j+n {
				screen[int(j)][int(i)] = '*'
			}
		}

	}
}

func printScreen(screen [LENGTH][WIDTH]rune) {
	for i := 0; i < LENGTH; i++ {
		for j := 0; j < WIDTH; j++ {
			if screen[j][i] != 0 {
				fmt.Printf("%c ", screen[j][i])
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
