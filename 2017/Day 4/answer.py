#!/usr/bin/env python3

# Part A
with open('AoC_input') as AoCInput:
    valid_count = 0
    for line in AoCInput:
        input_value = line.rstrip()
        word_list = input_value.split()
        word_list.sort()

        valid = True
        for word_pos in range(len(word_list)-1):
            if word_list[word_pos] == word_list[word_pos+1]:
                valid = False
                break
        if valid:
            valid_count += 1

    print(valid_count)

# Part B
with open('AoC_input') as AoCInput:
    valid_count = 0
    for line in AoCInput:
        input_value = line.rstrip()
        word_list = input_value.split()
        word_list = [''.join(sorted(word)) for word in word_list]
        word_list.sort()

        valid = True
        for word_pos in range(len(word_list)-1):
            if word_list[word_pos] == word_list[word_pos+1]:
                valid = False
                break
        if valid:
            valid_count += 1

    print(valid_count)
