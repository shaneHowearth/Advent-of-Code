#!/usr/bin/env python3

# Part A
with open('AoC_input') as AoCInput:
    counter = 0
    for line in AoCInput:
        line = line.split()
        line = [int(v) for v in line]
        line.sort()
        counter += line[-1] -line[0]

    print(counter)

# Part B
with open('AoC_input') as AoCInput:
    counter = 0
    for line in AoCInput:
        line = line.split()
        line = [int(v) for v in line]
        line.sort()
        for val in enumerate(line):
            for i in range(val[0] + 1, len(line)):
                if line[i] % val[1] == 0:
                    counter += line[i] / val[1]

    print(counter)
