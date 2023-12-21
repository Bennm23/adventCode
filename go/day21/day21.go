package main

import (
	"advent/lib"
	"advent/lib/structures"
	"fmt"
	"math"
)

func main() {
	solve()
}

// type Position [2]int
type Coord [2]int
type Position struct {
	coords   Coord
	useCount int
}

var GRID_LEN int = 0
var GRID_WIDTH int = 0

func solve() {

	grid := lib.ReadFileToGrid("sample")

	GRID_LEN = len(grid)
	GRID_WIDTH = len(grid[0])

	rockCount := 0

	startPosition := Position{}
	for i, row := range grid {
		fmt.Println(string(row))
		for j, col := range row {
			if col == 'S' {
				startPosition = Position{Coord{i, j}, 1}
			}
			if col == '#' {
				rockCount++
			}
		}
	}

	fmt.Println("GRID LENGTH = ", GRID_LEN)
	fmt.Println("GRID WIDTH = ", GRID_WIDTH)

	fmt.Println("Start Pos = ", startPosition)
	// exploreQuantumI(&grid, 1, []Position{startPosition})
	// exploreQuantumCount(&grid, 1, &map[Coord]int {startPosition.coords: 1,})
	startPoint := MDP{
		x:          startPosition.coords[0],
		y:          startPosition.coords[1],
		dimensions: structures.Set[Dimension]{{0, 0}},
	}
	explore(&grid, 1, &map[Coord]MDP{startPosition.coords: startPoint})

	fmt.Println("Part 1 = ", P1) //3677

	fmt.Println("Rock Count = ", rockCount)
}

func mapCoordinate(x, y int) (a, b int) {

	a = x % GRID_LEN
	b = y % GRID_WIDTH

	if a < 0 {
		a += GRID_LEN
	}
	if b < 0 {
		b += GRID_WIDTH
	}

	return a, b
}

func mapCoordinateD(x, y int) (a, b int, dimension []int) {
	//Given Some Coordinate X and Y in the infinite space
	//We must map them to the original grid

	// fmt.Printf("Virtual Coordinate (%d, %d)\n", x, y)

	//Ex. Coordinate 14, 20
	// 	  Sample Grid is [0:11]:11
	//	 [ 0 .. 10] 11 12 13 14
	//	So 14 in X actually translates to 3 in the original grid

	a = x % GRID_LEN
	b = y % GRID_WIDTH

	if a < 0 {
		a += GRID_LEN
	}
	if b < 0 {
		b += GRID_WIDTH
	}

	aDiff := int(math.Floor(float64(x) / float64(GRID_LEN)))
	bDiff := int(math.Floor(float64(y) / float64(GRID_WIDTH)))

	return a, b, []int{aDiff, bDiff} //, fmt.Sprintf("%d,%d", aDiff, bDiff)
}

var P1 int = 0

type Dimension [2]int
type MDP struct {
	x, y       int
	dimensions structures.Set[Dimension] //A dimension is a array with length two where [0] is how many x's away and [1] is y's away
}

