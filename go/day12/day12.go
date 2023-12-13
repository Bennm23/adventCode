package main

import (
	"advent/lib"
	"fmt"
	"regexp"
	"strconv"
	"strings"

)

func main() {

	lines, err := lib.ReadFile("day12.txt")

	if err != nil {
		panic("Failed To Read Day12")
	}

	var maps []string
	var groups []string

	for _, line := range lines {

		split := strings.Split(line, " ")
		maps = append(maps, split[0])
		groups = append(groups, split[1])
	}

	for _, m := range maps {
		fmt.Println(m)
	}
	fmt.Println("Groups")
	for _, m := range groups {
		fmt.Println(m)
	}

	sum := 0
	for i, groupString := range groups {
		sum += solveRow(maps[i], groupString)
	}


	fmt.Println("Part 1 = ", sum)
}

var poundRegex regexp.Regexp = *regexp.MustCompile(`\#+`)

func solveRow(line, groupings string) int {

	// unknowns := questionRegex.FindAllStringIndex(line, -1);
	var group []int

	for _, g := range strings.Split(groupings, ",") {
		val, err := strconv.Atoi(g)

		if err != nil {
			panic("ATOI Failed")
		}
		group = append(group, val)
	}

	permutations := make([]string, 0)

	generatePermutations(&permutations, line, []rune(line), 0)

	sum := 0
	var match string = "*"
	for _, g := range group {

		for i:=0; i < g; i++ {
			match += "#"
		}
	}
	for _,p := range permutations {

		if validLine(p, group) {
			sum++
		}
	}

	return sum//7110
}

func validLine(line string, group []int) bool {
	groupings := poundRegex.FindAllStringIndex(line, -1)
	if len(groupings) != len(group) {
		return false
	}

	//For each set of #, if its length does not match group[i] return false
	for i, option := range groupings {

		if group[i] != (option[1] - option[0]) {
			return false
		}
	}

	return true
}

func generatePermutations(permutations *[]string, line string, permutation []rune, index int) {
	//At each index, try both '.' and '#'
	if index == len(line) {
		*permutations = append(*permutations, string(permutation))
		return
	}

	if line[index] != '?' {
		permutation[index] = rune(line[index])
		generatePermutations(permutations, line, permutation, index + 1)
	} else {
		permutation[index] = '#'
		generatePermutations(permutations, line, permutation, index + 1)
		permutation[index] = '.'
		generatePermutations(permutations, line, permutation, index + 1)
		permutation[index] = '?'
	}

}

type State struct {
	ID          int
	Transitions map[rune][]*State
	IsAccepting bool
	MatchCount  int
}

func NewState(id int, isAccepting bool) *State {
	return &State{
		ID:          id,
		Transitions: make(map[rune][]*State),
		IsAccepting: isAccepting,
		MatchCount:  0,
	}
}

func AddTransition(from, to *State, symbol rune) {
	from.Transitions[symbol] = append(from.Transitions[symbol], to)
}

func constructNFA(pattern string) *State {
	startState := NewState(0, false)
	currentState := startState

	for i, char := range pattern {
		nextState := NewState(i+1, false)

		switch char {
		case '#':
			AddTransition(currentState, nextState, '#')
		case '*':
			AddTransition(currentState, currentState, '\x00') // Epsilon transition (any character)
			AddTransition(currentState, nextState, '#')
		case '.':
			AddTransition(currentState, nextState, '\x00') // Epsilon transition (any character)
			AddTransition(currentState, nextState, '.')
		case '?':
			AddTransition(currentState, nextState, '#')
			AddTransition(currentState, nextState, '.')
		default:
			AddTransition(currentState, nextState, char)
		}

		currentState = nextState
	}

	currentState.IsAccepting = true
	return startState
}

func matchNFA(nfa *State, input string) int {
	currentStates := []*State{nfa}
	matches := 0

	for _, char := range input {
		var nextStates []*State

		for _, state := range currentStates {
			if transitions, exists := state.Transitions[char]; exists {
				nextStates = append(nextStates, transitions...)
			}
			if epsilonTransitions, exists := state.Transitions['\x00']; exists {
				nextStates = append(nextStates, epsilonTransitions...)
			}
		}

		if len(nextStates) == 0 {
			return 0
		}

		currentStates = nextStates
	}

	for _, state := range currentStates {
		if state.IsAccepting {
			state.MatchCount++
			matches += state.MatchCount
		}
	}

	return matches
}