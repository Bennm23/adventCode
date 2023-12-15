package main

import (
	"advent/lib"
	"fmt"
	"regexp"
	"strings"
)

func main() {

	lib.RunAndPrintDuration(func() {
		lines := lib.ReadFile("day4.txt")

		p1, p2 := solve(lines)
		fmt.Println("PART 1 = ", p1)//24542
		fmt.Println("PART 2 = ", p2)//8736438
	})//51529
}

type SetMap map[string]struct{}

type ScoreMap map[int]int

func solve(lines []string) (int, int) {
	sum := 0;
	numFinder := regexp.MustCompile(`[0-9]+`)
	winnerMap := make(ScoreMap)

	for row, line := range lines {
		card := strings.Split(line, ":")[1]

		winners := make(SetMap)

		split := strings.Split(card, "|")
		winSplit := split[0]
		numSplit := split[1]

		winningNums := numFinder.FindAllStringIndex(winSplit, -1)
		myNums := numFinder.FindAllStringIndex(numSplit, -1)


		for _, winner := range winningNums {
			winners[winSplit[winner[0]:winner[1]]] = struct{}{} 
		}

		cardCount := 0;
		matches := 0;
		for _, num := range myNums {
			strval := numSplit[num[0]:num[1]]

			_, ok := winners[strval]

			if ok {
				delete(winners, strval)
				matches++
				if cardCount == 0 {
					cardCount = 1
				} else {
					cardCount = cardCount<<1
				}
			} 
		}
		winnerMap[row] = matches
		sum += cardCount
	}
	
	p2 := 0;

	for i := 0; i < len(lines); i++ {
		//iterate through the lines and solutions
		p2 += 1
		//if row numWinners == 0, continue
		numWinners, ok := winnerMap[i]
		if !ok {
			fmt.Println("Failed to Access I = " , i)
			panic("uhoh")
		}
		if numWinners == 0 {
			//When score is 0, we are done here
			continue
		}

		//When score is not 0, we must recurse through the next X rows
		total := calculateCards(&lines, i+1, numWinners, winnerMap)
		p2 += total
		
	}


	return sum, p2
}


func calculateCards(lines *[]string, startIndex int, numWinners int, winnerMap map[int]int) int {
	//loop through copy cards, startIndex .. startIndex + numWinners

	childCards := 0;

	for i := startIndex; i < startIndex + numWinners; i++ {
		childCards++
		newWinners, ok := winnerMap[i]
		if !ok {
			panic("Failed to query for map val in calculate")
		}

		if newWinners == 0 {
			continue;
		}
		childCards += calculateCards(lines, i+1, newWinners, winnerMap)
	}


	return childCards
}
