package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x int
	y int
}

type Region struct {
	area      int
	perimeter int
	points    map[Point]struct{}
}

func main() {
	grid := parse()
	fmt.Println(solveP1(grid))
	fmt.Println(solveP2(grid))
}

func getRegions(grid [][]string) []Region {
	var regions []Region
	viz := make(map[Point]struct{})
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if _, ok := viz[Point{i, j}]; !ok {
				region := Region{points: make(map[Point]struct{})}
				viz[Point{i, j}] = struct{}{}
				buildFullRegion(grid, i, j, viz, &region)
				regions = append(regions, region)
			}
		}
	}

	return regions
}

func buildFullRegion(grid [][]string, i int, j int, viz map[Point]struct{}, r *Region) {
	r.area++
	r.perimeter += 4
	r.points[Point{i, j}] = struct{}{}

	if i > 0 && grid[i][j] == grid[i-1][j] {
		if _, ok := viz[Point{i - 1, j}]; !ok {
			viz[Point{i - 1, j}] = struct{}{}
			buildFullRegion(grid, i-1, j, viz, r)
		}
		r.perimeter--
	}

	if i < len(grid)-1 && grid[i][j] == grid[i+1][j] {
		if _, ok := viz[Point{i + 1, j}]; !ok {
			viz[Point{i + 1, j}] = struct{}{}
			buildFullRegion(grid, i+1, j, viz, r)
		}
		r.perimeter--
	}

	if j > 0 && grid[i][j] == grid[i][j-1] {
		if _, ok := viz[Point{i, j - 1}]; !ok {
			viz[Point{i, j - 1}] = struct{}{}
			buildFullRegion(grid, i, j-1, viz, r)
		}
		r.perimeter--
	}

	if j < len(grid[0])-1 && grid[i][j] == grid[i][j+1] {
		if _, ok := viz[Point{i, j + 1}]; !ok {
			viz[Point{i, j + 1}] = struct{}{}
			buildFullRegion(grid, i, j+1, viz, r)
		}
		r.perimeter--
	}

}

func solveP1(grid [][]string) int {
	regions := getRegions(grid)
	price := 0
	for _, region := range regions {
		price += region.area * region.perimeter
	}
	return price
}

func solveP2(grid [][]string) int {
	regions := getRegions(grid)
	price := 0
	for _, region := range regions {
		price += region.area * sides(region)
	}
	return price
}

func sides(region Region) int {
	sides := 0

	for point, _ := range region.points {
		x := point.x
		y := point.y
		_, w := region.points[Point{x, y - 1}]
		_, n := region.points[Point{x - 1, y}]
		_, e := region.points[Point{x, y + 1}]
		_, s := region.points[Point{x + 1, y}]
		_, nw := region.points[Point{x - 1, y - 1}]
		_, ne := region.points[Point{x - 1, y + 1}]
		_, se := region.points[Point{x + 1, y + 1}]
		_, sw := region.points[Point{x + 1, y - 1}]
		if !w && !n {
			sides++
		}
		if !e && !n {
			sides++
		}
		if !w && !s {
			sides++
		}
		if !e && !s {
			sides++
		}
		if !nw && n && w {
			sides++
		}
		if !ne && n && e {
			sides++
		}
		if !sw && s && w {
			sides++
		}
		if !se && s && e {
			sides++
		}
	}
	return sides
}

func parse() [][]string {
	var grid [][]string
	file, _ := os.Open("D:\\code\\aoc\\inputs\\data12.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		temp := strings.Split(line, "")
		grid = append(grid, temp)
	}
	return grid
}
