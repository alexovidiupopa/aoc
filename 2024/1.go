package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("D:\\code\\aoc\\2024\\1\\data1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	list1 := []int{}
	list2 := []int{}
	occs := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error parsing numbers:", line)
			continue
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)
		occs[num2] = occs[num2] + 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	//solve1(list1, list2)

	solve2(list1, occs)

}

func solve1(list1 []int, list2 []int) {
	sort.Ints(list1)
	sort.Ints(list2)
	diff := float64(0)

	for i := range list1 {
		diff = diff + math.Abs(float64(list1[i]-list2[i]))
	}

	fmt.Println(diff)
}

func solve2(list1 []int, occs map[int]int) {
	sort.Ints(list1)
	diff := int64(0)

	for i := range list1 {
		diff = diff + int64(list1[i]*occs[list1[i]])
	}

	fmt.Println(diff)
}
