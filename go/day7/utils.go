package main

import (
	"bufio"
	"os"
)

func Check[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

type Numeric interface {
	int | int64
}

func Abs[T Numeric](value T) T {
	if value < 0 {
		return -value
	}
	return value
}

func ReadLines(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	lines := make([]string, 0)

	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines
}
