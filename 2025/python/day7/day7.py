from typing import Optional, Set, Tuple
from lib import read_to_input

type Position = tuple[int, int]

ROW = 0
COL = 1
SPLITTER = '^'

grid = [[c for c in line] for line in read_to_input("day7/input")]
GRID_HEIGHT = len(grid)
GRID_WIDTH = len(grid[0])

start_pos: Position = (0, 0)
splitters: Set[Position] = set()

for rix,row in enumerate(grid):
    print("".join(row))
    for cix, col in enumerate(row):
        if col == 'S':
            start_pos = (rix, cix)
        elif col == SPLITTER:
            splitters.add((rix, cix))


def get_split(position: Position) -> Optional[tuple[Position, Position]]:
    if grid[position[ROW]][position[COL]] != SPLITTER:
        return None
    
    left_split = (position[ROW], position[COL] - 1)
    right_split = (position[ROW], position[COL] + 1)

    return (left_split, right_split)


def count_splits(position: Position, visited: Set[Position], post_split: bool = True) -> int:

    # We already get the split pos
    if post_split:
        new_pos = (position[ROW], position[COL])
    else:
        new_pos = (position[ROW] + 1, position[COL])


    if new_pos[ROW] >= GRID_HEIGHT:
        return 0
    if new_pos[COL] < 0 or new_pos[COL] >= GRID_WIDTH:
        return 0

    # Is this possible for non splits?
    if new_pos in visited:
        return 0

    visited.add(new_pos)

    # Check for splitter
    splits = get_split(new_pos)
    split_count = 0

    if splits is None:
        return count_splits(new_pos, visited, False)
    else:

        left_pos, right_pos = splits
        split_count = 1

        if left_pos not in visited:
            split_count += count_splits(left_pos, visited)
        
        if right_pos not in visited:
            split_count += count_splits(right_pos, visited)

    return split_count

from functools import lru_cache

@lru_cache(maxsize=None)
def count_timelines(position: Position, post_split: bool = True) -> int:

    # We already get the split pos
    if post_split:
        new_pos = (position[ROW], position[COL])
    else:
        new_pos = (position[ROW] + 1, position[COL])


    if new_pos[ROW] >= GRID_HEIGHT:
        return 0
    if new_pos[COL] < 0 or new_pos[COL] >= GRID_WIDTH:
        return 0

    # Check for splitter
    splits = get_split(new_pos)
    timeline_count = 0

    if splits is None:
        return count_timelines(new_pos, False)
    else:

        left_pos, right_pos = splits
        timeline_count = 1 # At split, 1 path becomes 2
        timeline_count += count_timelines(left_pos)
        timeline_count += count_timelines(right_pos)

    return timeline_count

visited = set()
def print_visited():
    for rix, row in enumerate(grid):
        for cix, col in enumerate(row):
            if (rix, cix) in visited and grid[rix][cix] != SPLITTER:
                print("|", end='')
            else:
                print(col, end='')
        print()

p1 = count_splits(start_pos, visited)
print("Part 1 = ", p1) #1687
p2 = count_timelines(start_pos) + 1 # Add 1 for the initial timeline
print("Part 2: ", p2) # 390684413472684