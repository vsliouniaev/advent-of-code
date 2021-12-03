#!/usr/bin/env python3
from itertools import groupby
import math

inputRange = range(353096, 843212)


def numDigits(n):
    return math.ceil(math.log(n, 10))


def split(n):
    return [(n // (10**i)) % 10 for i in range(numDigits(n) - 1, -1, -1)]


def part1():
    def chkpass(num):
        d = split(num)
        for i in range(1, numDigits(num)):
            if d[i] < d[i - 1]:
                return 0
        return any([len(list(v)) > 1 for _, v in groupby(d)])

    return sum([chkpass(n) for n in inputRange])


def part2():
    def chkpass(num):
        d = split(num)
        for i in range(1, numDigits(num)):
            if d[i] < d[i - 1]:
                return 0
        return any([len(list(v)) == 2 for _, v in groupby(d)])

    return sum([chkpass(n) for n in inputRange])


print(part1())  # 579
print(part2())  # 358
