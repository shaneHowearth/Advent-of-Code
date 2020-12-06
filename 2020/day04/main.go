// Package main -
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
	internal := false
	if len(results) == 8 {
		internal = true
	}
	if len(results) == 7 {
		if _, ok := results["cid"]; !ok {
			internal = true
		}
	}
	if internal {
		byr, err := strconv.Atoi(results["byr"])
		if err != nil {
			log.Printf("BYR error %v %#v", err, results["byr"])
			return false
		}
		if byr < 1920 || byr > 2002 {
			return false
		}

		iyr, err := strconv.Atoi(results["iyr"])
		if err != nil {
			log.Printf("IYR error %v %#v", err, results["iyr"])
			return false
		}
		if iyr < 2010 || iyr > 2020 {
			return false
		}

		eyr, err := strconv.Atoi(results["eyr"])
		if err != nil {
			log.Printf("EYR error %v %#v", err, results["eyr"])
			return false
		}
		if eyr < 2020 || eyr > 2030 {
			return false
		}

		hgt, err := strconv.Atoi(results["hgt"][:len(results["hgt"])-2])
		if err != nil {
			log.Printf("hgt error %v %#v", err, results["hgt"])
			return false
		}
		bottom := 59
		top := 76
		if strings.Contains(results["hgt"], "cm") {
			bottom = 150
			top = 193
		}
		if hgt < bottom || hgt > top {
			return false
		}

		internal, err = regexp.MatchString("#[0-9a-f]{6}", results["hcl"])
		if err != nil {
			log.Printf("hcl error %v %#v", err, results["hcl"])
			return false
		}
		if !internal {
			return false
		}

		ecl := results["ecl"]
		if !strings.Contains(ecl, "amb") && !strings.Contains(ecl, "blu") && !strings.Contains(ecl, "brn") && !strings.Contains(ecl, "gry") && !strings.Contains(ecl, "grn") && !strings.Contains(ecl, "hzl") && !strings.Contains(ecl, "oth") {
			return false
		}

		internal, err = regexp.MatchString("^[0-9]{9}$", results["pid"])
		if err != nil {
			log.Printf("pid error %v %#v", err, results["pid"])
			return false
		}
		if !internal {
			return false
		}
		return true
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
