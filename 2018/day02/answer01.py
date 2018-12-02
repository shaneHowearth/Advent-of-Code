#!/usr/bin/env python

import collections
from itertools import groupby
from operator import mul

val = []
with open('input.txt') as f:
    for line in f:
        a = list(line.strip())
        a.sort()
        b=collections.Counter(a)
        c = set(b.values())
        val.extend(c)

print(reduce(mul, collections.Counter([v for v in val if v>1]).values(), 1))
