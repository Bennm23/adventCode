package main

import (
	"advent/lib"
	"fmt"
	"strconv"
	"strings"
)

type Entry struct {
	label string;
	lense int;
}
type Box struct {
	children []*Entry;
	boxNum int;
}
func (b *Box) containsLabel(label string) (int, bool) {

	for i, child := range b.children {
		if child.label == label {
			return i, true
		}
	}
	return -1, false
}

func (b *Box) addEntry(entry *Entry) {
	b.children = append(b.children, entry)
}

func (b *Box) removeIfExists(label string) {
	index, found := b.containsLabel(label)
	if !found {
		return
	}

	b.children = append(b.children[:index], b.children[index + 1:]...)
}

func NewBox(boxNum int) *Box{
	return &Box {
		make([]*Entry, 0),
		boxNum,
	}
}

func main() {

	labels := lib.ReadOneLineToChunks("day15.txt", ",")
	boxes := make([]*Box, 0);
	for i:=0; i < 256; i++ {
		boxes = append(boxes, NewBox(i))
	}

	var p1, p2 int = 0, 0
	for _, s := range labels {
		p1 += hashit(s)

		var boxNum int
		if strings.Contains(s, "=") {

			split := strings.Split(s, "=")
			label := split[0]
			boxNum = hashit(label)
			lensNum, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}

			index, found := boxes[boxNum].containsLabel(label);
			if found {
				//If the box contains the label already, replace the old lens
				boxes[boxNum].children[index].lense = lensNum
			} else {

				boxes[boxNum].addEntry(&Entry{label, lensNum})
			}

		} else {
			//Cut off the last '-'
			label := s[:len(s) - 1]
			boxNum = hashit(label)
			boxes[boxNum].removeIfExists(label)
		}

	}

	for i, box := range boxes {
		//For every lense in all boxes
		//sum += (i + 1)*(indexOf(lense)+1)*(lens number)
		for lx, lense := range box.children {

			p2 += (i + 1)*(lx + 1)*(lense.lense)
		}
	}

	fmt.Println("Part 1 = ", p1)//515974
	fmt.Println("Part 2 = ", p2)
}

func hashit(str string) int{
	var result int

	for _, c := range str {
		result += int(c)
		result *= 17
		result %= 256
	}

	return result
}


