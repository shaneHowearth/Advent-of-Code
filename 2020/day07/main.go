// Package main -
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func canHold(s string, set []string) bool {
	for _, v := range set {
		if strings.Contains(v, s) {
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
	data := map[string][]string{}
	for scanner.Scan() {
		tmp := scanner.Text()
		// split on space
		words := strings.Split(tmp, " ")
		bags := strings.Split(strings.Join(words[4:], " "), ",")
		data[strings.Join(words[:2], " ")] = bags
	}

	possibles := map[string]struct{}{}
	seen := map[string]struct{}{}
	for k := range data {
		// does this k have shiny gold bags
		if canHold("shiny gold", data[k]) {
			possibles[k] = struct{}{}
			seen[k] = struct{}{}
		}
	}

	fmt.Println(len(possibles))
	newPossibles := map[string]struct{}{}
	for {
		changed := false
		for k := range possibles {
			for j := range data {
				if canHold(k, data[j]) {
					if _, ok := seen[j]; !ok {
						changed = true
						seen[j] = struct{}{}
					}
					newPossibles[j] = struct{}{}
				}
			}
		}
		possibles = newPossibles
		if !changed {
			break
		}
	}
	fmt.Println(len(seen))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
