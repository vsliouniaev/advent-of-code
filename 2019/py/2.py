#!/usr/bin/env python3
import math

filepath = "in/2"
f = open(filepath)
lines = f.readlines()
f.close()
buf = [int(x) for x in lines[0].split(",")]


def add(buf, i):
    buf[buf[i + 3]] = buf[buf[i + 1]] + buf[buf[i + 2]]
    return i + 4


def mul(buf, i):
    buf[buf[i + 3]] = buf[buf[i + 1]] * buf[buf[i + 2]]
    return i + 4


def hlt(buf, i):
    return None


ops = {1: add, 2: mul, 99: hlt}


def part1(data, n, v):
    buf = data.copy()
    buf[1] = n
    buf[2] = v
    i = 0
    while i != None:
        i = ops[buf[i]](buf, i)

    return buf[0]


def part2(data):
    for n in range(99):
        for v in range(99):
            if part1(data, n, v) == 19690720:
                return (100 * n) + v


print(part1(buf, 12, 2))  # 3306701
print(part2(buf))  # 7621
