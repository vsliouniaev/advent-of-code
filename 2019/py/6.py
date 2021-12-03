#!/usr/bin/env python3
filepath = "in/6"
f = open(filepath)
lines = f.readlines()
f.close()
data = dict(l.rstrip().split(")")[::-1] for l in lines)


def part1():
    i = len(data)
    for k in data:
        v = data[k]
        while v in data:
            v = data[v]
            i = i + 1
    return i


def part2():
    def path2Com(loc):
        p = data[loc]
        d, i = {}, 0
        while p in data:
            p = data[p]
            i += 1
            d[p] = i
        return d

    u = path2Com('YOU')
    s = path2Com('SAN')
    # Common planets between u and s |> min combined distance to planet
    return min([u[k] + s[k] for k in u.keys() & s.keys()])


print(part1())  # 417916
print(part2())  # 523
