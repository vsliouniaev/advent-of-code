#!/usr/bin/env python3
import math
import networkx as nx

filepath = "in/6"
f = open(filepath)
lines = f.readlines()
f.close()

G = nx.Graph()
G.add_edges_from([l.rstrip().split(")")[::-1] for l in lines])


def part1():
    return sum([nx.shortest_path_length(G, n, 'COM') for n in G.nodes])


def part2():
    return nx.shortest_path_length(G, 'SAN', 'YOU') - 2


print(part1())  # 417916
print(part2())  # 523
