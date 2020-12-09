#!/usr/bin/env python

val = 0
with open('input.txt') as f:
    for line in f:
        if line.startswith('-'):
            val -= int(line[1:])
        if line.startswith('+'):
            val += int(line[1:])
        if line.startswith('*'):
            val *= int(line[1:])
        if line.startswith('/'):
            val /= int(line[1:])

print(val)
