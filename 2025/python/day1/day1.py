
def read_to_input():
    sample = []
    with open("input") as f:
        sample = [line.strip() for line in f.readlines()]
    return sample


def p1():
    data = read_to_input()
    total = 0
    dial = 50
    for line in data:
        dir = line[0]
        amt = int(line[1:])
        if dir == "L":
            dial -= amt
            dial %= 100
        if dir == "R":
            dial += amt
            dial %= 100

        # print("The dial is rotated ", line, " to point at ", dial)

        if dial == 0:
            total += 1
    

    return total
def p2(p1_score):
    data = read_to_input()
    in_rotation_total = 0
    dial = 50
    for line in data:
        dir = line[0]
        amt = int(line[1:])

        if dir == "L": 
            inc = -1
        if dir == "R":
            inc = 1

        for _ in range(amt - 1):
            dial += inc

            if dial % 100 == 0:
                print(f"Rotation {line} hit 0")
                in_rotation_total += 1

        dial += inc
        dial %= 100


    return in_rotation_total + p1_score


# Part 1: 1076
part1 = p1()
# Part 2: 6379
part2 = p2(part1)

print("Part 1:", part1)
print("Part 2:", part2)