from collections import namedtuple
from queue import PriorityQueue
import sys
import time
from typing import Dict
from utils import filePath


State = namedtuple("State", ["position", "dir", "movesSoFar"])
    
MOVES = [(1, 0), (-1, 0), (0, 1), (0, -1)]

def solve(grid, maxMoves = 3, minMoves = 1) -> int:
    openSet = PriorityQueue()
    best : Dict[State, int] = {}
    
    for move in MOVES:
        openSet.put((0, State(position=(0, 0), dir=move, movesSoFar=0)))

    while not openSet.empty():
        (value, curr) = openSet.get()
        if curr in best:
            continue
        
        best[curr] = value

        for neighbor in getNeighbors(curr, maxMoves, minMoves):
            if neighbor not in best:
                weight = best[curr] + grid[neighbor.position[0]][neighbor.position[1]]
                openSet.put((weight, neighbor))
    
    l = {k:v for (k, v) in best.items() if k.position == (len(grid) - 1, len(grid[0]) - 1)}
    
    min = sys.maxsize
    for _, v in l.items():
        if v < min:
            min = v
    
    return min
    

def getNeighbors(state : State, maxMoves, minMoves) -> list[State]:
    neighbors = []
    
    for dir in MOVES:
        newPosition = (state.position[0] + dir[0], state.position[1] + dir[1])
        if dir == state.dir:
            moves = state.movesSoFar + 1
        else:
            moves = 1

        (x, y) = newPosition

        if moves > maxMoves:
            continue
        if dir != state.dir and state.movesSoFar < minMoves:
            continue
        if x < 0 or x >= GRID_LEN or y < 0 or y >= GRID_WIDTH:
            continue
        if dir[0] * -1 == state.dir[0] and dir[1] * -1 == state.dir[1]:
            continue

        neighbors.append(State(position=newPosition, dir=dir, movesSoFar=moves))

    return neighbors

    
start = time.perf_counter_ns() * 1e-6

fi = open(filePath("day17.txt"))
grid = [[int(x) for x in f[:-1]] for f in fi]
fi.close()

GRID_LEN = len(grid)
GRID_WIDTH = len(grid[0])

solve(grid)
print("Part 1 = ", solve(grid))
print("Part 2 = ", solve(grid, 10, 4))


end = time.perf_counter_ns() * 1e-6
print('Duration MS  = ', end - start)#8516

    