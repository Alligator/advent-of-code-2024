package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

type numeric interface {
	int | int64
}

func abs[T numeric](value T) T {
	if value < 0 {
		return -value
	}
	return value
}

func readLines(filename string) [][]int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	lines := make([][]int, 0)

	for s.Scan() {
		fields := strings.Fields(s.Text())
		line := make([]int, 0)
		for _, field := range fields {
			n := check(strconv.ParseInt(field, 10, 64))
			line = append(line, int(n))
		}
		lines = append(lines, line)
	}

	return lines
}

func isSafe(line []int) bool {
	prev := 0
	increasing := true
	decreasing := true
	safe := true

	for i, n := range line {
		if i == 0 {
			prev = n
			continue
		}

		diff := n - prev
		if abs(diff) < 1 || abs(diff) > 3 {
			safe = false
			break
		}
		increasing = increasing && diff > 0
		decreasing = decreasing && diff < 0

		prev = n
	}

	safe = safe && (increasing || decreasing)
	return safe
}

func main() {
	lines := readLines("input.txt")

	total := 0
	for _, line := range lines {
		if isSafe(line) {
			total++
		}
	}

	fmt.Println("part 1:", total)

	// brain dev part 2 solution
	total = 0

outer:
	for _, line := range lines {
		if isSafe(line) {
			total++
			continue
		}

		for i := 0; i < len(line); i++ {
			clone := make([]int, len(line))
			copy(clone, line)

			modifiedLine := append(clone[:i], clone[i+1:]...)

			if isSafe(modifiedLine) {
				total++
				continue outer
			}
		}
	}

	fmt.Println("part 2:", total)
}
