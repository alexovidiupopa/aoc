package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseInput() ([]int, error) {
	file, err := os.Open("D:\\aoc\\2024\\22.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		number, _ := strconv.Atoi(line)

		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return numbers, nil
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	numbers, _ := parseInput()

	sum := 0
	for k := range numbers {
		for i := 0; i < 2000; i++ {
			numbers[k] = mixAndPrune(numbers[k], numbers[k]*64)
			numbers[k] = mixAndPrune(numbers[k], int(numbers[k]/32))
			numbers[k] = mixAndPrune(numbers[k], numbers[k]*2048)
		}
		sum += numbers[k]
	}
	return sum
}

func mixAndPrune(a, b int) int {
	return (a ^ b) % 16777216
}

type pair struct {
	diff    int
	bananas int
}

type window struct {
	x, y, z, w int
}

func part2() int {
	numbers, _ := parseInput()
	cache := make(map[window]int)
	for k := range numbers {
		vis := make(map[window]bool)
		del := make([]pair, 0, 1999)
		currWindow := make([]int, 0, 8)
		prev := numbers[k] % 10
		for i := 0; i < 2000; i++ {

			numbers[k] = mixAndPrune(numbers[k], numbers[k]*64)
			numbers[k] = mixAndPrune(numbers[k], int(numbers[k]/32))
			numbers[k] = mixAndPrune(numbers[k], numbers[k]*2048)

			diff, bananas := numbers[k]%10-prev, numbers[k]%10
			del = append(del, pair{diff: diff, bananas: bananas})
			currWindow = append(currWindow, diff)

			if len(currWindow) == 4 {
				key := window{x: currWindow[0], y: currWindow[1], z: currWindow[2], w: currWindow[3]}
				if !vis[key] {
					cache[key] += numbers[k] % 10
					vis[key] = true
				}
				currWindow = currWindow[1:]
			}

			prev = numbers[k] % 10
		}
	}
	mx := 0
	for _, v := range cache {
		mx = max(mx, v)
	}
	return mx

}
