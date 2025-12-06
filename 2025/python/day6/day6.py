

from typing import List
from lib import read_to_input

file = "day6/input"

lines = read_to_input(file)
operations = [op for op in lines[-1].strip().split()]

def part1() -> int:
    columns: List[List[int]] = []
    for _ in range(0, len(operations)):
        columns.append([])

    for rdx, row in enumerate(lines[:-1]):
        cols = row.strip().split()
        for cix, col in enumerate(cols):
            columns[cix].append(int(col))

    p1 = 0
    for idx, column in enumerate(columns):
        if operations[idx] == "*":
            total = 1
        else:
            total = 0
        for c in column:
            if operations[idx] == "*":
                total *= c
            else:
                total += c

        p1 += total
    return p1


def convert_column(col: List[str]) -> List[int]:
    res = []
    # Convert the column of aligned strings
    # first element of result, is rightmost char of all rows combined
    for i in reversed(range(0, len(col[0]))):
        col_val = ""
        for s in col:
            col_val += s[i]
        res.append(int(col_val))
    return res

def part2() -> int:
    p2 = 0

    lines = []
    with open(file, "r") as f:
        # Only strip newline
        lines = [line[:-1] for line in f.readlines()]

    # Get the max digit in each column since they are aligned along that spacing
    digits_by_col = [-1 for _ in range(0, len(operations))]
    for line in lines[:-1]:
        s = line.strip().split()
        for i,item in enumerate(s):
            digits_by_col[i] = max(digits_by_col[i], len(item))

    columns: List[List[str]] = [[] for _ in range(0, len(operations))]

    # create the list of columns, each entry being len = digits_by_col[col] for that column
    for row in lines[:-1]:
        cix = 0
        i = 0
        while i < len(row):
            columns[cix].append(row[i:i + digits_by_col[cix]])
            i += digits_by_col[cix] + 1
            cix += 1

    # Convert columns to int lists
    converted = [ convert_column(col) for col in columns]
    for idx, column in enumerate(converted):
        if operations[idx] == "*":
            total = 1
            op = lambda total,v: total * v
        else:
            total = 0
            op = lambda total,v: total + v
        for c in column:
            if c == 0: # There isn't a single 0 in the input somehow but if there was
                continue
            total = op(total, c)

        p2 += total


    return p2

print("Part 1 = ", part1()) # 5361735137219
print("Part 2 = ", part2()) # 11744693538946