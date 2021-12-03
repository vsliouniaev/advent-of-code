#!/usr/bin/env python3
from collections import deque
import enum


class _ioState(enum.Enum):
    OutputProduced = 0
    InputRequired = 1
    Halted = 2


class IntcodeComputer:
    def __init__(self, program, name=None, log=False):
        self._name = name
        self._memory = [0] * 10000
        self._memory[:len(program)] = program
        self._inputBuffer = deque()
        self._outputBuffer = []
        self._state = None
        self._instructionPointer = 0
        self._relativeBase = 0
        self._logEnabled = log

    def WriteInput(self, val):
        # self._log("inp {}".format(val))
        if type(val) is deque: val = list(val)
        if type(val) is int: val = [val]
        for v in list(val):
            self._inputBuffer.append(int(v))
        return self

    def GetOutput(self):
        o = self._outputBuffer.copy()
        self._outputBuffer = []
        return o

    def RunToOutput(self):
        return self._run(_ioState.OutputProduced)

    def RunToInput(self):
        return self._run(_ioState.InputRequired)

    def RunToHalt(self):
        return self._run()

    def IsHalted(self):
        return self._state == _ioState.Halted

    def _run(self, pauseFor=_ioState.Halted):
        if self._state == _ioState.Halted:
            raise Exception(f"{self._name} halted")
        self._state = None
        while self._state not in [_ioState.Halted, pauseFor]:
            i = self._instruction(self._memory[self._instructionPointer])
            self._instructionPointer, self._state = self._ops[i.opcode](
                self, self._instructionPointer, i.modes)
        return self

    def _log(self, l):
        if not self._logEnabled:
            return
        if self._name == None:
            print(l)
        else:
            print("{}: {}".format(self._name, l))

    def _getParams(self, n, i, modes):
        return [self._getWithMode(i + x + 1, modes[x]) for x in range(n)]

    def _getWithMode(self, i, mode):
        return self._operand(i, self._memory, mode, self._relativeBase)

    # 01 - add
    # adds together numbers read from two positions and stores the result in a third position.
    # The three integers immediately after the opcode tell you these three positions - the first two indicate the positions
    # from which you should read the input values, and the third indicates the position at which the output should be stored.
    def _add(self, i, modes):
        p = self._getParams(3, i, modes)
        v = p[0].value + p[1].value
        self._memory[p[2].address] = v
        return i + 4, None

    # 02 - multiply
    # works exactly like opcode 1, except it multiplies the two inputs instead of adding them
    def _mul(self, i, modes):
        p = self._getParams(3, i, modes)
        v = p[0].value * p[1].value
        self._memory[p[2].address] = v
        return i + 4, None

    # 03 - input
    # takes a single integer as input and saves it to the position given by its only parameter
    def _inp(self, i, modes):
        if len(self._inputBuffer) == 0:
            return i, _ioState.InputRequired
        v = self._inputBuffer.popleft()
        p = self._getParams(1, i, modes)
        self._memory[p[0].address] = v
        return i + 2, None

    # 04 - output
    # outputs the value of its only parameter. For example, the instruction 4,50 would output the value at address 50.
    def _out(self, i, modes):
        p = self._getParams(1, i, modes)
        self._log("read out {}".format(p[0]))
        self._outputBuffer.append(p[0].value)
        return i + 2, _ioState.OutputProduced

    # 05 - jump-if-true:
    # if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
    def _jit(self, i, modes):
        p = self._getParams(2, i, modes)
        if p[0].value != 0:
            return p[1].value, None
        return i + 3, None

    # 06 - jump-if-false:
    # if the first parameter is zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing
    def _jif(self, i, modes):
        p = self._getParams(2, i, modes)
        if p[0].value == 0:
            return p[1].value, None
        return i + 3, None

    ## 07 - less than:
    # if the first parameter is less than the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0
    def _les(self, i, modes):
        p = self._getParams(3, i, modes)
        # Parameters that an instruction writes to will never be in immediate mode.
        v = 1 if p[0].value < p[1].value else 0
        self._memory[p[2].address] = v
        return i + 4, None

    ## 08 - equals:
    # if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0
    def _equ(self, i, modes):
        p = self._getParams(3, i, modes)
        v = 1 if p[0].value == p[1].value else 0
        self._memory[p[2].address] = v
        return i + 4, None

    ## 09 - relative-base-adjust
    ## adjusts the relative base by the value of its only parameter. The relative base increases (or decreases, if the value is negative)
    # by the value of the parameter
    def _rba(self, i, modes):
        p = self._getParams(1, i, modes)
        self._relativeBase += p[0].value
        self._log("set relative base to {}".format(self._relativeBase))
        return i + 2, None

    # 99 - halt
    # means that the program is finished and should immediately halt
    def _hlt(self, i, modes):
        return None, _ioState.Halted

    _ops = {
        1: _add,
        2: _mul,
        3: _inp,
        4: _out,
        5: _jit,
        6: _jif,
        7: _les,
        8: _equ,
        9: _rba,
        99: _hlt
    }

    class _instruction:
        def __init__(self, num):
            s = str(num)
            self.opcode = int(s[-2:])
            m = [int(x) for x in s[:-2][::-1]]
            self.modes = m + [0] * (5 - len(m))  # pad to 5

    class _operand:
        def __init__(self, i, memory, mode, relBase):
            if mode == 0:
                self.address = memory[i]
                self.str = "@"
            if mode == 1:
                self.address = i
            if mode == 2:
                self.address = memory[i] + relBase
                self.str = "@+({})".format(relBase)
            self.value = memory[self.address]