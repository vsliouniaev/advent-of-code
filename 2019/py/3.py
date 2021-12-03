#!/usr/bin/env python3
import math

filepath = "in/3"
f = open(filepath)
lines = f.readlines()
f.close()

distances = []
for vectors in [w.rstrip() for w in lines]:
    x = 0
    y = 0
    d = 0
    posDist = {}
    distances.append(posDist)
    for v in [v for v in vectors.split(",")]:
        h = v[0]
        l = int(v[1:])
        for _ in range(l):
            if h == "R":
                x += 1
            elif h == "L":
                x -= 1
            elif h == "D":
                y -= 1
            elif h == "U":
                y += 1
            else:
                raise Exception("Unknown heading '{}'".format(h))
            d += 1
            posDist[(x, y)] = d

intersections = distances[0].keys() & distances[1].keys()


def part1():
    manhattan = lambda t: abs(t[0]) + abs(t[1])
    return min(map(manhattan, intersections))


def part2():
    distance = lambda t: distances[0][t] + distances[1][t]
    return min(map(distance, intersections))


print(part1())  # 1983
print(part2())  # 107754
