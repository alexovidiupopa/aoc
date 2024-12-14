package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Robot struct {
	px, py, vx, vy int
}

type Quadrants struct {
	q1, q2, q3, q4 int
}

func parseInput(filename string) ([]Robot, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var robots []Robot
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var robot Robot
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.py, &robot.px, &robot.vy, &robot.vx)
		robots = append(robots, robot)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return robots, nil
}

const n = 103
const m = 101
const iter = 100

var tree = [][]int{
	{0, 0, 0, 0, 1, 0, 0, 0, 0},
	{0, 0, 0, 1, 1, 1, 0, 0, 0},
	{0, 0, 1, 1, 1, 1, 1, 0, 0},
	{0, 1, 1, 1, 1, 1, 1, 1, 0},
	{1, 1, 1, 1, 1, 1, 1, 1, 1},
}

func main() {
	robots, err := parseInput("D:\\code\\aoc\\inputs\\data14.txt")
	if err != nil {
		log.Fatal(err)
	}

	tiles := make([][]int, n)
	for i := range n {
		tiles[i] = make([]int, m)
	}

	for _, robot := range robots {
		tiles[robot.px][robot.py] += 1
	}

	//q := getQuadrants(robots, tiles)
	//fmt.Println(q.q1 * q.q2 * q.q3 * q.q4)
	fmt.Println(getSecondsUntilPattern(robots, tiles))
}

func getQuadrants(robots []Robot, tiles [][]int) Quadrants {
	for range iter {
		for i, robot := range robots {
			tiles[robot.px][robot.py] -= 1
			robot.px = (((robot.px + robot.vx) % n) + n) % n
			robot.py = (((robot.py + robot.vy) % m) + m) % m
			tiles[robot.px][robot.py] += 1
			robots[i] = robot
		}
	}

	var q Quadrants

	for i := 0; i < n/2; i++ {
		for j := 0; j < m/2; j++ {
			q.q1 += tiles[i][j]
			q.q2 += tiles[i][(j+1)+(m/2)]
			q.q3 += tiles[(i+1)+(n/2)][j]
			q.q4 += tiles[(i+1)+(n/2)][(j+1)+(m/2)]
		}
	}
	return q
}

func getSecondsUntilPattern(robots []Robot, tiles [][]int) int {
	seconds := 1
	for {
		for i, robot := range robots {
			tiles[robot.px][robot.py] -= 1
			robot.px = (((robot.px + robot.vx) % n) + n) % n
			robot.py = (((robot.py + robot.vy) % m) + m) % m
			tiles[robot.px][robot.py] += 1
			robots[i] = robot
		}
		for i := 0; i < n-len(tree); i++ {
			for j := 0; j < m-len(tree[0]); j++ {
				if isTreePattern(tiles, i, j) {
					return seconds
				}
			}
		}
		seconds += 1
	}
}

func isTreePattern(tiles [][]int, x int, y int) bool {
	for i := 0; i < len(tree); i++ {
		for j := 0; j < len(tree[0]); j++ {
			if tree[i][j]-tiles[i+x][j+y] == 1 {
				return false
			}
		}
	}
	return true
}
