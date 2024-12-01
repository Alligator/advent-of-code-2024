package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)

	left := make([]int, 0)
	right := make([]int, 0)

	for s.Scan() {
		line := s.Text()
		sp := strings.Fields(line)
		l := check(strconv.ParseInt(sp[0], 10, 64))
		r := check(strconv.ParseInt(sp[1], 10, 64))
		left = append(left, int(l))
		right = append(right, int(r))
	}

	// part 1
	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		l := left[i]
		r := right[i]
		diff := l - r
		if diff < 0 {
			sum += -diff
		} else {
			sum += diff
		}
	}
	fmt.Println("part 1:", sum)

	// part 2
	cache := make(map[int]int)

	simScore := 0
	for _, l := range left {
		if count, present := cache[l]; present {
			simScore += l * count
			continue
		}

		count := 0
		for _, r := range right {
			if r == l {
				count++
			}
		}

		cache[l] = count
		simScore += l * count
	}

	fmt.Println("part 2:", simScore)
}
