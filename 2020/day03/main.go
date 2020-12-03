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
	var start int
	var row int
	for scanner.Scan() {
		row++
		if row == 1 {
			continue
		}
		start += 3
		text := scanner.Text()
		if string(text[start%len(text)]) == "#" {
			result++
		}
	}
	fmt.Printf("Result %d\n", result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
