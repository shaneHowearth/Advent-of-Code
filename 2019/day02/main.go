// Package main
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
	var tmp []string
	for scanner.Scan() {
		tmp = strings.Split(scanner.Text(), ",")
	}
	input := []int{}
	for i := range tmp {
		if t, terr := strconv.Atoi(tmp[i]); terr == nil {
			input = append(input, t)
		} else {
			log.Fatal("Unable to convert", terr)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Start", input)
	idx := 0
	for {
		switch input[idx] {
		case 1:
			input[input[input[idx+3]]] = input[input[idx+1]] + input[input[idx+2]]
		case 2:
			input[input[input[idx+3]]] = input[input[idx+1]] * input[input[idx+2]]
		case 99:
			goto end
		}
		idx += 4
	}
end:
	fmt.Println(input)
}
