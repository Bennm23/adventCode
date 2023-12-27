package main

import (
	"advent/lib"
	"advent/lib/avstrings"
	"advent/lib/structures"
	"fmt"
	"math"
)

type Point [3]int

type Block struct {
	points     []Point
	identifier int
	blockedBy, supports structures.Set[int]
}

var ID = 0

func NewBlock(line string) *Block {
	nums := avstrings.SplitTextToInts(line)

	var left, right Point = Point(nums[:3]), Point(nums[3:])

	if right[2] > Z_MAX {
		Z_MAX = right[2]
	}

	points := make([]Point, 0)
	for x := left[0]; x <= right[0]; x++ {
		for y := left[1]; y <= right[1]; y++ {
			for z := left[2]; z <= right[2]; z++ {
				points = append(points, Point{x, y, z})
			}
		}
	}

	ID++
	return &Block{
		identifier: ID,
		points:     points,
		blockedBy:   structures.Set[int]{},
	}
}

var Z_MAX int = math.MinInt

func main() {
	lib.RunAndPrintDurationMillis(func() { solve() })//2200 - 2500
}

func solve() {

	lines := lib.ReadFile("day22.txt")

	blocks := map[int]*Block{}

	for _, line := range lines {
		b := NewBlock(line)
		blocks[b.identifier] = b
	}

	//Indexed by the z level
	zPlane := structures.Create3DArray[int]([]int{Z_MAX + 1, 10, 10})
	zMap := map[int][]*Block{}

	for _, block := range blocks {
		minZ := math.MaxInt
		for _, point := range block.points {
			zPlane[point[2]][point[0]][point[1]] = block.identifier
			if point[2] < minZ {
				minZ = point[2]
				
			}
		}
		if val, found := zMap[minZ]; found {
			zMap[minZ] = append(val, block)
		} else {
			zMap[minZ] = []*Block{block}
		}
	}

	//Now we have to drop the blocks down
	for plane := 1; plane < len(zPlane); plane++ {

		blocksOnPlane, found := zMap[plane]
		if !found {
			continue
		}

		//For any block on this plane, shift it down until there is a block beneath it
		for _, block := range blocksOnPlane {

			blocked := false
			for !blocked {

				for _, point := range block.points {
					x, y, z := point[0], point[1], point[2]

					//If any point is at ground level, we are blocked
					if z == 0 {
						blocked = true
						break
					}
					//If any point has another identifier beneath it, we are blocked
					beneath := zPlane[z-1][x][y]
					if beneath != 0 && beneath != block.identifier {
						//Then this is an invalid move and we are blocked
						blocked = true
						block.blockedBy.Add(beneath)
						blocks[beneath].supports.Add(block.identifier)
					}
				}

				if !blocked {
					newPoints := []Point{}
					//If we aren't blocked, then shift all points down
					for _, point := range block.points {
						x, y, z := point[0], point[1], point[2]

						zPlane[z][x][y] = 0
						zPlane[z-1][x][y] = block.identifier
						newPoints = append(newPoints, Point{x, y, z - 1})
					}
					block.points = newPoints
				}
			}
		}
	}

	necessaryBlocks := structures.Set[int]{}

	for _, block := range blocks {
		if len(block.blockedBy) == 1 {
			necessaryBlocks.AddAll(block.blockedBy...)
		}
	}

	fmt.Println("Part 1 = ", len(blocks)-len(necessaryBlocks)) //430

	p2 := 0

	//For each block that causes others to fall, walk up the support ladder, if any block is supported only by blocks in this chain reaction, it joins the chain reaction
	for _, blockId := range necessaryBlocks {
		base := blocks[blockId]

		fallers := structures.Set[int]{base.identifier}
		countFallers(base, &blocks, &fallers)
		p2 += len(fallers) - 1 //Not including itself
	}
	fmt.Println("Part 2 = ", p2) //60558
}

func countFallers(block *Block, blocks *map[int]*Block, fallers *structures.Set[int]) {
	newFallers := []*Block{}
	//For each support, if is blocked only by the bricks that would fall, then it will fall as well
	for _, supportId := range block.supports {

		support := (*blocks)[supportId]

		if !fallers.ContainsAll(support.blockedBy...) {
			continue
		}

		newFallers = append(newFallers, support)
		fallers.Add(supportId)
	}

	for _, b := range newFallers {
		countFallers(b, blocks, fallers)
	}
}
