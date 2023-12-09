package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("/home/alex/aoc/inputs/5-2015.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	total2 := 0
	for scanner.Scan() {

		s := scanner.Text()

		if nice(s) {
			total++
		}

		if nice2(s) {
			total2++
		}
	}

	fmt.Println(total)
	fmt.Println(total2)
}

func nice(s string) bool {
	return niceVowels(s) && niceRepeated(s) && niceGroups(s)
}

func niceVowels(s string) bool {
	return strings.Count(s, "a")+strings.Count(s, "e")+strings.Count(s, "i")+strings.Count(s, "o")+strings.Count(s, "u") > 2
}

func niceRepeated(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func niceGroups(s string) bool {
	return strings.Count(s, "ab")+strings.Count(s, "cd")+strings.Count(s, "pq")+strings.Count(s, "xy") == 0
}

func nice2(s string) bool {
	return niceTriplet(s) && niceNotOverlappingPairs(s)
}

func niceTriplet(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func niceNotOverlappingPairs(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if strings.Count(s, s[i:i+2]) >= 2 {
			return true
		}
	}
	return false
}
