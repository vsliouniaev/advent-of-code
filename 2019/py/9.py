#!/usr/bin/env python3
from IntcodeComputer import IntcodeComputer

filepath = "in/9"
f = open(filepath)
dat = f.read().replace("\n", "")
f.close()
data = [int(x) for x in dat.split(",")]


def part1():
    return IntcodeComputer(data).WriteInput(1).RunToHalt().GetOutput()


def part2():
    return IntcodeComputer(data).WriteInput(2).RunToHalt().GetOutput()


print(part1())  # 4234906522
print(part2())  # 60962