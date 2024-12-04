package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	solve()
}

func solve() {
	var grid [][]string
	file, _ := os.Open("~/aoc/4-input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		temp := strings.Split(line, "")
		grid = append(grid, temp)
	}

	total1 := horiz(grid) + vert(grid) + diag(grid)
	total2 := smallXmas(grid)
	fmt.Println(total1)
	fmt.Println(total2)
}

func horiz(g [][]string) int {
	xmas := 0
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i])-3; j++ {
			word := strings.Join(g[i][j:j+4], "")
			if word == "XMAS" || word == "SAMX" {
				xmas += 1
			}
		}
	}
	return xmas
}

func vert(g [][]string) int {
	xmas := 0
	for j := 0; j < len(g[0]); j++ {
		for i := 0; i < len(g)-3; i++ {
			var letters []string
			for k := 0; k < 4; k++ {
				letters = append(letters, g[i+k][j])
			}
			word := strings.Join(letters, "")
			if word == "XMAS" || word == "SAMX" {
				xmas += 1
			}
		}
	}
	return xmas
}

func diag(g [][]string) int {
	xmas := 0

	for j := 3; j < len(g[0]); j++ {
		for i := 0; i < len(g)-3; i++ {
			var letters []string
			for k := 0; k < 4; k++ {
				letters = append(letters, g[i+k][j-k])
			}
			word := strings.Join(letters, "")
			if word == "XMAS" || word == "SAMX" {
				xmas += 1
			}
		}
	}

	for j := 0; j < len(g[0])-3; j++ {
		for i := 0; i < len(g)-3; i++ {
			var letters []string
			for k := 0; k < 4; k++ {
				letters = append(letters, g[i+k][j+k])
			}
			word := strings.Join(letters, "")
			if word == "XMAS" || word == "SAMX" {
				xmas += 1
			}
		}
	}
	return xmas
}

func smallXmas(g [][]string) int {
	xmas := 0
	for j := 0; j <= len(g[0])-3; j++ {
		for i := 0; i <= len(g)-3; i++ {
			var right []string
			var left []string
			for k := 0; k < 3; k++ {
				right = append(right, g[i+k][j+k])
				left = append(left, g[i+2-k][j+k])
			}
			wordRight := strings.Join(right, "")
			wordLeft := strings.Join(left, "")
			if (wordRight == "MAS" || wordRight == "SAM") && (wordLeft == "MAS" || wordLeft == "SAM") {
				xmas += 1
			}
		}
	}
	return xmas
}
