#!/usr/bin/env python3

# Part A
with open('AoC_input') as AoCInput:
    maze = []
    for line in AoCInput:
        maze.append(int(line.rstrip()))

    position = 0
    steps = 0
    while True:
        new_position = position + maze[position]
        maze[position] += 1
        steps += 1
        position = new_position
        # Check if we have escaped the maze
        if position >= len(maze):
            break
    print(steps)

# Part B
with open('AoC_input') as AoCInput:
    maze = []
    for line in AoCInput:
        maze.append(int(line.rstrip()))

    position = 0
    steps = 0
    while True:
        new_position = position + maze[position]
        if maze[position] >=3:
            maze[position] -= 1
        else:    
            maze[position] += 1
        steps += 1
        position = new_position
        # Check if we have escaped the maze
        if position >= len(maze):
            break
    print(steps)
