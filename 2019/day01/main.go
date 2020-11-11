// Package file
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		tmp, terr := strconv.Atoi(scanner.Text())
		if terr != nil {
			log.Fatal("Maaan the scanner died with error", terr)
		}
		module := (tmp/3 - 2)
		sum += module
		extra := module/3 - 2
		for {
			if extra > 0 {
				sum += extra
			} else {
				break
			}
			extra = extra/3 - 2
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}
