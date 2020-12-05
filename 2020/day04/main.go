// Package main -
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func check(holder []string) bool {
	//check if valid
	results := map[string]string{}
	for idx := range holder {
		for _, group := range strings.Split(holder[idx], " ") {
			field := strings.Split(group, ":")
			results[field[0]] = field[1]
		}
	}
	if len(results) == 8 {
		return true
	}
	if len(results) == 7 {
		if _, ok := results["cid"]; !ok {
			return true
		}
	}
	return false
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
			if check(holder) {
				valid++
			}
			holder = []string{}
		} else {
			holder = append(holder, text)
		}
	}
	if len(holder) != 0 {
		if check(holder) {
			valid++
		}

	}
	fmt.Printf("Result %d\n", valid)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
