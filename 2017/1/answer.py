#!/usr/bin/env python3

# Part A
with open('AoC_input') as AoCInput:
    counter = 0
    for line in AoCInput:
        line = line.rstrip()
        for char in enumerate(line):
            if (char[0] < len(line) - 1 and
                line[int(char[0])] == line[int(char[0])+1]):
                counter += (int(char[1]))
        if int(line[0]) == int(line[-1]):
            # handle wraparound
            counter += int(line[-1])
    print(counter)
