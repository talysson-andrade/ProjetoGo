package main

import (
	"fmt"
	"math"
)

const (
	LENGTH = 15
	WIDTH  = 15
	raio   = 5
)

func main() {
	var screen [LENGTH][WIDTH]rune
	p1 := [2]int{1, 1}
	p2 := [2]int{14, 1}
	p3 := [2]int{1, 10}
	p4 := [2]int{14, 10}

	connectPoints(&screen, p1, p2)
	connectPoints(&screen, p1, p3)
	connectPoints(&screen, p4, p3)
	connectPoints(&screen, p4, p2)

	printScreen(screen)
}

func verifyPoint(point *[2]int) bool {
	if point[0] < LENGTH && point[1] < WIDTH && point[0] >= 0 && point[1] >= 0 {
		return true
	}
	return false
}

func connectPoints(screen *[LENGTH][WIDTH]rune, p1 [2]int, p2 [2]int) { // Bresenham's Line Drawing Algorithm
	dx := int(math.Abs(float64(p1[0] - p2[0])))
	dy := int(math.Abs(float64(p1[1] - p2[1])))
	sx := 1
	if p2[0] < p1[0] {
		sx = -1
	}
	sy := 1
	if p2[1] < p1[1] {
		sy = -1
	}
	err := dx - dy

	for {
		if verifyPoint(&p1) {
			screen[p1[0]][p1[1]] = '*'
		}
		if p1[0] == p2[0] && p1[1] == p2[1] {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			p1[0] += sx
		}
		if e2 < dx {
			err += dx
			p1[1] += sy
		}
	}
}

func printScreen(screen [LENGTH][WIDTH]rune) {
	for i := 0; i < LENGTH; i++ {
		for j := 0; j < WIDTH; j++ {
			if screen[i][j] != 0 {
				fmt.Printf("%c ", screen[i][j])
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
