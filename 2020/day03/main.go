// Package main -
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	results := map[int]int{}
	var row int
	slopes := []int{1, 3, 5, 7, 1}
	for scanner.Scan() {
		if row == 0 {
			row++
			continue
		}
		text := scanner.Text()
		for idx, s := range slopes {
			start := row * s
			if idx == 4 {
				if row%2 != 0 {
					continue
				}
				start = row * s / 2
			}
			if string(text[start%len(text)]) == "#" {
				(results[idx])++
			}

		}
		row++
	}
	result := 1
	for k := range results {
		result *= results[k]
	}
	fmt.Printf("Result %d\n", result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
