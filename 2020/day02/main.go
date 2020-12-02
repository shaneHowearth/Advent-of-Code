// Package main -
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result int
	for scanner.Scan() {
		// split on :
		input := strings.Split(scanner.Text(), ":")
		// split on space
		rules := strings.Split(input[0], " ")
		// split on dash (and convert to ints)
		counts := strings.Split(rules[0], "-")
		bottom, berr := strconv.Atoi(counts[0])
		if berr != nil {
			log.Printf("bottom didn't convert with error %v", berr)
		}

		top, terr := strconv.Atoi(counts[1])
		if terr != nil {
			log.Printf("top didn't convert with error %v", terr)
		}
		actualCount := strings.Count(input[1], rules[1])
		if actualCount >= bottom && actualCount <= top {
			result++
		}
	}
	fmt.Printf("Result %d\n", result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
