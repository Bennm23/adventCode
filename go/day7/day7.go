package main

import (
	"advent/lib"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

const (
	HIGH = 1
	PAIR = 2
	TWO_PAIR = 3
	THREE_KIND = 4
	FULL_HOUSE = 5
	FOUR_KIND = 6
	FIVE_KIND = 7
)
type Hand struct {

	bid int;
	runes []rune;
	handType int;
	test string;

}

func parseHand(handString string, bid int, rankings map[rune]int) Hand {
	var hand Hand

	hand.bid = bid
	hand.handType = HIGH
	hand.test = handString

	useCount := make(map[rune]int)

	for _, c := range handString {

		val, ok := useCount[c]
		
		hand.runes = append(hand.runes, c)

		if !ok {
			useCount[c] = 1
			continue
		}
		useCount[c] = val + 1
	}

	var bestKey rune
	bestScore := math.MinInt

	fmt.Println(hand)
	jokerCount, ok := useCount['J']
	if ok && len(useCount) != 1 {
		for k,_ := range useCount {
			if k == 'J' {
				continue
			}
			copy := lib.CopyMap[rune, int](useCount)

			copy[k] += jokerCount
			delete(copy, 'J')

			score := getHandType(copy)
			if score > bestScore {
				bestKey = k
				bestScore = score
			}

		}
		useCount[bestKey] += jokerCount
		delete(useCount, 'J')
	}

	hand.handType = getHandType(useCount)


	return hand
}

func getHandType(useCount map[rune]int) int {
	fmt.Println(useCount)
	if len(useCount) == 1 {
		return FIVE_KIND
	} else if len(useCount) == 2 {//FOUR OF A KIND OR FULL HOUSE
		for _, v := range useCount {
			if v == 2 || v == 3 {
				return FULL_HOUSE
			} else {
				return FOUR_KIND
			}
		}
	} else if len(useCount) == 3 {//TWO_PAIR or THREE_KIND

		for _,v := range useCount {
			if v == 2 {
				return TWO_PAIR
			} else if v == 3 {
				return THREE_KIND
			}
		}
	} else if len(useCount) == 4 {
		return PAIR
	} else if len(useCount) == 5 {
		return HIGH
	}

	panic("UHOH!")
}

func main() {
	lines := lib.ReadFile("day7.txt")

	var hands []Hand

	rankings := make(map[rune]int)
	rankings['A'] = 1
	rankings['K'] = 2
	rankings['Q'] = 3
	// rankings['J'] = 4
	rankings['T'] = 5
	rankings['9'] = 6
	rankings['8'] = 7
	rankings['7'] = 8
	rankings['6'] = 9
	rankings['5'] = 10
	rankings['4'] = 11
	rankings['3'] = 12
	rankings['2'] = 13
	rankings['J'] = 14

	for _, l := range lines {
		splits := strings.Split(l, " ")
		hand := splits[0]
		bid, err := strconv.Atoi(splits[1])
		if err != nil {
			panic("Failed to Atoi")
		}
		hands = append(hands, parseHand(hand, bid, rankings))
	}

	sort.Slice(hands[:], func(i, j int) bool {

		if hands[i].handType != hands[j].handType {
			return hands[i].handType < hands[j].handType
		} else {
			for index, runeVal := range hands[i].runes {
				if runeVal != hands[j].runes[index] {

					return rankings[runeVal] > rankings[hands[j].runes[index]]
				}
			}
		}

		return true
	})

	// fmt.Println(hands)

	score := 0
	for rank, hand := range hands {

		fmt.Println("+ ", rank + 1 , " * ", hand.bid)
		score += (rank + 1) * hand.bid

	}

	fmt.Println(score)//251481660
}
