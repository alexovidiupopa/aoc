package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("D:\\code\\aoc\\inputs\\data2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	count1 := 0
	count2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		list := []int{}
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err == nil {
				list = append(list, num)
			}
		}

		if isLegal1(list) {
			count1 += 1
		}
		if isLegal2(list) {
			count2 += 1
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(count1)
	fmt.Println(count2)

}

func isLegal1(list []int) bool {
	asc := false
	if list[0] < list[1] {
		asc = true
	} else if list[0] > list[1] {
		asc = false
	}
	for i := 0; i < len(list)-1; i++ {
		diff := int64(math.Abs(float64(list[i] - list[i+1])))
		if (diff < 1 || diff > 3) || (list[i] < list[i+1] && !asc) || (list[i] > list[i+1] && asc) {
			return false
		}

	}
	return true
}

func remove(slice []int, s int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return append(newSlice[:s], newSlice[s+1:]...)
}

func isLegal2(list []int) bool {
	final := false

	for i := 0; i < len(list)-1; i++ {
		slice1 := remove(list, i)
		slice2 := remove(list, i+1)
		final = final || isLegal1(slice1) || isLegal1(slice2)
	}

	return final
}
