package main

import (
	"advent/lib"
	"advent/lib/avstrings"
	"advent/lib/maths"
	"fmt"
	"strings"
)

type Node struct {
	name string;
	leftNode string;
	rightNode string;
}

func (n *Node) str() string {
	return fmt.Sprintf("%s [%s, %s]", n.name, n.leftNode, n.rightNode)	
}
func main() {
	lines, err := lib.ReadFile("day8.txt")
	if err != nil {
		panic("Failed to read day8")
	}
	nodes := make(map[string]*Node)
	var startNodes []*Node

	for _, line := range lines[2:] {
		split := strings.Split(line, " = ")
		key := split[0]
		vals := strings.Split(avstrings.ParseTextInParens(split[1]), ", ")

		node := Node {
			key,
			vals[0],
			vals[1],
		}
		nodes[key] = &node

		if key[2] == 'A' {
			startNodes = append(startNodes, &node)
		}
	}

	fmt.Println("Part 1 = ", part1(nodes, lines[0]))//21409
	fmt.Println("Part 2 = ", part2(nodes, lines[0], startNodes))//21165830176709
}

func part2(nodes map[string]*Node, commands string, startNodes []*Node) int64 {
	var solutions []int64

	for _, starter := range startNodes {
		solutions = append(
			solutions,
			solveAtNode(nodes, commands, starter.name, func(s string) bool {
				return s[2] == 'Z'
			}),
		)
	}

	var lcm int64
	lcm = maths.Lcm[int64](solutions[0], solutions[1])

	for _, solution := range solutions[2:] {
		lcm = maths.Lcm[int64](lcm, solution)
	}

	return lcm

}
func part1(nodes map[string]*Node, commands string) int64 {
	return solveAtNode(nodes, commands, "AAA", func(s string) bool {
		return s == "ZZZ"
	})
}

type Completer func(string) bool

func solveAtNode(nodes map[string]*Node, commands string, start string, completer Completer) int64 {
	var count int64 = 0
	curr, ok := nodes[start]

	if !ok {
		panic("Start String Not Present in Nodes")
	}

	fmt.Println("Solving at ", curr.str())

	solved := false;

	for !solved {

		for _, c := range commands {
			count++
			if c == 'L' {
				curr, ok = nodes[curr.leftNode]
				if !ok {
					panic("Left Move Failed to find node")
				}
			} else {
				curr, ok = nodes[curr.rightNode]
				if !ok {
					panic("Right Move Failed to find node")
				}
			}

			if completer(curr.name) {
				solved = true;
				break;
			}
		}
	}
	return count

}
