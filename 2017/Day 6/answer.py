#!/usr/bin/env python3

# Part A
with open('AoC_input') as AoCInput:
    for line in AoCInput:
        previous_instances = {}
        input_value = line.rstrip().split()
        input_values = [int(x) for x in input_value]
        previous_instances[tuple(input_values)] = 1

        # Redistribute
        while True:
            max_bank = input_values.index(max(input_values))
            store = input_values[max_bank]
            input_values[max_bank] = 0
            for i in range(1, store+1):
                input_values[(max_bank+i)%len(input_values)] +=1

            if tuple(input_values) not in previous_instances:
                previous_instances[tuple(input_values)] = 1
            else:
                break

        print(len(previous_instances))

# Part B
with open('AoC_input') as AoCInput:
    for line in AoCInput:
        previous_instances = {}
        input_value = line.rstrip().split()
        input_values = [int(x) for x in input_value]
        previous_instances[tuple(input_values)] = 1
        been_here = False
        new_loop_instances = {}

        # Redistribute
        while True:
            max_bank = input_values.index(max(input_values))
            store = input_values[max_bank]
            input_values[max_bank] = 0
            for i in range(1, store+1):
                input_values[(max_bank+i)%len(input_values)] +=1

            if tuple(input_values) not in previous_instances:
                previous_instances[tuple(input_values)] = 1
            else:
                if not been_here:
                    new_loop_instances[tuple(input_values)] = 1
                    been_here = True
                else:
                    if tuple(input_values) not in new_loop_instances:
                        new_loop_instances[tuple(input_values)] = 1
                    else:
                        break

        print(len(new_loop_instances))
