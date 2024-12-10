package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(solve10P1())
	fmt.Println(solve10P2())
}

func parse() map[Point]int {
	graph := make(map[Point]int)
	var grid [][]string
	file, _ := os.Open("D:\\code\\aoc\\inputs\\data10.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		temp := strings.Split(line, "")
		grid = append(grid, temp)
	}

	for i, line := range grid {
		for j, c := range line {
			v, _ := strconv.Atoi(c)
			graph[Point{i, j}] = v
		}
	}
	return graph
}

type Point struct {
	x int
	y int
}

func zeroes(graph map[Point]int) []Point {
	var z []Point
	for k, v := range graph {
		if v == 0 {
			z = append(z, k)
		}
	}
	return z
}

func possibleMoves(graph map[Point]int, curr Point, height int) []Point {
	var moves []Point
	adjs := []Point{{curr.x + 1, curr.y}, {curr.x, curr.y + 1}, {curr.x - 1, curr.y}, {curr.x, curr.y - 1}}
	for _, adj := range adjs {
		if c, ok := graph[adj]; ok && c == height+1 {
			moves = append(moves, adj)
		}
	}
	return moves
}

func possiblePaths(graph map[Point]int, distinct bool) int {
	s := 0
	zeros := zeroes(graph)
	for _, point := range zeros {
		s += bfs(graph, point, distinct)
	}
	return s

}

func bfs(graph map[Point]int, point Point, distinct bool) int {
	nines := 0
	viz := make(map[Point]bool)
	queue := []Point{point}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if _, ok := viz[curr]; ok && distinct {
			continue
		}
		viz[curr] = true

		if graph[curr] == 9 {
			nines++
			continue
		}

		moves := possibleMoves(graph, curr, graph[curr])
		queue = append(queue, moves...)

	}
	return nines
}

func solve10P1() int {
	graph := parse()
	return possiblePaths(graph, true)
}

func solve10P2() int {
	graph := parse()
	return possiblePaths(graph, false)
}
