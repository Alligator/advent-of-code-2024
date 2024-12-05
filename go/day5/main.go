package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed test.txt
var testInput string

//go:embed input.txt
var input string

type rule struct {
	before string
	after  string
}

func part1(input string) {
	sections := strings.Split(input, "\n\n")

	rules := make([]rule, 0)

	for _, r := range strings.Split(sections[0], "\n") {
		// all 2 digit
		before := r[:2]
		after := r[3:]
		rules = append(rules, rule{before, after})
	}

	total := 0
	for _, update := range strings.Split(sections[1], "\n") {
		if update == "" {
			continue
		}

		// parse pages
		pages := make(map[string]int)
		middleNum := int64(0)
		sp := strings.Split(update, ",")
		for i, page := range sp {
			pages[page] = i
			if i == len(sp)/2 {
				middleNum = Check(strconv.ParseInt(page, 10, 64))
			}
		}

		// check all rules
		valid := true
		for _, rule := range rules {
			indexBefore, beforePresent := pages[rule.before]
			if !beforePresent {
				continue
			}

			indexAfter, afterPresent := pages[rule.after]
			if !afterPresent {
				continue
			}

			if indexBefore >= indexAfter {
				valid = false
				break
			}
		}

		if valid {
			total += int(middleNum)
		}
	}

	fmt.Println("part 1:", total)
}

func part2(input string) {
	sections := strings.Split(input, "\n\n")

	rules := make([]rule, 0)

	for _, r := range strings.Split(sections[0], "\n") {
		// all 2 digit
		before := r[:2]
		after := r[3:]
		rules = append(rules, rule{before, after})
	}

	total := int64(0)
	for _, update := range strings.Split(sections[1], "\n") {
		if update == "" {
			continue
		}

		// parse pages
		pages := make(map[string]int)
		sp := strings.Split(update, ",")
		for i, page := range sp {
			pages[page] = i
		}

		// check all rules
		valid := false
		wasInvalid := false
		for !valid {
			valid = true
			for _, rule := range rules {
				indexBefore, beforePresent := pages[rule.before]
				if !beforePresent {
					continue
				}

				indexAfter, afterPresent := pages[rule.after]
				if !afterPresent {
					continue
				}

				if indexBefore >= indexAfter {
					valid = false
					wasInvalid = true
					// swap and retry
					tmp := pages[rule.before]
					pages[rule.before] = pages[rule.after]
					pages[rule.after] = tmp
					break
				}
			}
		}

		if wasInvalid {
			for k, v := range pages {
				if v == len(sp)/2 {
					total += Check(strconv.ParseInt(k, 10, 64))
				}
			}
		}
	}

	fmt.Println("part 2:", total)
}

func main() {
	part1(input)
	part2(input)
}
