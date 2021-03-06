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
		var interim int
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
		if string(input[1][bottom]) == rules[1] {
			interim++
		}
		if string(input[1][top]) == rules[1] {
			interim++
		}
		if interim == 1 {
			result++
		}
	}
	fmt.Printf("Result %d\n", result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
