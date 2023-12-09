package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("/home/alex/aoc/inputs/3-2015.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		path := scanner.Text()

		fmt.Println(presents(path))
		fmt.Println(presents2(path))
	}

}

type Point struct {
	x int
	y int
}

func presents(path string) int {

	x := 0
	y := 0

	points := make([]Point, 0, 1000)

	for _, d := range path {
		if d == '^' {
			x--
		}
		if d == '>' {
			y++
		}
		if d == '<' {
			y--
		}
		if d == 'v' {
			x++
		}
		if !contains(points, Point{x, y}) {
			points = append(points, Point{x, y})
		}
	}

	return len(points)
}

func presents2(path string) int {

	xSanta := 0
	ySanta := 0

	xRobo := 0
	yRobo := 0

	points := make([]Point, 0, 1000)
	points = append(points, Point{0, 0})

	for i := 0; i < len(path); i += 2 {
		if path[i] == '^' {
			xSanta--
		}
		if path[i] == '>' {
			ySanta++
		}
		if path[i] == '<' {
			ySanta--
		}
		if path[i] == 'v' {
			xSanta++
		}

		if path[i+1] == '^' {
			xRobo--
		}
		if path[i+1] == '>' {
			yRobo++
		}
		if path[i+1] == '<' {
			yRobo--
		}
		if path[i+1] == 'v' {
			xRobo++
		}

		if !contains(points, Point{xSanta, ySanta}) {
			points = append(points, Point{xSanta, ySanta})
		}
		if !contains(points, Point{xRobo, yRobo}) {
			points = append(points, Point{xRobo, yRobo})
		}
	}

	return len(points)
}

func contains(points []Point, point Point) bool {
	for i := 0; i < len(points); i++ {
		if points[i].x == point.x && points[i].y == point.y {
			return true
		}
	}
	return false
}
