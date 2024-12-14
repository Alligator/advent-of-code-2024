package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed test.txt
var testInput string

//go:embed input.txt
var input string

type antenna struct {
	X, Y int
}

func part1(input string) {
	antennaByType := make(map[rune][]antenna)

	w := strings.Index(input, "\n")
	h := strings.Count(input, "\n")

	for y, line := range strings.Split(input, "\n") {
		for x, c := range line {
			if c == '.' {
				continue
			}

			if antennas, present := antennaByType[c]; present {
				antennaByType[c] = append(antennas, antenna{x, y})
			} else {
				arr := make([]antenna, 1)
				arr[0] = antenna{x, y}
				antennaByType[c] = arr
			}
		}
	}

	antiNodes := make(map[antenna]bool)

	addAntiNode := func(a, b antenna) {
		antiNode := antenna{
			X: a.X*2 - b.X,
			Y: a.Y*2 - b.Y,
		}
		if antiNode.X < 0 || antiNode.X >= w || antiNode.Y < 0 || antiNode.Y >= h {
			return
		}
		antiNodes[antiNode] = true
	}

	for _, antennas := range antennaByType {
		for i, ant := range antennas {
			for _, otherAnt := range antennas[i+1:] {
				addAntiNode(ant, otherAnt)
				addAntiNode(otherAnt, ant)
			}
		}
	}

	fmt.Println("part 1:", len(antiNodes))
}

func part2(input string) {
	antennaByType := make(map[rune][]antenna)

	w := strings.Index(input, "\n")
	h := strings.Count(input, "\n")

	for y, line := range strings.Split(input, "\n") {
		for x, c := range line {
			if c == '.' {
				continue
			}

			if antennas, present := antennaByType[c]; present {
				antennaByType[c] = append(antennas, antenna{x, y})
			} else {
				arr := make([]antenna, 1)
				arr[0] = antenna{x, y}
				antennaByType[c] = arr
			}
		}
	}

	antiNodes := make(map[antenna]bool)

	addAntiNode := func(a, b antenna) {
		xd := a.X - b.X
		yd := a.Y - b.Y

		x := a.X + xd
		y := a.Y + yd

		for x >= 0 && x < w && y >= 0 && y < h {
			antiNode := antenna{x, y}
			antiNodes[antiNode] = true
			x += xd
			y += yd
		}

		x = a.X - xd
		y = a.Y - yd

		for x >= 0 && x < w && y >= 0 && y < h {
			antiNode := antenna{x, y}
			antiNodes[antiNode] = true
			x -= xd
			y -= yd
		}
	}

	for _, antennas := range antennaByType {
		for i, ant := range antennas {
			for _, otherAnt := range antennas[i+1:] {
				addAntiNode(ant, otherAnt)
				addAntiNode(otherAnt, ant)
			}
		}
	}

	fmt.Println("part 2:", len(antiNodes))
}

func main() {
	part1(input)
	part2(input)
}