func explore(grid *[][]rune, step int, exploreSet *map[Coord]MDP) {
	//As we walk
	//When we cross dimensions, we now have one position representing 2 distinct points
	//Every time we walk into a mapped point, if the source point is in a new dimension for the mapped point,
	//	increment move count by 1
	//A point is represented by an x, y and an array of Z dimensions
	//As we walk, if we cross into a new dimension
	//	loop through each Z point.
	//		determine the new dimension based upon the existing Z value
	//		the point is no longer in the old dimension, so the current Z values can go away
	//	Now we have the new Z values
	//	If this point (x, y) exists in the exploreSet
	//		then add any new Z values to the new points Z array and increment use count by 1
	//		!!!DONT NEED USE COUNT, we just need len(Z array)
	// If not in explore, add to explore
	if step > STEPS {
		return
	}
	if step == STEPS {
		fmt.Println("reached end len = ", len(*exploreSet))
		for _, point := range *exploreSet {
			P1 += len(point.dimensions)
		}
		fmt.Println("----------------")
		for i, row := range *grid {
			for j, col := range row {
				if mdp, found := (*exploreSet)[Coord{i, j}]; found {
					fmt.Printf("%d ", len(mdp.dimensions))
				} else {
					fmt.Print(string(col) + " ")
				}
			}
			fmt.Println()
		}
		return
	}

	fmt.Println("Exploring Step = ", step)

	toTry := map[Coord]MDP{}

	for _, point := range *exploreSet {
		// fmt.Printf("Trying Point (%d, %d)\n", point.x, point.y)
		// fmt.Println("Curr Dimensions = ", point.dimensions)
		//If this point is already in the tryMap, then two moves have merged, must merge the dimensions
		up := Coord{point.x + 1, point.y}
		if mdp, found := toTry[up]; found {
			//Then we have to merge
			// fmt.Println("Intersected UP, merging")
			// fmt.Println("Others Dimensions = ", mdp.dimensions)
			mdp.dimensions.AddAll(point.dimensions...)
			toTry[up] = mdp
			// fmt.Println("post merge = ", mdp.dimensions)
		} else {
			toTry[up] = point
		}

		down := Coord{point.x - 1, point.y}
		if mdp, found := toTry[down]; found {
			//Then we have to merge
			// fmt.Println("Intersected DOWN, merging")
			// fmt.Println("Others Dimensions = ", mdp.dimensions)
			mdp.dimensions.AddAll(point.dimensions...)
			toTry[down] = mdp
			// fmt.Println("post merge = ", mdp.dimensions)
		} else {
			toTry[down] = point
		}

		left := Coord{point.x, point.y - 1}
		if mdp, found := toTry[left]; found {
			//Then we have to merge
			// fmt.Println("Intersected LEFT, merging")
			// fmt.Println("Others Dimensions = ", mdp.dimensions)
			mdp.dimensions.AddAll(point.dimensions...)
			toTry[left] = mdp
			// fmt.Println("post merge = ", mdp.dimensions)
		} else {
			toTry[left] = point
		}

		right := Coord{point.x, point.y + 1}
		if mdp, found := toTry[right]; found {
			//Then we have to merge
			// fmt.Println("Intersected RIGHT, merging")
			// fmt.Println("Others Dimensions = ", mdp.dimensions)
			mdp.dimensions.AddAll(point.dimensions...)
			toTry[right] = mdp
			// fmt.Println("post merge = ", mdp.dimensions)
		} else {
			toTry[right] = point
		}
	}

	// fmt.Println("Trying ", toTry)

	results := make(map[Coord]MDP)
	for coord, mdp := range toTry {
		x, y := coord[0], coord[1]
		//If we have wrapped into a new dimension, account for that
		if x < 0 || y < 0 || x >= len(*grid) || y >= len((*grid)[0]) {
			var newDimension []int
			x, y, newDimension = mapCoordinateD(x, y)

			if (*grid)[x][y] == '#' {
				continue
			}
			// fmt.Println("----------")
			// fmt.Printf("Moving Coord (%d, %d) To New Dimension\n", coord[0], coord[1])

			//If we are now in a new dimension, we must check all the current dimensions
			//For each, change their value
			newDimensions := structures.Set[Dimension]{}

			// fmt.Println("New Dimension = ", newDimension)

			for _, dimension := range mdp.dimensions {
				// fmt.Println("Dim Old = ", dimension)
				newDimensions.Add(Dimension{dimension[0] + newDimension[0], dimension[1] + newDimension[1]})
				// fmt.Println("Dim After = ", Dimension{dimension[0] + newDimension[0], dimension[1] + newDimension[1]})
			}
			//Move to know point in new dimensions
			mdp.dimensions = newDimensions
			mdp.x, mdp.y = x, y

			if other, found := results[Coord{x, y}]; found {
				mdp.dimensions.AddAll(other.dimensions...)
			}

			// fmt.Println("Merged And Updated Dimensions = ", mdp.dimensions)

			// fmt.Printf("Moved To New dimension At (%d, %d)\n", x, y)

			results[Coord{x, y}] = mdp

		} else {
			if (*grid)[x][y] == '#' {
				continue
			}

			//If we have not wrapped into a new dimension, then we just move
			//everything at this coord to the new coord

			// fmt.Printf("Moving (%d, %d) to (%d, %d)\n", mdp.x, mdp.y, x, y)

			mdp.x, mdp.y = x, y
			if other, found := results[Coord{x, y}]; found {
				mdp.dimensions.AddAll(other.dimensions...)
			}
			results[coord] = mdp
		}
	}

	fmt.Println("----------------")
	for i, row := range *grid {
		for j, col := range row {
			if val, found := results[Coord{i, j}]; found {
				fmt.Printf("%d ", len(val.dimensions))
			} else {
				fmt.Print(string(col) + " ")
			}
		}
		fmt.Println()
	}

	explore(grid, step+1, &results)
}

var STEPS int =21

