package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("/home/alex/aoc/inputs/6-2015.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mat := [1000][1000]int{}

	for scanner.Scan() {

		s := scanner.Text()
		split := strings.Split(s, " ")
		if len(split) == 4 { // toggle
			startX, _ := strconv.Atoi(strings.Split(split[1], ",")[0])
			startY, _ := strconv.Atoi(strings.Split(split[1], ",")[1])
			endX, _ := strconv.Atoi(strings.Split(split[3], ",")[0])
			endY, _ := strconv.Atoi(strings.Split(split[3], ",")[1])
			for i := startX; i <= endX; i++ {
				for j := startY; j <= endY; j++ {
					mat[i][j] = mat[i][j] + 2
				}
			}
		} else {
			// turn
			val := 1
			if split[1] == "off" {
				val = -1
			}
			startX, _ := strconv.Atoi(strings.Split(split[2], ",")[0])
			startY, _ := strconv.Atoi(strings.Split(split[2], ",")[1])
			endX, _ := strconv.Atoi(strings.Split(split[4], ",")[0])
			endY, _ := strconv.Atoi(strings.Split(split[4], ",")[1])
			for i := startX; i <= endX; i++ {
				for j := startY; j <= endY; j++ {
					mat[i][j] = int(math.Max(float64(mat[i][j]+val), 0))
				}
			}
		}

	}

	open := uint64(0)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			open += uint64(mat[i][j])
		}
	}
	fmt.Println(open)
}
