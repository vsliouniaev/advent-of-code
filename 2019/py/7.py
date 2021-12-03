#!/usr/bin/env python3
import itertools
from collections import deque
from IntcodeComputer import IntcodeComputer

filepath = "in/7"
f = open(filepath)
dat = f.read().replace("\n", "")
f.close()
data = [int(x) for x in dat.split(",")]

# Assumes one output, one input only
# def checkPhases(phases):
#     amps = [IntcodeComputer(data, i) for i in range(0, len(phases))]
#     io = 0
#     for a in range(len(amps)):
#         amps[a].WriteInput(phases[a])
#     while True:
#         for a in range(len(amps)):
#             amps[a].WriteInput(io).RunToOutput()
#             if amps[a].IsHalted():
#                 return io
#             else:
#                 io = amps[a].GetOutput()


# Can deal with multiple outputs from one stage into the next
def checkPhases(phases):
    amps = [IntcodeComputer(data, i) for i in range(0, len(phases))]
    dataLine = deque([0])
    for a in range(len(amps)):
        amps[a].WriteInput(phases[a])
    while True:
        for a in range(len(amps)):
            # Checking at the beginning ensures that we only stop when all amps are done
            # Doing this a the end quits when the first amp is done
            if amps[a].IsHalted():
                return list(dataLine)
            dataLine = deque(
                amps[a].WriteInput(dataLine).RunToInput().GetOutput())


def part1():
    phaseses = {}
    for phases in itertools.permutations(range(0, 5)):
        amps = [IntcodeComputer(data, i) for i in range(0, 5)]
        io = 0
        for a in range(len(amps)):
            amps[a].WriteInput(phases[a]).WriteInput(io).RunToOutput()
            io = amps[a].GetOutput()
        phaseses[phases] = io
    mx = max(phaseses, key=phaseses.get)
    return "{} : {}".format(phaseses[mx], mx)


def part2():
    phaseses = {}
    for phases in itertools.permutations(range(5, 10)):
        phaseses[phases] = checkPhases(phases)

    mx = max(phaseses, key=phaseses.get)
    return "{} : {}".format(phaseses[mx], mx)


print(part1())  # 118936 : (2, 1, 4, 3, 0)
print(part2())  # 57660948 : (9, 7, 6, 5, 8)
