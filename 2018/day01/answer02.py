#!/usr/bin/env python

import itertools

# Collate all the values into a list
data = [int(x) for x in open("input.txt").readlines()]
freq = 0
seen = set([0])

# Cycle over the data until we find the result
for val in itertools.cycle(data):
    freq += val
    if freq in seen:
        print(freq)
        break
    seen.add(freq)
