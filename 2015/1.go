package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("/home/alex/aoc/inputs/1-2015.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		total := 0

		s := scanner.Text()

		for i := 0; i < len(s); i++ {
			if s[i] == '(' {
				total++
			} else {
				total--
			}
			if total == -1 {
				fmt.Println(i + 1)
			}
		}

		fmt.Println(total)

	}
}
