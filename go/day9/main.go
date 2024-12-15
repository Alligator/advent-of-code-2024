package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string
var testInput string = "2333133121414131402"

func printBlocks(blocks []int) {
	for _, b := range blocks {
		if b == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(b)
		}
	}
	fmt.Print("\n")
}

func part1(input string) {
	// -1 = free
	blocks := make([]int, 0)

	fileId := 0
	for i, c := range strings.TrimSpace(input) {
		n := StrToInt(string(c))
		f := fileId
		if i%2 == 1 {
			// free block
			f = -1
		} else {
			fileId++
		}
		for range n {
			blocks = append(blocks, f)
		}
	}

	freeBlockIndex := 0
	for i, block := range blocks {
		if block == -1 {
			freeBlockIndex = i
			break
		}
	}

	blockIndexToMove := len(blocks) - 1

	for blockIndexToMove > freeBlockIndex {
		blocks[freeBlockIndex] = blocks[blockIndexToMove]
		blocks[blockIndexToMove] = -1

		for i := blockIndexToMove - 1; i >= 0; i-- {
			b := blocks[i]
			if b != -1 {
				blockIndexToMove = i
				break
			}
		}

		for i, b := range blocks[freeBlockIndex:] {
			if b == -1 {
				freeBlockIndex += i
				break
			}
		}
	}

	checkSum := 0
	for i, b := range blocks {
		if b != -1 {
			checkSum += i * b
		}
	}

	fmt.Println("part 1:", checkSum)
}

type block struct {
	FileId int // -1 = free
	Size   int
	Moved  bool
	Next   *block
	Prev   *block
}

func (b *block) print() {
	blk := b
	for blk != nil {
		for _ = range blk.Size {
			if blk.FileId == -1 {
				fmt.Print(".")
			} else {
				fmt.Print(blk.FileId)
			}
		}
		blk = blk.Next
	}
	fmt.Print("\n")
}

func part2(input string) {
	var head *block
	var tail *block

	fileId := 0
	for i, c := range strings.TrimSpace(input) {
		n := StrToInt(string(c))
		f := fileId
		if i%2 == 1 {
			// free block
			f = -1
		} else {
			fileId++
		}

		blk := block{FileId: f, Size: n, Moved: false}

		if head == nil {
			head = &blk
			tail = &blk
		} else {
			tail.Next = &blk
			blk.Prev = tail
			tail = &blk
		}
	}

	head.Moved = true

	blkToMove := tail

	findFreeBlock := func(start *block) *block {
		blk := start.Next
		for blk != nil {
			if blk == blkToMove {
				return nil
			}

			if blk.FileId == -1 {
				return blk
			}
			blk = blk.Next
		}
		return nil
	}

	findBlockToMove := func(start *block) {
		if start == nil {
			blkToMove = nil
			return
		}

		blk := start.Prev
		for blk != nil {
			if !blk.Moved && blk.FileId != -1 {
				blkToMove = blk
				return
			}
			blk = blk.Prev
		}
		blkToMove = nil
	}

	for blkToMove != nil {
		btmPrev := blkToMove.Prev

		freeBlk := findFreeBlock(head)

		for freeBlk != nil {
			if freeBlk.Size == blkToMove.Size {
				// swap the two blocks
				freeBlk.FileId = blkToMove.FileId
				blkToMove.FileId = -1
				freeBlk.Moved = true

				freeBlk = freeBlk.Next
				findBlockToMove(btmPrev)
				break
			} else if freeBlk.Size > blkToMove.Size {
				// shrink the free block
				newBlk := block{
					FileId: blkToMove.FileId,
					Size:   blkToMove.Size,
					Moved:  true,
					Prev:   freeBlk.Prev,
					Next:   freeBlk,
				}

				freeBlk.Size -= blkToMove.Size
				if freeBlk.Prev != nil {
					freeBlk.Prev.Next = &newBlk
				}
				freeBlk.Prev = &newBlk

				blkToMove.FileId = -1
				findBlockToMove(btmPrev)
				break
			}

			freeBlk = findFreeBlock(freeBlk)
		}
		findBlockToMove(btmPrev)
	}

	checkSum := 0
	blk := head

	i := 0
	for blk != nil {
		for _ = range blk.Size {
			if blk.FileId != -1 {
				checkSum += blk.FileId * i
			}
			i++
		}
		blk = blk.Next
	}

	fmt.Println("part 2:", checkSum)
}

func main() {
	part1(input)
	part2(input)
}
