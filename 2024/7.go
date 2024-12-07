package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	solve7()
}

func solve7() {
	var testValues []int
	var numbers [][]int
	file, _ := os.Open("D:\\code\\aoc\\inputs\\data7.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		temp := strings.Split(line, ":")
		tval, _ := atoi(temp[0])
		testValues = append(testValues, tval)
		var nums []int
		temp2 := strings.Split(strings.TrimPrefix(temp[1], " "), " ")
		for _, s := range temp2 {
			num, _ := atoi(s)
			nums = append(nums, num)
		}
		numbers = append(numbers, nums)
	}

	total1 := 0
	total2 := 0
	for i := 0; i < len(testValues); i++ {

		if good(testValues[i], numbers[i], 1, numbers[i][0], false) {
			total1 += testValues[i]
		} else if good(testValues[i], numbers[i], 1, numbers[i][0], true) {
			total2 += testValues[i]
		}
	}
	fmt.Println(total1)
	fmt.Println(total1 + total2)

}

func atoi(s string) (int, error) {
	return strconv.Atoi(s)
}

func good(target int, values []int, i int, acc int, shouldConcat bool) bool {
	if acc > target {
		return false
	}
	if i == len(values) {
		return acc == target
	}
	if shouldConcat {
		return good(target, values, i+1, acc+values[i], shouldConcat) || good(target, values, i+1, acc*values[i], shouldConcat) || good(target, values, i+1, concat(acc, values[i]), shouldConcat)
	}
	return good(target, values, i+1, acc+values[i], shouldConcat) || good(target, values, i+1, acc*values[i], shouldConcat)

}

func concat(a int, b int) int {
	r, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return r
}
