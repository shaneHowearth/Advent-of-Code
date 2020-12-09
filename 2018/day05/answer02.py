#!/usr/bin/env python

import string
a = []
with open('input.txt') as f:
    for line in f:
        sca = list(line.strip())

# figure shit out
skip = False
b = []
for letter in list(string.ascii_lowercase):
    a = filter(lambda x: x.lower()!=letter, sca)


    while True:
        for i in  range(len(a)-1):
            if skip:
                skip = False
                continue
            if (a[i] == a[i+1].lower() and a[i].upper() == a[i+1]):
                skip = True
                continue
            if (a[i] == a[i+1].upper() and a[i].lower() == a[i+1]):
                skip = True
                continue
            b.append(a[i])
        b.append(a[-1:][0])
        if a != b:
            a = b
            b = []
        else:
            break

    print letter, len(b)
    b = []
