#!/bin/bash

# O(n**2)
while read l
    do
        x=`agrep -1 -c $l input.txt`
        if [ $x -eq 2 ]
            then
                echo $l
        fi
    done < input.txt
