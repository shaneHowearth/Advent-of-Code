#!/usr/bin/env python3

# Part A
with open('AoC_input') as AoCInput:
    for line in AoCInput:
        input_value = int(line.rstrip())
        test_input_values = [1, 12, 23, 1024]
        north_seed_diff = 3
        south_seed_diff = 7
        east_seed_diff = 1
        west_seed_diff = 5
        
        # Generate Axis
        # Everyone starts at 1
        north = east = west = south = 1

        step_counter = 0
        while True:
            # If any are bigger stop
            if north >= input_value or \
                south >= input_value or \
                east >= input_value or \
                west >= input_value:
                break
            north += north_seed_diff + 8 * step_counter
            south += south_seed_diff + 8 * step_counter
            east += east_seed_diff + 8 * step_counter
            west += west_seed_diff + 8 * step_counter
            step_counter += 1
        # Result is the number of layers out,
        # plus the distance to the nearest axis
        print(step_counter + min([abs(north - input_value),
            abs(east - input_value), abs(west - input_value), 
            abs(south - input_value)]))

# Part B
with open('AoC_input') as AoCInput:
    for line in AoCInput:
        input_value = int(line.rstrip())
        headline_value = 0
        row = 0
        col = 0
        spiral_values = {(row, col): 1}
        col += 1
        ADD_LEFT = 0
        ADD_ABOVE = 0
        ADD_BELOW = 0
        ADD_RIGHT = 0

        while headline_value <= input_value:
            # In a matrix, there are 8 possible adjacent values
            tmp_total = 0
            for row_adj in range(-1, 2):
                for col_adj in range(-1, 2):
                   if (row+row_adj, col+col_adj) in spiral_values:
                       tmp_total += spiral_values[(row+row_adj, col+col_adj)] 
            # We sum those for the new headline_value
            headline_value = tmp_total
            spiral_values[(row, col)] = \
                headline_value

            # We only change direction when there are only one adjacent places
            # already filled
            if (row, col-1) not in spiral_values and \
                (row, col+1) not in spiral_values:
                # No Left or right
                if (row-1, col) in spiral_values:
                    # Turning Left
                    ADD_LEFT = -1
                    ADD_ABOVE = 0
                    ADD_BELOW = 0
                    ADD_RIGHT = 0

                elif (row+1, col) in spiral_values:
                    # Turning Right
                    ADD_LEFT = 0
                    ADD_ABOVE = 0
                    ADD_BELOW = 0
                    ADD_RIGHT = 1

            elif (row-1, col) not in spiral_values and \
                (row+1, col) not in spiral_values:
                # No Up or Down
                if (row, col+1) in spiral_values:
                    # Turning down
                    ADD_LEFT = 0
                    ADD_ABOVE = 0
                    ADD_BELOW = -1
                    ADD_RIGHT = 0

                elif (row, col-1) in spiral_values:
                    # Turning Upwards
                    ADD_LEFT = 0
                    ADD_ABOVE = 1
                    ADD_BELOW = 0
                    ADD_RIGHT = 0



            # Ready for the next loop
            row += ADD_BELOW + ADD_ABOVE
            col += ADD_LEFT + ADD_RIGHT

        print( headline_value)
