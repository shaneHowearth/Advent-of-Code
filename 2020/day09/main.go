// Package main -
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type command struct {
	instr string
	val   int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := []int{}
	for scanner.Scan() {
		tmp := scanner.Text()
		num, _ := strconv.Atoi(tmp)
		input = append(input, num)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	position := 25
	for {
	next:
		window := map[int]bool{}
		for i := position - 25; i < position; i++ {
			window[input[i]] = true
		}
		for i := 1; i <= 25; i++ {
			if window[input[position]-input[position-i]] {
				position++
				goto next
			}
		}
		// Number that isn't the sum of any pair of previous 25 vals
		fmt.Println(input[position])
		break
	}

	bottom := 0
	top := 1
	for top <= len(input) {
		a := 0
		for _, v := range input[bottom:top] {
			a += v
		}
		if a == input[position] {
			largest := 0
			smallest := input[position]
			for _, v := range input[bottom:top] {
				if v < smallest {
					smallest = v
				}
				if v > largest {
					largest = v
				}
			}
			// Smallest and largest vals in set of contiguous vals that add up
			// to be the first number that was found
			fmt.Println(largest + smallest)
			break
		}
		if a < input[position] {
			top++
		} else {
			bottom++
		}
	}
}
