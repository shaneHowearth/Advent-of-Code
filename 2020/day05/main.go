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
	var result int
	seats := [128][8]int{}
	seatMap := map[int]struct{}{}
	for scanner.Scan() {
		row := []int{0, 127}
		seat := []int{0, 7}
		for _, v := range scanner.Text() {
			if v == 'F' {
				row[1] = row[0] + ((row[1] - row[0]) / 2)
			}
			if v == 'B' {
				row[0] = row[1] - ((row[1] - row[0]) / 2)
			}
			if v == 'R' {
				seat[0] = seat[1] - ((seat[1] - seat[0]) / 2)
			}
			if v == 'L' {
				seat[1] = seat[0] + ((seat[1] - seat[0]) / 2)
			}
		}
		seatID := row[0]*8 + seat[0]
		seats[row[0]][seat[0]] = seatID
		seatMap[seatID] = struct{}{}
		if seatID > result {
			result = seatID
		}
	}
	for r := range seats {
		for c := range seats[r] {
			if r > 0 && r < 127 && seats[r][c] == 0 {
				id := r*8 + c
				if _, ok := seatMap[id-1]; ok {
					if _, ok := seatMap[id+1]; ok {
						fmt.Printf("R: %d, C: %d ID: %d\n", r, c, id)
					}
				}
			}
		}
	}
	fmt.Printf("Result %d\n", result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
