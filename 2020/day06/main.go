// Package main -
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

func check(holder []string) int {
	results := map[rune]bool{}
	counter := 0
	for group := range holder {
		for _, v := range holder[group] {
			if !results[v] {
				results[v] = true
				counter++
			}
		}
	}
	return counter
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	valid := 0
	var holder []string
	for scanner.Scan() {
		text := scanner.Text()
		if utf8.RuneCountInString(text) == 0 {
			valid += check(holder)
			holder = []string{}
		} else {
			holder = append(holder, text)
		}
	}

	if len(holder) != 0 {
		valid += check(holder)
	}

	fmt.Printf("Result %d\n", valid)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
