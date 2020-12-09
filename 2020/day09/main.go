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
		fmt.Println(input[position])
		break
	}
}
