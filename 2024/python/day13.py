
import re

from sympy import symbols, Eq, solve

def get_ints(s : str) -> list:
    return [int(v) for v in re.findall(r'\d+', s)]

with open("../inputs/day13.txt") as file:
    
    total = 0
    groups = []
    subgroup = []

    for line in file:
        if line == "\n":
            groups.append(subgroup)
            subgroup = []
        else:
            subgroup.append(line.strip())

    groups.append(subgroup)
    
    for group in groups:
        
        a = get_ints(group[0])
        b = get_ints(group[1])
        p = get_ints(group[2])
        # p[0] += 10000000000000
        # p[1] += 10000000000000

        pA, pB = symbols('pA pB', integer=True, nonnegative=True)

        eq1 = Eq(a[0]*pA + b[0]*pB, p[0])
        eq2 = Eq(a[1]*pA + b[1]*pB, p[1])

        solution = solve((eq1, eq2), (pA, pB))

        if solution:
            total += solution[pA] * 3 + solution[pB]
            
    print("P2 = ", total)