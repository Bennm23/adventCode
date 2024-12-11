package main

import (
	"advent/lib"
	"advent/lib/maths"
	"advent/lib/structures"
)

func main() {
	lib.RunAndScore("Part 1", p1)//Score = 6346871685398. Total Time 3596 us
	lib.RunAndScore("Part 2", p2)//Score = 6373055193464. Total Time 359928 us
}

type Block struct {
	id        int
	block_len int
	space_len int
}

func getBlocks() ([]int, []Block) {
	line := lib.ReadFile("day9.txt")[0]
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
	return block_list, blocks
}
func p1() int {
	sum := 0

	blockList, _ := getBlocks()

	front := 0
	back := len(blockList) - 1

	for back != front {
		for blockList[front] != -1 {
			front++
		}
		if front > back {
			break
		}
		blockList[front] = blockList[back]
		blockList[back] = -1
		back--
	}

	ix := 0
	for blockList[ix] != -1 {
		sum += ix * blockList[ix]
		ix++
	}
	return sum
}

func p2() int {
	sum := 0

	_, blocks := getBlocks()
	var moved structures.Set[int]

	back := len(blocks) - 1
	for back != 0 {

		if moved.Contains(blocks[back].id) {
			back--
			continue
		}

		for i := 0; i < back; i++ {
			space_gap := blocks[i].space_len - blocks[back].block_len
			if space_gap < 0 {
				continue
			}

			if back-1 == i {
				blocks[back].space_len += blocks[back-1].space_len
			} else {
				blocks[back-1].space_len += blocks[back].block_len + blocks[back].space_len
				blocks[back].space_len = space_gap
			}
			blocks[i].space_len = 0

			toMove := blocks[back]
			blocks = append(blocks[:back], blocks[back+1:]...)
			blocks = append(blocks[:i+1], append([]Block{toMove}, blocks[i+1:]...)...)

			back += 1

			moved.Add(toMove.id)
			break
		}
		back--
	}

	ix := 0
	for _, b := range blocks {

		for i := 0; i < b.block_len; i++ {
			sum += ix * b.id
			ix++
		}
		ix += b.space_len
	}
	return sum
}