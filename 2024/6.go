package main

import (
	"bufio"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"os"
)

const Max = 130
const (N uint8 = iota E S W)

type Point struct {
	x, y int
}

func main() {
	fmt.Println(solve6P1())
	fmt.Println(solve6P2())
}

func parse() (map[Point]struct{}, Point, uint8) {
	file, _ := os.Open("~/aoc/6-input.txt")
	defer file.Close()

	var start Point
	scanner := bufio.NewScanner(file)
	obs := map[Point]struct{}{}
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j, c := range line {
			switch c {
			case '#':
				obs[Point{i, j}] = struct{}{}
			case '^':
				start = Point{i, j}
			}
		}
	}
	return obs, start, N
}

func solve6P1() int {
	obs, start, dir := parse()
	vis := mapset.NewSet[Point]()
	vis.Add(start)

	for {
		var future Point

		switch dir {
		case N:
			future = Point{start.x - 1, start.y}
		case E:
			future = Point{start.x, start.y + 1}
		case S:
			future = Point{start.x + 1, start.y}
		case W:
			future = Point{start.x, start.y - 1}
		}

		if !(future.x >= 0 && future.x < Max && future.y >= 0 && future.y < Max) {
			break
		}

		if _, ok := obs[future]; ok {
			dir = (dir + 1) % 4
		} else {
			start = future
			vis.Add(start)
		}
	}

	return vis.Cardinality()
}

type State struct {
	point Point
	dir   uint8
}

func solve6P2() int {
	obs, start, dir := parse()

	total := 0
	for i := 0; i < Max; i++ {
		for j := 0; j < Max; j++ {
			if loopOnObstacle(obs, Point{i, j}, start, dir) {
				total++
			}
		}
	}

	return total
}

func loopOnObstacle(obs map[Point]struct{}, target Point, start Point, dir uint8) bool {
	if target == start {
		return false
	}

	if _, ok := obs[target]; ok {
		return false
	}

	obs[target] = struct{}{}
	defer delete(obs, target)

	vis := mapset.NewSet[State]()
	vis.Add(State{start, dir})

	for {
		var future Point
		switch dir {
		case N:
			future = Point{start.x - 1, start.y}
		case E:
			future = Point{start.x, start.y + 1}
		case S:
			future = Point{start.x + 1, start.y}
		case W:
			future = Point{start.x, start.y - 1}
		}

		if !(future.x >= 0 && future.x < Max && future.y >= 0 && future.y < Max) {
			return false
		}

		if _, ok := obs[future]; ok {
			dir = (dir + 1) % 4
		} else {
			start = future
		}

		futureState := State{start, dir}
		if vis.Contains(futureState) {
			return true
		}

		vis.Add(futureState)
	}

}
