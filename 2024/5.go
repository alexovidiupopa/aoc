package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	solve5()
}

func solve5() {
	file, _ := os.Open("~/aoc/5-input.txt")
	defer file.Close()

	rules := make(map[int][]int)
	var updates [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		temp := strings.Split(line, "|")
		n1, _ := atoi(temp[0])
		n2, _ := atoi(temp[1])
		rules[n1] = append(rules[n1], n2)
	}

	for scanner.Scan() {
		line := scanner.Text()
		temp := strings.Split(line, ",")
		var tmp []int
		for _, s := range temp {
			num, _ := atoi(s)
			tmp = append(tmp, num)
		}
		updates = append(updates, tmp)
	}

	sum1 := 0

	for _, update := range updates {
		good := true
		for i := 0; i < len(update)-1; i++ {
			for j := i + 1; j < len(update); j++ {
				if slices.Contains(rules[update[j]], update[i]) {
					good = false
					break
				}
			}
			if !good {
				break
			}
		}
		if good {
			sum1 += update[len(update)/2]
		}
	}

	fmt.Println(sum1)

	sum2 := 0

	for _, update := range updates {
		sort.SliceStable(update, func(i, j int) bool {
			return slices.Contains(rules[update[i]], update[j])
		})
		sum2 += update[len(update)/2]
	}
	fmt.Println(sum2 - sum1)
}

func atoi(s string) (int, error) {
	return strconv.Atoi(s)
}
