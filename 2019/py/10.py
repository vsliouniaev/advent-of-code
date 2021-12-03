#!/usr/bin/env python3

import math

filepath = "in/10"
with open(filepath) as f:
    lines = f.readlines()

asteroids = set()
for y in range(len(lines)):
    for x in range(len(lines[y])):
        if lines[y][x] == '#':
            asteroids.add((x, y))


class asteroid:
    def __init__(self, p1, p2):
        x = p2[0] - p1[0]
        y = p2[1] - p1[1]
        self.angle = math.atan2(y, x)
        self.angle += (math.pi / 2)  # Up is zero
        self.angle %= (math.pi * 2)  # convert negative to pos
        self.pos = p2
        self.distance = abs(x) + abs(y)  # manhattan distance

    def __repr__(self):
        return f"p:{self.pos} a:{self.angle} d:{self.distance}"


def part1():
    max = 0
    locMax = (0, 0)
    for loc in asteroids:
        visible = set()
        for a in asteroids:
            if a != loc:
                visible.add(asteroid(loc, a).angle)
        if len(visible) > max:
            max = len(visible)
            locMax = loc

    return max, locMax


def part2():
    l = part1()[1]
    visible = {}
    for a in asteroids:
        if a != l:
            ast = asteroid(l, a)
            if ast.angle not in visible:
                visible[ast.angle] = []
            visible[ast.angle].append(ast)
            sorted(visible[ast.angle], key=lambda x: x.distance, reverse=True)

    i, v = 0, (None, None)
    while True:
        for a in sorted(visible):
            if len(visible[a]) > 0:
                v = visible[a].pop().pos
                i += 1
            if i == 200:
                return v[0] * 100 + v[1]


print(part1())  # 274
print(part2())  # 305
