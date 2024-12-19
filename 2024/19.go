package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parse(filename string) ([]string, []string) {
	var towels []string
	var designs []string
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			towels = strings.Split(line, ", ")
		} else if line != "" {
			designs = append(designs, line)
		}
		i++
	}
	return towels, designs
}

func getPossible(design string, towels []string, cache map[string]int) int {
	c := 0

	for _, towel := range towels {
		if n, ok := cache[design]; ok {
			return n
		}
		if strings.HasPrefix(design, towel) {
			c += getPossible(design[len(towel):], towels, cache)
		} else if len(design) == 0 {
			return 1
		}
	}

	cache[design] = c
	return c
}

func main() {
	towels, designs := parse("D:\\code\\aoc\\2024\\19.txt")
	//fmt.Println(towels)
	//fmt.Println(designs)
	total := 0
	for _, design := range designs {
		cache := map[string]int{}
		possible := getPossible(design, towels, cache)
		// part 1
		//if possible > 0 {
		//	total++
		//}
		if possible > 0 {
			total += possible
		}
	}
	fmt.Println(total)
}
