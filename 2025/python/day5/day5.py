
import sys
from typing import List, Tuple
from lib import read_to_input

lines = read_to_input("day5/input")

ranges = []
ids = []

for line in lines:
    if line is None or line == "":
        continue
    if "-" in line:
        parts = line.split("-")
        ranges.append((int(parts[0]), int(parts[1])))
    else:
        ids.append(int(line))

p1 = 0
for id in ids:
    for r in ranges:
        if id >= r[0] and id <= r[1]:
            p1 += 1
            break

def is_overlap(tup1: Tuple[int, int], tup2: Tuple[int, int]) -> bool:
    # Lower Overlap
    if tup1[0] <= tup2[0] and tup1[1] >= tup2[0]:
        return True
    
    # Upper overlap
    if tup1[1] > tup2[1] and tup1[0] <= tup2[1]:
        return True

    # Complete Overlap
    if tup1[1] > tup2[1] and tup1[0] < tup2[0]:
        return True

    return False

class Overlap:

    def __init__(self, rng: tuple[int, int] | None):
        if rng:
            self.groups: List[Tuple, Tuple] = [rng]
        else:
            self.groups: List[Tuple, Tuple] = []

    def in_any_range(self, rng: tuple[int, int]) -> bool:
        for range in self.groups:
            if is_overlap(range, rng) or is_overlap(rng, range):
                return True
        return False
    
    def reduce(self) -> int:
        minv = sys.maxsize
        maxv = -1

        for group in self.groups:
            minv = min(minv, group[0])
            maxv = max(maxv, group[1])

        return maxv - minv + 1

def merge_overlaps(overlaps: List[Overlap]) -> Tuple[List[Overlap], bool]:
    merge_indices: Tuple[int, int] | None = None
    for i in range(0, len(overlaps) - 1):
        for j in range(i+1, len(overlaps)):

            for item in overlaps[j].groups:
                if overlaps[i].in_any_range(item):
                    merge_indices = (i,j)
                    break
            if merge_indices is not None:
                break
        if merge_indices is not None:
            break

    if merge_indices is not None:
        new = Overlap(None)
        new.groups.extend(overlaps[merge_indices[0]].groups)
        new.groups.extend(overlaps[merge_indices[1]].groups)

        newlaps = []
        newlaps.append(new)
        for i in range(0, len(overlaps)):
            if i == merge_indices[0] or i == merge_indices[1]:
                continue
            newlaps.append(overlaps[i])
        
        return newlaps, True
        

    return overlaps, False

overlaps: List[Overlap] = []

for old in ranges:
    updated = False
    for group in overlaps:

        if group.in_any_range(old):
            group.groups.append(old)
            updated = True
            break

    if not updated:
        overlaps.append(Overlap(old))

    while True:
        newlaps, merged = merge_overlaps(overlaps)

        if merged:
            overlaps = newlaps
        else:
            break

p2 = 0
for lap in overlaps:
    p2 += lap.reduce()

print(f"Part 1: {p1}") #712
print(f"Part 2: {p2}") #332998283036769
