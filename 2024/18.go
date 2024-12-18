package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const N = 70 + 1
const BYTES = 1024

type Point struct {
	x, y int
}

func initializeMatrix(filename string) ([][]bool, error) {
	matrix := make([][]bool, N)
	for i := range matrix {
		matrix[i] = make([]bool, N)
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		matrix[y][x] = true
		b++
		if b == BYTES {
			break
		}
	}

	return matrix, nil
}

func bfs(matrix [][]bool) int {
	var q = []Point{{0, 0}}
	visited := make([][]bool, N)
	for i := range visited {
		visited[i] = make([]bool, N)
	}
	visited[0][0] = true

	var dx = []int{0, 1, 0, -1}
	var dy = []int{1, 0, -1, 0}

	steps := 0
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			point := q[0]
			q = q[1:]

			if point.x == N-1 && point.y == N-1 {
				return steps
			}

			for d := 0; d < 4; d++ {
				nx, ny := point.x+dx[d], point.y+dy[d]
				if nx >= 0 && nx < N && ny >= 0 && ny < N && !matrix[ny][nx] && !visited[ny][nx] {
					visited[ny][nx] = true
					q = append(q, Point{nx, ny})
				}
			}
		}
		steps++
	}

	return -1
}

func solve2(filename string) (int, error) {
	matrix := make([][]bool, N)
	for i := range matrix {
		matrix[i] = make([]bool, N)
	}

	file, err := os.Open(filename)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		matrix[y][x] = true
		if bfs(matrix) == -1 {
			fmt.Println(x, y)
			return -1, nil
		}
	}

	return bfs(matrix), nil
}

func main() {
	matrix, err := initializeMatrix("~/aoc/18.txt")
	
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(matrix)

	fmt.Println(bfs(matrix))
	fmt.Println(solve2("~/aoc/18.txt"))
}