func exploreQuantumCount(grid *[][]rune, step int, exploreSet *map[Coord]int) int {
	if step > STEPS {
		return 0
	}
	if step == STEPS {
		for _, count := range *exploreSet {
			P1 += count
		}
	}

	fmt.Println("===================")
	fmt.Println("Explore Set = ", *exploreSet)
	// toTry := structures.List[Position]{}
	toTry := map[Coord]int{}

	//For each position in the explore set, check out each direction
	for p, useCount := range *exploreSet {

		// count, found := (*exploreSet)[p]
		//Need to remove ourself from current block
		// if found {
		// fmt.Printf("Moving From (%d, %d) Decrement Count\n", p[0], p[1])
		// (*exploreSet)[p] = count - useCount

		// if count - useCount == 0 {

		// }
		// }

		up := Position{Coord{p[0] - 1, p[1]}, useCount}
		if count, found := toTry[up.coords]; found {
			toTry[up.coords] = count + useCount
		} else {
			toTry[up.coords] = useCount
		}
		// toTry.Add(up)
		down := Position{Coord{p[0] + 1, p[1]}, useCount}
		if count, found := toTry[down.coords]; found {
			toTry[down.coords] = count + useCount
		} else {
			toTry[down.coords] = useCount
		}
		// toTry.Add(down)
		left := Position{Coord{p[0], p[1] - 1}, useCount}
		if count, found := toTry[left.coords]; found {
			toTry[left.coords] = count + useCount
		} else {
			toTry[left.coords] = useCount
		}
		// toTry.Add(left)
		right := Position{Coord{p[0], p[1] + 1}, useCount}
		if count, found := toTry[right.coords]; found {
			toTry[right.coords] = count + useCount
		} else {
			toTry[right.coords] = useCount
		}
		// toTry.Add(right)

		delete(*exploreSet, p)
	}

	fmt.Println("-------------")

	//For each move from each explore attempt
	//Validate and correlate it
	for coords, useCount := range toTry {

		// x, y := position.coords[0], position.coords[1]
		x, y := coords[0], coords[1]
		//If we have wrapped into a new plane, account for that
		if x < 0 || y < 0 || x >= len(*grid) || y >= len((*grid)[0]) {
			x, y = mapCoordinate(x, y)
		}
		// position.coords[0], position.coords[1] = x, y
		coords[0], coords[1] = x, y

		if x < 0 || y < 0 {
			panic("LESS THEN ZERO SOMEHOW")

		}

		//If this mapping is a rock, we can't use it
		if (*grid)[x][y] == '#' {
			continue
		}

		//If our new X and Y already map to something in exploreSet, increment it
		//Else init

		fmt.Printf("Position (%d, %d, %d) Is VAlid\n", x, y, useCount)

		count, found := (*exploreSet)[coords]
		if found {
			fmt.Println("Exists in map, adding")
			(*exploreSet)[coords] = count + useCount
		} else {
			fmt.Println("Doesn't exist, creating")
			(*exploreSet)[coords] = useCount
		}

		// count, found := (*exploreSet)[position.coords]
		// if found {
		// 	fmt.Println("Exists in map, adding")
		// 	(*exploreSet)[position.coords] = count + position.useCount
		// } else {
		// 	fmt.Println("Doesn't exist, creating")
		// 	(*exploreSet)[position.coords] = position.useCount
		// }
	}

	//Now on recurse, the exploreSet is already sanitized
	fmt.Println("Len = ", len(*exploreSet))
	exploreQuantumCount(grid, step+1, exploreSet)

	return 0
}

// func exploreQuantumI(grid *[][]rune, step int, exploreSet structures.Set[Position]) int {
// 	if step > STEPS {
// 		return 0
// 	}

// 	//The set has already removed duplicates, no need to do it twice
// 	validPositions := structures.List[Position]{}
// 	for _, position := range exploreSet {

// 		x, y := position.coords[0], position.coords[1]
// 		if x < 0 || y < 0 || x >= len(*grid) || y >= len((*grid)[0]) {
// 			// var dimension string
// 			x, y = mapCoordinate(x, y)
// 			// position.coords = Coord{x, y}
// 			// position.dimension = dimension
// 		}

// 		if (*grid)[x][y] == '#' {
// 			continue
// 		}
// 		validPositions = append(validPositions, position)
// 	}

// 	if step == STEPS {
// 		P1 = len(validPositions)
// 	}

// 	//Try All 4 directions
// 	possibleDestinations := structures.Set[Position]{}
// 	for _, p := range validPositions {
// 		possibleDestinations.Add(Position{Coord{p.coords[0] - 1, p.coords[1]}, p.useCount})
// 		possibleDestinations.Add(Position{Coord{p.coords[0] + 1, p.coords[1]}, p.useCount})
// 		possibleDestinations.Add(Position{Coord{p.coords[0], p.coords[1] - 1}, p.useCount})
// 		possibleDestinations.Add(Position{Coord{p.coords[0], p.coords[1] + 1}, p.useCount})
// 	}
// 	exploreQuantumI(grid, step + 1, possibleDestinations)

// 	return 0
// }

// func exploreQuantum(grid *[][]rune, step int, exploreSet structures.Set[Position]) int {
// 	if step > STEPS {
// 		return 0
// 	}

// 	//The set has already removed duplicates, no need to do it twice
// 	validPositions := structures.List[Position]{}
// 	for _, position := range exploreSet {

// 		x, y := position[0], position[1]
// 		if x < 0 || y < 0 || x >= len(*grid) || y >= len((*grid)[0]) {
// 			continue
// 		}

// 		if (*grid)[x][y] == '#' {
// 			continue
// 		}
// 		validPositions = append(validPositions, position)
// 	}

// 	if step == STEPS {
// 		P1 = len(validPositions)
// 	}

// 	//Try All 4 directions
// 	possibleDestinations := structures.Set[Position]{}
// 	for _, p := range validPositions {
// 		possibleDestinations.Add(Position{p[0] - 1, p[1]})
// 		possibleDestinations.Add(Position{p[0] + 1, p[1]})
// 		possibleDestinations.Add(Position{p[0], p[1] - 1})
// 		possibleDestinations.Add(Position{p[0], p[1] + 1})
// 	}
// 	exploreQuantum(grid, step + 1, possibleDestinations)

// 	return 0
// }
