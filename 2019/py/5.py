#!/usr/bin/env python3
import math

filepath = "in/5"
f = open(filepath)
dat = f.read().replace("\n", "")
f.close()
data = [int(x) for x in dat.split(",")]

inputBuffer = []
outputBuffer = []


def getParams(n, mem, modes, i):
    return [getWithMode(mem, modes[x], i + x + 1) for x in range(n)]


def getWithMode(mem, mode, i):
    if mode == 0:
        return mem[mem[i]]
    if mode == 1:
        return mem[i]


# 01
def add(mem, i, modes):
    p = getParams(2, mem, modes, i)
    mem[mem[i + 3]] = p[0] + p[1]
    return i + 4


# 02
def mul(mem, i, modes):
    p = getParams(2, mem, modes, i)
    mem[mem[i + 3]] = p[0] * p[1]
    return i + 4


# 03
def inp(mem, i, modes):
    mem[mem[i + 1]] = inputBuffer.pop()
    return i + 2


# 04
def out(mem, i, modes):
    p = getParams(1, mem, modes, i)
    outputBuffer.append(p[0])
    # outputBuffer.append(mem[mem[i + 1]])
    return i + 2


# 05 - jump-if-true:
# if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
def jit(mem, i, modes):
    p = getParams(2, mem, modes, i)
    if p[0] != 0:
        return p[1]
    return i + 3


# 06 - jump-if-false:
# if the first parameter is zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing
def jif(mem, i, modes):
    p = getParams(2, mem, modes, i)
    if p[0] == 0:
        return p[1]
    return i + 3


## 07 - less than:
# if the first parameter is less than the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0
def les(mem, i, modes):
    p = getParams(3, mem, modes, i)
    # Parameters that an instruction writes to will never be in immediate mode.
    mem[mem[i + 3]] = 1 if p[0] < p[1] else 0
    return i + 4


## 08 - equals:
# if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0
def equ(mem, i, modes):
    p = getParams(3, mem, modes, i)
    # Parameters that an instruction writes to will never be in immediate mode.
    mem[mem[i + 3]] = 1 if p[0] == p[1] else 0
    return i + 4


# 99
def hlt(mem, i, modes):
    return None


ops = {1: add, 2: mul, 3: inp, 4: out, 5: jit, 6: jif, 7: les, 8: equ, 99: hlt}


class instruction:
    def __init__(self, num):
        s = str(num)
        self.opcode = int(s[-2:])
        m = [int(x) for x in s[:-2][::-1]]
        self.modes = m + [0] * (5 - len(m))  # pad to 5


def run(data):
    mem = data.copy()
    x = 0
    while x != None:
        # print("{:03}".format(x))
        i = instruction(mem[x])
        x = ops[i.opcode](mem, x, i.modes)

    return outputBuffer


inputBuffer.append(1)
print(run(data)[-1:])  # 6745903
inputBuffer.append(5)
print(run(data)[-1:])  # 9168267