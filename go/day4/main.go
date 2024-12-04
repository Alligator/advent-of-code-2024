package main

import (
	_ "embed"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
)

//go:embed test.txt
var testInput string

//go:embed input.txt
var input string

func check(cells []string, x int, y int, w int, search string) bool {
	index := y*w + x
	return cells[index] == search
}

func part1(input string) {
	w := strings.Index(input, "\n")
	h := strings.Count(input, "\n")
	cells := strings.Split(strings.ReplaceAll(input, "\n", ""), "")

	var wg sync.WaitGroup
	var count atomic.Int64

	for i, cell := range cells {
		if cell == "X" {
			wg.Add(1)
			go func() {
				x := i % w
				y := i / w

				// check up
				if y > 2 &&
					check(cells, x, y-1, w, "M") &&
					check(cells, x, y-2, w, "A") &&
					check(cells, x, y-3, w, "S") {
					count.Add(1)
				}

				// check up right
				if x < w-3 &&
					y > 2 &&
					check(cells, x+1, y-1, w, "M") &&
					check(cells, x+2, y-2, w, "A") &&
					check(cells, x+3, y-3, w, "S") {
					count.Add(1)
				}

				// check right
				if x < w-3 &&
					check(cells, x+1, y, w, "M") &&
					check(cells, x+2, y, w, "A") &&
					check(cells, x+3, y, w, "S") {
					count.Add(1)
				}

				// check down right
				if x < w-3 &&
					y < h-3 &&
					check(cells, x+1, y+1, w, "M") &&
					check(cells, x+2, y+2, w, "A") &&
					check(cells, x+3, y+3, w, "S") {
					count.Add(1)
				}

				// check down
				if y < h-3 &&
					check(cells, x, y+1, w, "M") &&
					check(cells, x, y+2, w, "A") &&
					check(cells, x, y+3, w, "S") {
					count.Add(1)
				}

				// check down left
				if x > 2 &&
					y < h-3 &&
					check(cells, x-1, y+1, w, "M") &&
					check(cells, x-2, y+2, w, "A") &&
					check(cells, x-3, y+3, w, "S") {
					count.Add(1)
				}

				// check left
				if x > 2 &&
					check(cells, x-1, y, w, "M") &&
					check(cells, x-2, y, w, "A") &&
					check(cells, x-3, y, w, "S") {
					count.Add(1)
				}

				// check up left
				if x > 2 &&
					y > 2 &&
					check(cells, x-1, y-1, w, "M") &&
					check(cells, x-2, y-2, w, "A") &&
					check(cells, x-3, y-3, w, "S") {
					count.Add(1)
				}

				wg.Done()
			}()
		}
	}

	wg.Wait()

	fmt.Println("part 1:", count.Load())
}

func part2(input string) {
	w := strings.Index(input, "\n")
	h := strings.Count(input, "\n")
	cells := strings.Split(strings.ReplaceAll(input, "\n", ""), "")

	var wg sync.WaitGroup
	var count atomic.Int64

	for i, cell := range cells {
		if cell == "A" {
			wg.Add(1)
			go func() {
				x := i % w
				y := i / w

				if x < 1 || x > w-2 || y < 1 || y > h-2 {
					wg.Done()
					return
				}

				masCount := 0

				// M
				//  A
				//   S
				if check(cells, x-1, y-1, w, "M") && check(cells, x+1, y+1, w, "S") {
					masCount++
				}

				// S
				//  A
				//   M
				if check(cells, x-1, y-1, w, "S") && check(cells, x+1, y+1, w, "M") {
					masCount++
				}

				//   M
				//  A
				// S
				if check(cells, x+1, y-1, w, "M") && check(cells, x-1, y+1, w, "S") {
					masCount++
				}

				//   S
				//  A
				// M
				if check(cells, x+1, y-1, w, "S") && check(cells, x-1, y+1, w, "M") {
					masCount++
				}

				if masCount >= 2 {
					count.Add(1)
				}

				wg.Done()
			}()
		}
	}

	wg.Wait()

	fmt.Println("part 2:", count.Load())
}

func main() {
	part1(input)
	part2(input)
}
