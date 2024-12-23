package main

import (
	"advent/lib"
	"advent/lib/structures"
	"fmt"
	"slices"
	"strings"
)

func main() {
    lib.RunAndScore("Part 1", p1)//Result = 1151 : Total Time 18586 us
    lib.RunAndScore("Part 2", p2)//Result = ar,cd,hl,iw,jm,ku,qo,rz,vo,xe,xm,xv,ys : Total Time 71235 us
}

type ConnectionMap = map[string]structures.Set[string]

func buildConnectionMap() ConnectionMap {
    lines := lib.ReadFile("day23.txt")

    computer_mapping := make(ConnectionMap)

    for _, line := range lines {
        split := strings.Split(line, "-")
        left, right := split[0], split[1]

        try_insert(&computer_mapping, left, right)
        try_insert(&computer_mapping, right, left)
    }

    return computer_mapping
}

//Find all connections from a given computer
// for computer in connections
//  
//  if computer connections does not contain all previous
//   continue
//  else we found a new grouping
//  add grouping, and search again from this computer
//  including curr in visited
//      
func findAllInterConnections(
    connections *ConnectionMap,
    computer string,
    previous structures.Set[string],
    group []string,
    groupings *[][]string,
) {

    connected, found := (*connections)[computer]
    if !found {
        panic("Computer not connected to anything")
    }

    for connection := range connected {
        //Loop through all connected computers
        local_connections, found := (*connections)[connection]
        if !found {
            panic("LOCAL NOT FOUND")
        }
        //Local connections must hold all previous
        if !local_connections.ContainsAll(previous.Items()...) {
            continue
        }

        //If this connection is connected to all previous, recurse

        new_group := make([]string, 0);
        new_group = append(new_group, group...)
        new_group = append(new_group, connection)

        *groupings = append(*groupings, new_group)

        previous.Insert(connection)

        findAllInterConnections(
            connections,
            connection,
            previous,
            new_group,
            groupings,
        )
    }
}

func p1() int {

    computer_mapping := buildConnectionMap();
    groupings := structures.NewSet[[3]string]()

    for conn1, map1 := range computer_mapping {

        //Map1 is all connections to conn1
        //For each conn, iterate through all of its children
        for conn2 := range map1 {

            //map2 is all connections to conn2
            map2, found := computer_mapping[conn2]
            if !found {
                continue
            }
            for conn3 := range map2 {
                if map1.Contains(conn3) {
                    strv := []string{conn1, conn2, conn3}
                    slices.Sort(strv);
                    groupings.Insert([3]string(strv))
                }
            }
        }
    }
    sum := 0

    for group := range groupings {
        valid := false
        for _, g := range group {
            if g[0] == 't' {
                valid = true
            }
        }
        if !valid {
            continue
        }
        sum += 1
    }
    return sum
}

func try_insert(computers *ConnectionMap, key, connection string) {
    connections, found := (*computers)[key]

    if !found {
        connections = structures.NewSet[string]()
    }
    connections.Insert(connection)
    (*computers)[key] = connections
}

func p2() string {
    computer_mapping := buildConnectionMap();
    groupings := make([][]string, 0)

    for comp := range computer_mapping {
        findAllInterConnections(
            &computer_mapping,
            comp,
            structures.Set[string]{comp: struct{}{}},
            []string{comp},
            &groupings,
        )
    }

    best_group := make([]string, 0)

    for _, group := range groupings {
        if len(group) > len(best_group) {
            best_group = group
        }
    }

    slices.Sort(best_group)
    fmt.Println("Best Group = ", best_group)

    res := ""
    for i, g := range best_group {
        res += g
        if i != len(best_group) - 1 {
            res += ","
        }
    }
    fmt.Println("Password = ", res)

    return res
}
