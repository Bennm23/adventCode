from utils.reader import read_to_list


lines = read_to_list("day2/input")

for line in lines:
    print(line)

def part1() -> int:

    total: int = 0
    for line in lines:
        split = line.split(": ")
        letter = split[0].split(" ")[1]
        count = [int(num) for num in split[0].split(" ")[0].split("-")]
        tgt = split[1]

        cnt = {}
        for c in tgt:
            if c in cnt:
                cnt[c] += 1
            else:
                cnt[c] = 1

        if letter in cnt:
            if cnt[letter] >= count[0] and cnt[letter] <= count[1]:
                total += 1
        print(cnt)
    return total

def part2() -> int:

    total: int = 0
    for line in lines:
        split = line.split(": ")
        letter = split[0].split(" ")[1]
        indices = [int(num) - 1 for num in split[0].split(" ")[0].split("-")]
        tgt = split[1]

        if tgt[indices[0]] == letter and tgt[indices[1]] != letter:
            total += 1

        if tgt[indices[1]] == letter and tgt[indices[0]] != letter:
            total += 1

    return total

#Part 1 = 474
print("Part 1 = ", part1())
#Part 2 = 745
print("Part 2 = ", part2())