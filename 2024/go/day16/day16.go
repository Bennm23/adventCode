package main

import (
	"advent/lib"
	"advent/lib/maths"
	"advent/lib/structures"
)

func main() {
    lib.RunAndScore("Part 1", p1)
    lib.RunAndScore("Part 2", p2)
}

type Node struct {
    position maths.Position
    direction maths.Position
    prev *Node
    cost int
}
type NodeKey struct {
    position maths.Position
    direction maths.Position
}

func buildInput() ([][]rune, maths.Position, maths.Position) {
    grid := lib.ReadFileToGrid("day16.txt")

    var startPos, goal maths.Position

    for rix, row := range grid {
        for cix, col := range row {
            if col == 'S' {
                startPos = maths.Position{X: rix, Y: cix}
            } else if col == 'E' {
                goal = maths.Position{X: rix, Y: cix}
            }
        }
    }
    return grid, startPos, goal
}

func p1() int {
    grid, startPos, goal := buildInput()

    visited := structures.NewSet[NodeKey]()


    explorationSet := structures.NewStack[Node]()
    explorationSet.Push(Node{startPos, maths.NewPosition(0, 1), nil, 0})

    bestCost := -1
    for !explorationSet.IsEmpty() {
        curr := explorationSet.Pop()

        if curr.position == goal {
            bestCost = curr.cost
            break
        }

        visited.Insert(NodeKey{curr.position, curr.direction})

        linearPos := curr.position.Add(curr.direction)
        //If linear in bounds and not blocked
        if linearPos.InBounds(len(grid)) && grid[linearPos.X][linearPos.Y] != '#' {
            newNode := Node{linearPos, curr.direction, &curr, curr.cost + 1}
            if !visited.Contains(NodeKey{newNode.position, newNode.direction}) {
                explorationSet.PushEval(newNode, func(a, b Node) bool {
                    return a.cost < b.cost
                })
            }
        }

        //Try Clockwise turn
        clockwiseDir := curr.direction.TurnClockwise();
        clockwiseNode := Node{curr.position, clockwiseDir, &curr, curr.cost + 1000}
        if !visited.Contains(NodeKey{clockwiseNode.position, clockwiseNode.direction}) {
            explorationSet.PushEval(clockwiseNode, func(a, b Node) bool {
                return a.cost < b.cost
            })
        }

        counterClockwiseDir := curr.direction.TurnCounterClockwise();
        counterClockwiseNode := Node{curr.position, counterClockwiseDir, &curr, curr.cost + 1000}
        if !visited.Contains(NodeKey{counterClockwiseNode.position, counterClockwiseNode.direction}) {
            explorationSet.PushEval(counterClockwiseNode, func(a, b Node) bool {
                return a.cost < b.cost
            })
        }
    }
    return bestCost
}

type ExplorationSet = structures.Set[Node]


func p2() int {
    grid, startPos, goal := buildInput()

    visited := structures.NewSet[NodeKey]()

    explorationSet := structures.NewStack[Node]()
    explorationSet.Push(Node{startPos, maths.NewPosition(0, 1), nil, 0})

    goodPositions := structures.NewSet[maths.Position]()

    for !explorationSet.IsEmpty() {
        curr := explorationSet.Pop()

        if curr.position == goal {
            pos := &curr;

            for {
                goodPositions.Insert(pos.position)
                pos = pos.prev
                if pos == nil {
                    break
                }
            }
            continue
        }

        visited.Insert(NodeKey{curr.position, curr.direction})

        linearPos := curr.position.Add(curr.direction)

        //If linear in bounds and not blocked
        if linearPos.InBounds(len(grid)) && grid[linearPos.X][linearPos.Y] != '#' {
            newNode := Node{linearPos, curr.direction, &curr, curr.cost + 1}
            if !visited.Contains(NodeKey{newNode.position, newNode.direction}) {
                explorationSet.PushEval(newNode, func(a, b Node) bool {
                    return a.cost < b.cost
                })
            }
        }

        //Try Clockwise turn
        clockwiseDir := curr.direction.TurnClockwise();
        clockwiseNode := Node{curr.position, clockwiseDir, &curr, curr.cost + 1000}
        if !visited.Contains(NodeKey{clockwiseNode.position, clockwiseNode.direction}) {
            explorationSet.PushEval(clockwiseNode, func(a, b Node) bool {
                return a.cost < b.cost
            })
        }

        counterClockwiseDir := curr.direction.TurnCounterClockwise();
        counterClockwiseNode := Node{curr.position, counterClockwiseDir, &curr, curr.cost + 1000}
        if !visited.Contains(NodeKey{counterClockwiseNode.position, counterClockwiseNode.direction}) {
            explorationSet.PushEval(counterClockwiseNode, func(a, b Node) bool {
                return a.cost < b.cost
            })
        }
    }

    return len(goodPositions)
}
