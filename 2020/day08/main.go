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
	seen := map[int]struct{}{}
	commands := []command{}
	for scanner.Scan() {
		tmp := scanner.Text()
		// split on space
		words := strings.Split(tmp, " ")
		num, _ := strconv.Atoi(words[1])
		c := command{instr: words[0], val: num}
		commands = append(commands, c)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(commands)
	position := 0
	counter := 0
	for {
		if _, ok := seen[position]; !ok {
			seen[position] = struct{}{}
			//execute
			switch commands[position].instr {
			case "acc":
				counter += commands[position].val
				position++
			case "jmp":
				position += commands[position].val
			case "nop":
				position++
			}
		} else {
			break
		}
	}
	fmt.Println(counter)
}
