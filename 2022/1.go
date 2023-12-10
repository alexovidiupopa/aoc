// 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("1-2022.txt")
	scanner := bufio.NewScanner(file)

	numbers := []int{}
	local := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.EqualFold(line, "") {
			numbers = append(numbers, local)
			local = 0
		} else {
			num, _ := strconv.Atoi(line)
			local += num
		}
	}
	numbers = append(numbers, local)
	sort.Ints(numbers)
	fmt.Println(numbers)
}
