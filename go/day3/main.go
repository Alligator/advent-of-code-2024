package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func part1(src string) {
	re := Check(regexp.Compile("mul\\((\\d+),(\\d+)\\)"))

	matches := re.FindAllStringSubmatch(src, -1)
	total := 0
	for _, match := range matches {
		a := Check(strconv.ParseInt(match[1], 10, 64))
		b := Check(strconv.ParseInt(match[2], 10, 64))
		total += int(a * b)
	}
	fmt.Println("part 1:", total)
}

func part2(src string) {
	mulPattern := "(mul\\((\\d+),(\\d+)\\))"
	doPattern := "(do\\(\\))"
	dontPattern := "(don't\\(\\))"
	re := regexp.MustCompile(mulPattern + "|" + doPattern + "|" + dontPattern)

	matches := re.FindAllStringSubmatch(src, -1)
	total := 0
	mulEnabled := true
	for _, match := range matches {
		if mulEnabled && strings.HasPrefix(match[0], "mul") {
			a := Check(strconv.ParseInt(match[2], 10, 64))
			b := Check(strconv.ParseInt(match[3], 10, 64))
			total += int(a * b)
		} else if match[0] == "do()" {
			mulEnabled = true
		} else if match[0] == "don't()" {
			mulEnabled = false
		}
	}
	fmt.Println("part 2:", total)
}

func main() {
	lines := ReadLines("input.txt")
	src := strings.Join(lines, "")
	part1(src)
	part2(src)
}
