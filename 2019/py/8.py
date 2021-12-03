from itertools import groupby

filepath = "in/8"
f = open(filepath)
dat = f.read().replace("\n", "")
f.close()
data = [int(x) for x in dat]

w = 25
h = 6


def getLayers():
    lsize = w * h
    numlayers = len(data) // lsize
    layers = {}
    for lidx in range(0, numlayers):
        pixels = []
        for i in range(0, lsize):
            pixels.append(data[lidx * lsize + i])
        layers[lidx] = pixels
    return layers


layers = getLayers()


def part1():
    stats = {k: v.count(0) for (k, v) in layers.items()}
    target = min(stats, key=stats.get)
    return layers[target].count(1) * layers[target].count(2)


def part2():
    image = [2] * w * h
    for _, l in layers.items():
        for i in range(0, len(l)):
            if image[i] == 2:
                image[i] = l[i]

    out = ""
    for y in range(0, h):
        for x in range(0, w):
            p = image[(y * w) + x]
            if p == 0:
                out += " "
            else:
                out += "🎄"
        out += "\n"
    return out.strip()


print(part1())  # 1215
print(part2())  # LHCPH
