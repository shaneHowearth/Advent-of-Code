// Package main -
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("vim-go")
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := map[int]struct{}{}
	for scanner.Scan() {
		tmp, terr := strconv.Atoi(scanner.Text())
		if terr != nil {
			log.Fatal("Maaan the scanner died with error", terr)
		}
		data[tmp] = struct{}{}
	}

	for k := range data {
		for k2 := range data {
			if _, ok := data[2020-(k+k2)]; ok {
				fmt.Println(k, k2, 2020-(k+k2))
				fmt.Println(k * k2 * (2020 - (k + k2)))
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
