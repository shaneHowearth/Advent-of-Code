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

func canHold(s string, set []string) bool {
	for _, v := range set {
		if strings.Contains(v, s) {
			return true
		}
	}
	return false
}

var lines = map[string][]string{}

func getChildCount(bags []string) int {

	count := 1
	for _, bag := range bags {
		words := strings.Split(strings.TrimSpace(bag), " ")
		if len(words) <= 2 {
			break
		}
		newBag := strings.Split(strings.Join(words[1:3], " "), ",")
		num, _ := strconv.Atoi(words[0])
		bagName := strings.Join(newBag, " ")
		if !strings.Contains(bagName, "other bags.") {
			count += num * getChildCount(lines[bagName])
		}
	}
	return count
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	seen := map[string]struct{}{}
	for scanner.Scan() {
		tmp := scanner.Text()
		// split on space
		words := strings.Split(tmp, " ")
		bags := strings.Split(strings.Join(words[4:], " "), ",")
		lines[strings.Join(words[:2], " ")] = bags
	}

	seen["shiny gold"] = struct{}{}
	fmt.Println(getChildCount(lines["shiny gold"]) - 1)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
