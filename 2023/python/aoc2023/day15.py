from utils import filePath, runAndPrintMillis

class Element:
    label = ""
    lenseVal = -1
    def __init__(self, label: str, lenseVal: int):
        self.label = label
        self.lenseVal = lenseVal

class Box:
    id = -1
    lenses = []
    def __init__(self, id: int):
        self.id = id
        self.lenses = []

    def append(self, label, lenseVal):

        self.lenses.append(Element(label, lenseVal))

    def contains(self, label: str) -> (bool, int):
        for i, lense in enumerate(self.lenses):
            if lense[0] == label:
                return (True, i)

        return (False, -1)

def hash(string: str) -> int:

    total = 0
    
    for char in string:
        total += ord(char)
        total *= 17
        total %= 256

    return total
    

def solve():
    sum = 0
    line = open(filePath("day15.txt")).readline().strip()
        
    labels = line.split(",")
    
    boxes = []
    
    for i in range(0, 256):
        boxes.append(Box(i))
    
    for tag in line.split(","):
        sum += hash(tag)

        equalIdx = tag.find("=")

        #Then need to add this label and lense val to this box
        if equalIdx != -1:
            label = tag[:equalIdx]
            boxNum = hash(label)
            lenseVal = int(tag[equalIdx + 1:])

            found, index = boxes[boxNum].contains(label)
            
            if found:
                boxes[boxNum].lenses[index][1] = lenseVal
            else:
                boxes[boxNum].append(label, lenseVal)
            
            
    
    p2 = 0
    for box in boxes:
        
        for slot, pair in enumerate(box.lenses):
            p2 += (box.id + 1)*(slot + 1)*(pair[1])

            
    print(sum)#P1 = 515974
    print("Part 2 = ", p2)

solve()
