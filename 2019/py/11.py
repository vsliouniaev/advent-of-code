#!/usr/bin/env python3
from IntcodeComputer import IntcodeComputer

filepath = "in/11"
f = open(filepath)
dat = f.read().replace("\n", "")
f.close()
data = [int(x) for x in dat.split(",")]


class agent:
    def __init__(self):
        #               up     right   down      left
        self._dirs = [(0, 1), (1, 0), (0, -1), (-1, 0)]
        self._dir = 0  # Index into dirs array
        self._loc = (0, 0)

    def Move(self, inp):
        if inp == 0:  #left
            self._dir = (self._dir - 1) % len(self._dirs)
        elif inp == 1:  #right
            self._dir = (self._dir + 1) % len(self._dirs)
        else:
            raise Exception(f"Unknown direction {inp}")

        d = self._dirs[self._dir]
        # Add vector from dirs array to position
        self._loc = (self._loc[0] + d[0], self._loc[1] + d[1])
        return self._loc


def part1():
    vm = IntcodeComputer(data)
    panels = {}
    robot = agent()
    position = (0, 0)
    while True:
        camera = 0
        if position in panels:
            camera = panels[position]
        if vm.WriteInput(camera).RunToOutput().IsHalted():
            return len(panels)
        color = vm.GetOutput()[-1]
        if color != 0 and color != 1:
            raise Exception(f"color {color}")
        panels[position] = color
        vm.RunToOutput()
        direction = vm.GetOutput()[-1]
        position = robot.Move(direction)


def part2():
    vm = IntcodeComputer(data)
    panels = {(0, 0): 1}
    robot = agent()
    position = (0, 0)
    while True:
        camera = 0
        if position in panels:
            camera = panels[position]
        if vm.WriteInput(camera).RunToOutput().IsHalted():
            break
        color = vm.GetOutput()[-1]
        if color != 0 and color != 1:
            raise Exception(f"color {color}")
        panels[position] = color
        vm.RunToOutput()
        direction = vm.GetOutput()[-1]
        position = robot.Move(direction)

    # Guess the ranges
    out = ""
    for y in range(0, -7, -1):
        out += "\n"
        for x in range(-1, 45):
            if (x, y) in panels and panels[(x, y)] == 1:
                out += "H"
            else:
                out += " "
    return out


print(part1())  # 2041
print(part2())  # ZRZPKEZR
