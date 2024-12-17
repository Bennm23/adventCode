package main

import (
"advent/lib"
"fmt"
"strconv"
"strings"
)

func main() {
    lib.RunAndScore("Part 1", p1)
    lib.RunAndScore("Part 2", p2)
}

const (
    A = iota
    B
    C
)
func buildInput() ([]int, []int) {

    lines := lib.ReadFile("day17.txt")

    registers := make([]int, 0)

    for i := 0; i < 3; i++ {
        registers = append(registers, lib.ParseIntFromString(lines[i]))
    }

    program := lib.SplitStringToInts(strings.Split(lines[4], " ")[1], ",")

    return registers, program
}

func getOperandValue(instruction int, registers []int, val int) int {
    switch instruction {
    //Combo
    case 0, 2, 5, 6, 7:
        if val < 4 {
            return val
        }
        return registers[val - 4]
    default:
        //Literal
        return val
    }
}
func solveProgram(registers, program []int) string {
    ip := 0
    output := make([]string, 0)

    for ip < len(program) - 1 {
        instruction := program[ip]
        operand := getOperandValue(instruction, registers, program[ip + 1])
        switch instruction {
        case 0:
            registers[A] = registers[A] >> operand
            ip += 2
        case 1:
            registers[B] = registers[B] ^ operand
            ip += 2
        case 2:
            registers[B] = operand % 8
            ip += 2
        case 3:
            if registers[A] != 0 {
                ip = operand
            } else {
                ip += 2
            }
        case 4:
            registers[B] = registers[B] ^ registers[C]
            ip += 2
        case 5:
            output = append(output, strconv.Itoa((operand % 8)))
            ip += 2
        case 6:
            registers[B] = registers[A] >> operand
            ip += 2
        case 7:
            registers[C] = registers[A] >> operand
            ip += 2
        }
    }

    outs := ""
    for ix, out := range output {
        outs += out
        if ix != len(output) - 1 {
            outs += ","
        }
    }

    return outs
}
func p1() string {
    registers, program := buildInput()
    fmt.Println("Registers = ", registers)
    fmt.Println("Program = ", program)

    return solveProgram(registers, program)
}
func p2() int {
    _, program := buildInput()
    return solve(0, program)
}

/**
Pattern
- B = A % 8
- B = B ^ 5
- C = A >> B
- B = B ^ 6
- A = A >> 3
- B = B ^ C
- PRINT B%8
- if a != 0 GOTO 0


- A = 0 at end so must be 0-7
- A = A >> 3 + ? A only updated once
- B always 0-7
*/
func solve(ans int, program []int) int {
    fmt.Println("Ans = ", ans, " Prog = ", program)
    if len(program) == 0 {
        return ans
    }

    goal := program[len(program) - 1]

    a, b, c := 0, 0, 0

    for i := 0; i < 8; i++ {
        //The previous possible As are always the last A left shifted 3 plus any modulus of 8
        a = (ans << 3) + i
        // fmt.Println("A = ", a)

        b = a % 8
        b = b ^ 5
        c = a >> b
        b = b ^ 6
        // a = a >> 3
        b = b ^ c

        if b % 8 == goal {
            res := solve(a, program[:len(program)-1])
            if res != -1 {
                return res
            }
        }
    }
    fmt.Println("No Solution Found")
    return -1
}