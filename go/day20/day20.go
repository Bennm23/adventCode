package main

import (
	"advent/lib"
	"advent/lib/structures"
	"fmt"
	"strings"
)

type MODULE_TYPE byte

const (
    BROADCAST   MODULE_TYPE = 1
    FLIPFLOP    MODULE_TYPE = 2
    CONJUNCTION MODULE_TYPE = 3
)

type Module struct {
    destinations []string
    moduleType   MODULE_TYPE
    name         string

    onOff  bool
    memory map[string]bool
}

func BuildModule(split []string) Module {
    destinations := strings.Split(split[1], ",")
    d := make([]string, 0)
    for _, dest := range destinations {
        d = append(d, strings.Trim(dest, " "))
    }
    module := Module{
        destinations: d,
    }
    name := strings.Split(split[0], " ")[0]

    switch name[0] {
    case '%':
        module.moduleType = FLIPFLOP
        module.name = name[1:]
        module.onOff = false
    case '&':
        module.moduleType = CONJUNCTION
        module.name = name[1:]
        module.memory = make(map[string]bool)
    default:
        module.moduleType = BROADCAST
        module.name = name
    }
    return module
}

var lowPulses int64 = 0
var highPulses int64 = 0

func main() {

    lines := lib.ReadFile("day20.txt")

    moduleMap := make(ModuleMap)

    for _, line := range lines {
        module := BuildModule(strings.Split(line, " -> "))
        moduleMap[module.name] = &module
    }

    for _, module := range moduleMap {
        if module.moduleType != CONJUNCTION {
            continue
        }
        //For each module, if that module maps to this module, then add it to memory
        for _, other := range moduleMap {
            for _, dest := range other.destinations {
                if dest == module.name {
                    module.memory[other.name] = false
                }
            }
        }
    }

    broadcaster, found := moduleMap["broadcaster"]
    if !found {
        panic(moduleMap)
    }

	p1 := int64(0)

    i := 0
    for {

        //At the start of each for, push button
        lowPulses++
		reactions := structures.NewStack[Pulse]()
        results := broadcaster.sendPulse(false, moduleMap)
		reactions.PushAll(results)

        for reactions.Size() > 0 {
			reaction := reactions.Pop()

            sender, found := moduleMap[reaction.source]
            if !found {
                panic(found)
            }
			reactions.PushAll(sender.sendPulse(reaction.onOff, moduleMap))
        }

        if i == 999 {
            p1 = lowPulses * highPulses
            fmt.Println("Part One = ", p1)//737679780
            break
        }
        i++
    }

}

type ModuleMap map[string]*Module

func (module *Module) sendPulse(onOff bool, moduleMap ModuleMap) []Pulse {
    results := make([]Pulse, 0)

    for _, child := range module.destinations {
        if onOff {
            highPulses++
        } else {
            lowPulses++
        }
        c, found := moduleMap[child]
        if !found {
            continue
        }

        res, use := c.receivePulse(onOff, module.name)
        if use {
            results = append(results, res)
        }
    }
    return results
}

type Pulse struct {
    onOff  bool
    source string
}

func (module *Module) receivePulse(onOff bool, source string) (Pulse, bool) {

    outPulse := Pulse{source: module.name}
    if module.name == "rx" {
        fmt.Println("RX RECEIVED ", onOff)
    }

    switch module.moduleType {
    case BROADCAST:
        panic("BROADCAST RECEIVED PULSE")

    case FLIPFLOP:
        //Prefix %: Either On or Off, initally off. do nothing on high pulse, on low pulse toggle
        if !onOff {
            module.onOff = !module.onOff
            outPulse.onOff = module.onOff
            return outPulse, true
        }

    case CONJUNCTION:
        //Store the most recent pulse type from all connected input modules
        //Default is low pulse.

        //On pulse receive, update memory for input,
        // if all high pulses in memory send low pulse, else send high
        module.memory[source] = onOff

        allHigh := true
        for _, v := range module.memory {
            if !v {
                allHigh = false
                break
            }
        }
        outPulse.onOff = !allHigh
        return outPulse, true
    }

    return outPulse, false
}