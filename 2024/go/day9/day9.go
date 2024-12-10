package main

import (
	"advent/lib"
	"advent/lib/maths"
	"advent/lib/structures"
	"strconv"
	"strings"
)

func main() {
	// lib.RunAndScore("Part 1", p1)
	lib.RunAndScore("Part 2", p2)

}

type Block struct {
	id        int
	block_len int
	space_len int
}

func (b Block) printBlock() {
	block_str := strconv.Itoa(b.id)

	block_str = strings.Repeat(block_str, b.block_len)
	space_str := strings.Repeat(".", b.space_len)

	lib.Lognl(block_str, space_str)
}

func p1() int {
	sum := 0

	line := lib.ReadFile("day9.txt")[0]

	lib.Log("Line = ", line)

	id := 0

	var blocks []Block

	block_list := []int{}

	for i := 0; i < len(line); i += 2 {

		block_len := maths.ToInt(string(line[i]))
		var space_len int
		if i+1 == len(line) {
			space_len = 0
		} else {
			space_len = maths.ToInt(string(line[i+1]))
		}
		lib.Log("Block ID = ", id, " Block Len = ", block_len, " Space Len = ", space_len)

		block := Block{id, block_len, space_len}

		blocks = append(blocks, block)

		for i := 0; i < block_len; i++ {
			block_list = append(block_list, id)
		}
		for i := 0; i < space_len; i++ {
			block_list = append(block_list, -1)
		}
		id += 1
	}

	for _, block := range blocks {
		lib.Lognl("Block = ")
		block.printBlock()
		lib.Log()
	}

	lib.Log("BlockArr = ", block_list)

	front := 0
	back := len(block_list) - 1

	for back != front {

		for block_list[front] != -1 {
			front++
		}

		// lib.Log("Back = ", back, " Front = ", front)
		if front > back {
			break
		}

		block_list[front] = block_list[back]
		block_list[back] = -1
		// lib.Log("BlockArr = ", block_list)
		back--
	}

	ix := 0

	for block_list[ix] != -1 {
		sum += ix * block_list[ix]
		ix++
	}

	return sum
}
const DEBUG = false

func p2() int {
	sum := 0

	line := lib.ReadFile("day9.txt")[0]

	id := 0
	var blocks []Block

	for i := 0; i < len(line); i += 2 {

		block_len := maths.ToInt(string(line[i]))
		var space_len int
		if i+1 == len(line) {
			space_len = 0
		} else {
			space_len = maths.ToInt(string(line[i+1]))
		}
		lib.Log("Block ID = ", id, " Block Len = ", block_len, " Space Len = ", space_len)

		block := Block{id, block_len, space_len}
		blocks = append(blocks, block)
		id += 1
	}

	var moved structures.Set[int]

	for bix, b := range blocks {
		// b.printBlock()
		// fmt.Println()
		lib.Log("BIX = ", bix, " Block = ", b)
	}
	back := len(blocks) - 1
	nicePrint(blocks)
	lib.Log(blocks)
	for back != 0 {

		if moved.Contains(blocks[back].id) {
			// lib.Log("Already Moved ", blocks[back].id)
			back--
			continue
		}

        if DEBUG {
            lib.Log("----------")
            lib.Log("Back = ", back)

            lib.Log("BLOCKS = ", blocks)
            lib.Log("EVALUATING ID = ", blocks[back].id, " Space len = ", blocks[back].space_len)
        }
		for i := 0; i < back; i++ {
			space_gap := blocks[i].space_len - blocks[back].block_len
			if space_gap >= 0 {

				if back-1 == i {

					// blocks[back].space_len = blocks[back-1].space_len - blocks[back].block_len + blocks[back].space_len
                    //orig space change = orig.space_len - moved.block_len

					blocks[back].space_len += blocks[back-1].space_len

                    //orig space = 4, new len = 3, new space = 3
                    //(orig_space - moved_len) = 1 is leftover orignal space
                    //(moved_len + moved_space) =  is replaced + original space
                    //total = orig_space - new_len + new_len + new_space
					// blocks[back-1].space_len = 0

				} else {
					blocks[back-1].space_len += blocks[back].block_len + blocks[back].space_len
					blocks[back].space_len = space_gap
				}
				blocks[i].space_len = 0

                if DEBUG {
                    lib.Log("Swapping Block Id = ", blocks[back].id, " With ", blocks[i].id)
                    lib.Log("Space Gap = ", space_gap)
                }
				toMove := blocks[back]
				blocks = append(blocks[:back], blocks[back+1:]...)
				blocks = append(blocks[:i+1], append([]Block{toMove}, blocks[i+1:]...)...)

				back += 1

				moved.Add(toMove.id)
				break
			}
		}
        if DEBUG {
            nicePrint(blocks)
        }
		back--
	}

	ix := 0
	for _, b := range blocks {
		// b.printBlock()
		// fmt.Println()

		for i := 0; i < b.block_len; i++ {
			// fmt.Println("IX = ", ix, " B ID = ", b.id, " Sum = ", (ix * b.id))
			sum += ix * b.id
			ix++
		}
		// fmt.Println("Sum so far = ", sum)
		ix += b.space_len
	}
	nicePrint(blocks)

 //              6373055193464
	return sum //6373055544417 TOO HIGH 6373046941244 TOO LOW
}

func nicePrint(blocks []Block) {
	for _, b := range blocks {
		b.printBlock()
	}
	lib.Log()
}
