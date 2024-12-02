package main

import (
	"advent/lib"
	. "advent/lib/structures"
	"fmt"
	"regexp"
	"strconv"
	"z3"
)

type State struct {
	position, velocity Vector[float64]
}

func (state State) print() {
	fmt.Println("Pos: ", state.position, " @ Vel: ", state.velocity)
}

func BuildState(ints []string) State {

	pos, vel := Vector[float64]{}, Vector[float64]{}
	for i, val := range ints {
		c, _ := strconv.Atoi(val)
		if i > 2 {
			vel = append(vel, float64(c))
		} else {
			pos = append(pos, float64(c))
		}
	}

	return State {
		position: pos,
		velocity: vel,
	}
}

var MIN, MAX float64 = 200000000000000, 400000000000000

func main() {

	lines := lib.ReadFile("day24.txt")

	intFind := regexp.MustCompile(`-?[\d]+`)

	conf := z3.NewConfig()
	ctx := z3.NewContext(conf)
	conf.Close()
	defer ctx.Close()

	states := []State{}

	for _, line := range lines {
		fmt.Println(line)

		state := BuildState(intFind.FindAllString(line, -1))
		state.print()
		states = append(states, state)
	}

	//States are at time 0
	//Velocity is how far they will move after 1 nano second

	count := 0

	for i, state := range states {

		for j:=i+1; j < len(states); j++ {
			other := states[j]

			pX, pY, found := gi(state, other)
			if !found {
				continue
			}

			if pX <= MAX && pX >= MIN && pY <= MAX && pY >= MIN {
				count++
				fmt.Printf("Found Intersect (%.2f, %.2f)\n", pX, pY)
			} else {
				fmt.Printf("Found OB Intersect (%.2f, %.2f)\n", pX, pY)
			}
		}
	}

	fmt.Println("Num Occurences = ", count)//P1 = 20963
}

func gi(s1, s2 State) (float64, float64, bool) {
	p, q, r, s := s1.position[:2], s2.position[:2], s1.velocity[:2], s2.velocity[:2]

	rs := r.SimpleCross(s)
	qpr := q.Minus(p).SimpleCross(r)
	if rs == 0 && qpr == 0 {
		fmt.Println("Colinear")
		return -1, -1, false
	}
	if rs == 0 {
		fmt.Println("Parallel")
		return -1, -1, false
	}

	t := (q.Minus(p)).SimpleCross(s.Divide(rs))
	u := (q.Minus(p)).SimpleCross(r.Divide(rs))

	intersect := p.Plus(r.Times(t))

	if t < 0  || u < 0 {
		fmt.Println("Intersect Backwards")
		return -1, -1, false
	}
		
	return intersect[0], intersect[1], true
}