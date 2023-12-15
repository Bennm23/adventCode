from utils import runAndPrintMicros, filePath, runAndPrintMillis

from functools import cache

@cache
def count(cfg, nums):

    if cfg == '':
        return 1 if nums == () else 0
    if nums == ():
        return 0 if '#' in cfg else 1

    total = 0

    if cfg[0] in '.?':
        total += count(cfg[1:], nums)

    if cfg[0] in '#?':

        if nums[0] <= len(cfg) and \
            '.' not in cfg[:nums[0]] and\
            (nums[0] == len(cfg) or cfg[nums[0]] != '#'):

            total += count(cfg[nums[0] + 1:], nums[1:])

    return total


def solve():

    p1, p2 = 0, 0
    for line in open(filePath('day12.txt')):
        cfg,nums = line.split()
        nums = tuple(map(int, nums.split(',')))
        p1 += count(cfg, nums)

        cfg = '?'.join([cfg]*5)
        nums = nums * 5

        p2 += count(cfg, nums)


    print('Part 1 = ', p1)#7110
    print('Part 2 = ', p2)#1566786613613


runAndPrintMillis(solve)

#With @cache 187.6
#With homegrown cache 208
