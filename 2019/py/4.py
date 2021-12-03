#!/usr/bin/env python3
import math

inputRange = range(353096, 843212)


def numDigits(n):
    return 6


def split(n):
    return [
        (n // 100000) % 10,
        (n // 10000) % 10,
        (n // 1000) % 10,
        (n // 100) % 10,
        (n // 10) % 10,
        n % 10,
    ]


def part1():
    def chkpass(num):
        d = split(num)
        consecutive = 0
        for i in range(1, numDigits(num)):
            if d[i] == d[i - 1]:
                consecutive = 1
            if d[i] < d[i - 1]:
                return 0
        return consecutive

    return sum([chkpass(n) for n in inputRange])


def part2():
    def chkpass(num):
        d = split(num)
        consFound = False
        consecutive = 1
        for i in range(1, numDigits(num)):
            if not consFound:
                if d[i] == d[i - 1]:
                    consecutive += 1
                else:
                    consFound = True if consecutive == 2 else False
                    consecutive = 1
            if d[i] < d[i - 1]:
                return 0
        return 1 if consecutive == 2 or consFound else 0

    return sum([chkpass(n) for n in inputRange])


print(part1())  # 579
print(part2())  # 358