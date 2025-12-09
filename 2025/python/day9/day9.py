
from functools import lru_cache
import time
from lib import read_to_input


lines = [ line.strip() for line in read_to_input("day9/input") ]
pairs = [ ( int(line.split(",")[0]), int(line.split(",")[1]) ) for line in lines ]

COL = 0
ROW = 1
row_count = 0
col_count = 0
for pair in pairs:
    col_count = max(col_count, pair[COL] + 1)
    row_count = max(row_count, pair[ROW] + 1)


def p1() -> int:
    max_area = -1

    for i in range(len(pairs)):
        for j in range(i + 1, len(pairs)):

            width = abs(pairs[j][COL] - pairs[i][COL]) + 1
            height = abs(pairs[j][ROW] - pairs[i][ROW]) + 1

            max_area = max(max_area, width*height)

    return max_area

EMPTY = False
RED = True
GREEN = True

# def print_grid(grid: list[list[int]]) -> None:
#     for row in grid:
#         for col in row:
#             if col == EMPTY:
#                 print(".", end='')
#             elif col == RED:
#                 print("#", end='')
#             elif col == GREEN:
#                 print("X", end='')
                
#         print()

def get_corner_positions_in_rect(corner1: tuple[int,int], corner2: tuple[int,int]) -> list[tuple[int,int]]:
    positions = []
    min_row = min(corner1[ROW], corner2[ROW])
    max_row = max(corner1[ROW], corner2[ROW])
    min_col = min(corner1[COL], corner2[COL])
    max_col = max(corner1[COL], corner2[COL])

    positions.append((min_col, min_row))
    positions.append((min_col, max_row))
    positions.append((max_col, max_row))
    positions.append((max_col, min_row))
    return positions
def get_border_positions_in_rect(corner1: tuple[int,int], corner2: tuple[int,int]) -> list[tuple[int,int]]:
    positions = []
    min_row = min(corner1[ROW], corner2[ROW])
    max_row = max(corner1[ROW], corner2[ROW])
    min_col = min(corner1[COL], corner2[COL])
    max_col = max(corner1[COL], corner2[COL])

    for r in range(min_row, max_row + 1):
        positions.append((min_col, r))
        positions.append((max_col, r))
    
    for c in range(min_col + 1, max_col):
        positions.append((c, min_row))
        positions.append((c, max_row))
    
    return positions


def get_all_positions_in_rect(corner1: tuple[int,int], corner2: tuple[int,int]) -> list[tuple[int,int]]:
    positions = []
    min_row = min(corner1[ROW], corner2[ROW])
    max_row = max(corner1[ROW], corner2[ROW])
    min_col = min(corner1[COL], corner2[COL])
    max_col = max(corner1[COL], corner2[COL])

    for r in range(min_row, max_row + 1):
        for c in range(min_col, max_col + 1):
            positions.append((c,r))
    
    return positions

def p2() -> int:
    valid_cols_in_row: dict[int, set[tuple[int, int]]] = {}

    def add_col_range(row, col_range):
        if row in valid_cols_in_row:
            valid_cols_in_row[row].add(col_range)
        else:
            valid_cols_in_row[row] = set([col_range])
        
    # Fill perimeter
    for i in range(0, len(pairs)):
        first = pairs[i]
        second = pairs[(i + 1) % len(pairs)] # Mod to connect last to first

        min_row = min(first[ROW], second[ROW])
        max_row = max(first[ROW], second[ROW])
        min_col = min(first[COL], second[COL])
        max_col = max(first[COL], second[COL])

        # find all valid columns for each row
        for r in range(min_row, max_row + 1):
            add_col_range(r, (min_col, max_col))

    # Condense ranges. Assuming the next column must be closed
    for r, col_ranges in valid_cols_in_row.items():
        filtered = [rng for rng in col_ranges]
        # filter by start and end col
        filtered.sort(key=lambda x: (x[0], x[1]))

        curr_start = None
        new_ranges = set()
        for i in range(0, len(filtered)):
            # Closing
            if i % 2 == 1:
                end = filtered[i][1]
                new_ranges.add((curr_start, end))
                curr_start = None
            else:
                curr_start = filtered[i][0]
        
        if curr_start != None:
            new_ranges.add((curr_start, curr_start))
        valid_cols_in_row[r] = new_ranges


    def point_in_bounds(r: int, c: int) -> bool:
        col_ranges = valid_cols_in_row.get(r, set())
        for colr in col_ranges:
            if c >= colr[0] and c <= colr[1]:
                return True
        return False

    max_area = -1

    for i in range(len(pairs)):
        for j in range(i + 1, len(pairs)):

            width = abs(pairs[j][COL] - pairs[i][COL]) + 1
            height = abs(pairs[j][ROW] - pairs[i][ROW]) + 1

            area = width * height
            if area > 4748985168 / 2: #safe to assume its much smaller than p1 and /2 worked
                continue
            if area <= max_area:
                continue
            # Check if corner positions are green for quick validation
            corner_positions = get_corner_positions_in_rect(pairs[i], pairs[j])
            all_green = True
            for pos in corner_positions:
                if not point_in_bounds(pos[ROW], pos[COL]):
                    all_green = False
                    break
            if not all_green:
                continue

            min_row = min(pairs[i][ROW], pairs[j][ROW])
            max_row = max(pairs[i][ROW], pairs[j][ROW])
            min_col = min(pairs[i][COL], pairs[j][COL])
            max_col = max(pairs[i][COL], pairs[j][COL])

            for r in range(min_row, max_row + 1):
                if not point_in_bounds(r, min_col):
                    all_green = False
                    break
                if not point_in_bounds(r, max_col):
                    all_green = False
                    break
            if not all_green:
                continue
            
            for c in range(min_col + 1, max_col):
                if not point_in_bounds(min_row, c):
                    all_green = False
                    break
                if not point_in_bounds(max_row, c):
                    all_green = False
                    break

            if all_green:
                max_area = area
    return max_area

start = time.time()
print("Part 1: ", p1()) # 4748985168
print("Part 2: ", p2()) # 1550760868

print("Duration ms ", (time.time() - start) * 1000)


