package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var testInput string = `.....0.
..4321.
..5..2.
..6543.
..7..4.
..8765.
..9....
`

//go:embed input.txt
var input string

func find(input string, allowLoops bool) {
	w := strings.Index(input, "\n")
	h := strings.Count(input, "\n")
	topo := strings.ReplaceAll(input, "\n", "")

	candidates := make([]int, 0)
	for i, l := range topo {
		if l == '0' {
			candidates = append(candidates, i)
		}
	}

	dirs := [][2]int{
		{0, -1}, // up
		{0, 1},  // down
		{-1, 0}, // left
		{1, 0},  // right
	}

	score := 0
	for _, i := range candidates {
		front := []int{i}

		visited := make(map[int]bool)
		visited[i] = true

		numReachable := 0

		for len(front) > 0 {
			f := front[0]
			level := topo[f]

			x := f % w
			y := f / w
			visited[f] = true

			front = front[1:]

			for _, dir := range dirs {
				dx := x + dir[0]
				dy := y + dir[1]
				index := dx + dy*w

				if !allowLoops {
					if _, present := visited[index]; present {
						// already visited
						continue
					}
				}

				if dx >= 0 && dx < w && dy >= 0 && dy < h && topo[index] == level+1 {
					if topo[index] == byte('9') {
						// this is a valid trailhead
						numReachable += 1
						visited[index] = true
					} else {
						front = append(front, dx+dy*w)
					}
				}
			}
		}
		score += numReachable
	}

	fmt.Println(score)
}

func main() {
	fmt.Print("part 1: ")
	find(input, false)

	fmt.Print("part 2: ")
	find(input, true)
}
