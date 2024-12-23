package main

import (
	"advent/lib"
	"advent/lib/maths"
	"advent/lib/structures"
)

func main() {
    lib.RunAndScore("Part 1", p1)//Result = 17577894908 : Total Time   33338 us
    lib.RunAndScore("Part 2", p2)//Result = 1931        : Total Time 1271308 us
}

func buildInput() []int {
    return lib.ReadFileToTypeVec("day22.txt", func(s string) int {
        return maths.ToInt(s);
    })
}
const PRUNE = 16777216

func p1() int {
    sum := 0
    secrets := buildInput()

    for _, secret := range secrets {
        new_secret := secret
        for range 2000 {
            new_secret = ((new_secret * 64) ^ new_secret) % PRUNE
            new_secret = ((new_secret / 32) ^ new_secret) % PRUNE
            new_secret = ((new_secret * 2048) ^ new_secret) % PRUNE
        }
        sum += new_secret
    }

    return sum
}
func p2() int {
    secrets := buildInput()

    sequence_score := make(map[[4]int]int, 0)

    for _, secret := range secrets {
        new_secret := secret
        prev_bannanas := secret % 10
        sequence := make([]int, 0)

        found_this_secret := structures.NewSet[[4]int]()

        for range 2000 {
            new_secret = ((new_secret * 64) ^ new_secret) % PRUNE
            new_secret = ((new_secret / 32) ^ new_secret) % PRUNE
            new_secret = ((new_secret * 2048) ^ new_secret) % PRUNE

            bannanas := new_secret % 10
            delta := bannanas - prev_bannanas
            prev_bannanas = bannanas
            sequence = append(sequence, delta)

            if len(sequence) != 4 {
                continue
            }

            foured := [4]int(sequence)
            val, found := sequence_score[foured]
            if !found {
                sequence_score[foured] = bannanas
            } else if !found_this_secret.Contains(foured) {
                sequence_score[foured] = bannanas + val
            }
            found_this_secret.Insert(foured)
            sequence = sequence[1:]
        }
    }

    max := 0
    for _, total := range sequence_score {
        if total > max {
            max = total
        }
    }

    return max
}
