#!/usr/bin/env python3
import math

filepath = "in/1"
f = open(filepath)
lines = f.readlines()
f.close()
masses = [int(x) for x in lines]


def fuelRequired(m):
    return int(m) // 3 - 2


def fuelRequiredRec(mass):
    f = fuelRequired(mass)
    return 0 if f < 1 else f + fuelRequiredRec(f)


def part1():
    return sum(map(fuelRequired, masses))


def part2():
    fship = 0
    for m in masses:
        fship += fuelRequiredRec(m)
    return fship


print(part1())  # 3249140
print(part2())  # 4870838
