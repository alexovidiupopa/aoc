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
	file, err := os.Open("~/aoc/11.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// use a map of frequencies, as the order is irrelevant
	stones := make(map[int]int)
	for scanner.Scan() {
		s := scanner.Text()
		for _, st := range strings.Split(s, " ") {
			stt, _ := strconv.Atoi(st)
			stones[stt]++
		}
	}

	fmt.Println(solve11(stones, 25))
	fmt.Println(solve11(stones, 75))

}

func solve11(stones map[int]int, blinks int) int {
	for i := 0; i < blinks; i++ {
		state := make(map[int]int)
		for k, v := range stones {
			if k == 0 {
				state[1] += v
			} else if s := strconv.Itoa(k); len(s)%2 != 0 {
				state[k*2024] += v
			} else {
				p1, _ := strconv.Atoi(s[:len(s)/2])
				p2, _ := strconv.Atoi(s[len(s)/2:])
				state[p1] += v
				state[p2] += v
			}
		}
		stones = state
	}
	total := 0
	for _, v := range stones {
		total += v
	}

	return total
}
