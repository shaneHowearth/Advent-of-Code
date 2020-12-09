#!/usr/bin/env python
import numpy

b = {}
with open('input.txt') as f:
# with open('example.txt') as f:
    for line in f:
        a = line.strip().split(' ')
        b[a[0]] = [[int(i) for i in a[2][:-1].split(',')], [int(i) for i in a[3].split('x')]]

maxX = max(v[0][0] + v[1][0] for v in b.values())
maxY = max(v[0][1] + v[1][1] for v in b.values())
# create blank matrix
mx = numpy.zeros((maxX, maxY), dtype=int)

# apply the cutouts
for k in b:
    for i in range(b[k][1][0]):
        for j in range(b[k][1][1]):
            mx[b[k][0][0]+i,b[k][0][1]+j]+=1
# look for an unmolested cutout
for k in b:
    found = True
    for i in range(b[k][1][0]):
        for j in range(b[k][1][1]):
            if mx[b[k][0][0]+i,b[k][0][1]+j]!=1:
                found = False
    if found:
        print(k)
