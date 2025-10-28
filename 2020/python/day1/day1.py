from utils.reader import read_to_int_list

nums = read_to_int_list("day1/input")

def part1(nums: list[int]) -> int:

    for i in range(0, len(nums)):
        for j in range(i+1, len(nums)):

            if nums[i] + nums[j] == 2020:
                return nums[i] * nums[j]

    return -1

def part2(nums: list[int]) -> int:

    for i in range(0, len(nums)):
        for j in range(i+1, len(nums)):
            for k in range(j+1, len(nums)):

                if nums[i] + nums[j] + nums[k] == 2020:
                    return nums[i] * nums[j] * nums[k]

    return -1

p1 = part1(nums)

#Part 1 =  1019571
print("Part 1 = ", p1)
#Part 2 =  100655544
print("Part 2 = ", part2(nums))