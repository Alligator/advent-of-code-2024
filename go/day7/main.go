package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

//go:embed test.txt
var testInput string

//go:embed input.txt
var input string

type equation struct {
	Target int
	Nums   []int
}

func part1(input string) {
	// parse equations
	eqs := make([]equation, 0)

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		ci := strings.Index(line, ":")
		target := int(Check(strconv.ParseInt(line[:ci], 10, 64)))

		numStrs := strings.Fields(line[ci+1:])
		nums := make([]int, len(numStrs))
		for i, num := range numStrs {
			nums[i] = int(Check(strconv.ParseInt(num, 10, 64)))
		}

		eqs = append(eqs, equation{target, nums})
	}

	testTotal := 0

outer:
	for _, eq := range eqs {
		ops := 1 << (len(eq.Nums) - 1)

		for oi := range ops {
			total := eq.Nums[0]
			for i, num := range eq.Nums[1:] {
				if oi&(1<<i) == 0 {
					// add
					total += num
				} else {
					// mul
					total *= num
				}
			}

			if total == eq.Target {
				testTotal += eq.Target
				continue outer
			}
		}
	}

	fmt.Println("part 1:", testTotal)
}

func part2(input string) {
	// parse equations
	eqs := make([]equation, 0)

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		ci := strings.Index(line, ":")
		target := int(Check(strconv.ParseInt(line[:ci], 10, 64)))

		numStrs := strings.Fields(line[ci+1:])
		nums := make([]int, len(numStrs))
		for i, num := range numStrs {
			nums[i] = int(Check(strconv.ParseInt(num, 10, 64)))
		}

		eqs = append(eqs, equation{target, nums})
	}

	var wg sync.WaitGroup
	var sumTotal atomic.Int64

	checkEq := func(eq equation) {
		defer wg.Done()
		ops := int(math.Pow(float64(4), float64(len(eq.Nums))))

	ops:
		for oi := range ops {
			total := eq.Nums[0]
			for i, num := range eq.Nums[1:] {
				opBits := 3 << (i * 2)
				op := oi & opBits >> (i * 2)
				switch op {
				case 0:
					total += num
				case 1:
					total *= num
				case 2:
					// lazyyy
					total = int(Check(strconv.ParseInt(strconv.Itoa(total)+strconv.Itoa(num), 10, 64)))
				default:
					continue ops
				}
			}

			if total == eq.Target {
				sumTotal.Add(int64(eq.Target))
				return
			}
		}
	}

	for _, eq := range eqs {
		wg.Add(1)
		go checkEq(eq)
	}

	wg.Wait()

	fmt.Println("part 2:", sumTotal.Load())
}

func main() {
	part1(input)
	part2(input)
}
