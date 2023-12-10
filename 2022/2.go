// 2
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file, _ := os.Open("2-2022.txt")
	scanner := bufio.NewScanner(file)

	numbers := []int{}
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		game := strings.Split(line, " ")
		elf := game[0]
		me := game[1]
		score := 0
		if elf == "A" {
			if me == "X" {
				score = 3
			}
			if me == "Y" {
				score = 4
			}
			if me == "Z" {
				score = 8
			}
		}
		if elf == "B" {
			if me == "X" {
				score = 1
			}
			if me == "Y" {
				score = 5
			}
			if me == "Z" {
				score = 9
			}
		}
		if elf == "C" {
			if me == "X" {
				score = 2
			}
			if me == "Y" {
				score = 6
			}
			if me == "Z" {
				score = 7
			}
		}
		total += score

	}
	fmt.Println(total)
}
