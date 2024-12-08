package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	solve8P1()
	solve8P2()
}

func solve8P1() {
	grid, antennas := read()

	for _, points := range antennas {
		for i, p1 := range points {
			for j, p2 := range points {
				if i != j {
					x := p1.x + (p1.x - p2.x)
					y := p1.y + (p1.y - p2.y)
					if x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0]) {
						grid[x][y] = "#"
					}
				}
			}
		}
	}

	total1 := 0
	for _, line := range grid {
		for _, el := range line {
			if el == "#" {
				total1++
			}
		}
	}

	fmt.Println(total1)
}

func solve8P2() {
	grid, antennas := read()

	for _, points := range antennas {
		for i, p1 := range points {
			for j, p2 := range points {
				grid[p1.x][p1.y] = "#"
				if i != j {
					x := p1.x + (p1.x - p2.x)
					y := p1.y + (p1.y - p2.y)
					good := x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0])
					for good {
						grid[x][y] = "#"
						x += (p1.x - p2.x)
						y += (p1.y - p2.y)
						good = x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0])
					}
				}
			}
		}
	}

	total2 := 0
	for _, line := range grid {
		for _, el := range line {
			if el == "#" {
				total2++
			}
		}
	}

	fmt.Println(total2)
}

func read() ([][]string, map[string][]Point) {
	var grid [][]string
	antennas := make(map[string][]Point)
	file, _ := os.Open("D:\\code\\aoc\\inputs\\data8.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		temp := strings.Split(line, "")
		grid = append(grid, temp)
	}

	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] != "." {
				antennas[grid[i][j]] = append(antennas[grid[i][j]], Point{i, j})
			}
		}
	}
	return grid, antennas
}
