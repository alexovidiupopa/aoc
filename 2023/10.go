package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"modernc.org/mathutil"
)

func main() {
	solve()
}

type Point struct {
	x, y int
}

var mat = make([][]rune, 0)
var vis = make(map[Point]bool)
var pipe = make(map[Point]bool)
var start = Point{-1, -1}

func solve() {
	fileContent, err := ioutil.ReadFile("/home/alex/aoc/inputs/10-2023.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")

	for y, line := range lines {
		chars := []rune(line)
		mat = append(mat, chars)
		for x, c := range chars {
			if c == 'S' {
				start = Point{x, y}
			}
		}
	}

	fmt.Println(solve1(start))
	fmt.Println(solve2(start))
}

func solve1(start Point) int {
	p := start

	right := dfs(Point{p.x + 1, p.y}, 0)
	vis = make(map[Point]bool) // clear the vis after every search

	left := dfs(Point{p.x - 1, p.y}, 0)
	vis = make(map[Point]bool)

	down := dfs(Point{p.x, p.y + 1}, 0)
	vis = make(map[Point]bool)

	up := dfs(Point{p.x, p.y - 1}, 0)

	return mathutil.MaxVal(right, left, down, up) / 2 // div by 2 since it's a cycle
}

func solve2(start Point) int {

	p := start

	right := dfs(Point{p.x + 1, p.y}, 0)

	maxlen := right
	maxpipe := pipe

	vis = make(map[Point]bool)
	pipe = make(map[Point]bool)

	left := dfs(Point{p.x - 1, p.y}, 0)
	if left > maxlen {
		maxlen = left
		maxpipe = pipe
	}
	vis = make(map[Point]bool)
	pipe = make(map[Point]bool)

	down := dfs(Point{p.x, p.y + 1}, 0)
	if down > maxlen {
		maxlen = down
		maxpipe = pipe
	}
	vis = make(map[Point]bool)
	pipe = make(map[Point]bool)

	up := dfs(Point{p.x, p.y - 1}, 0)
	if up > maxlen {
		maxlen = up
		maxpipe = pipe
	}

	mat[start.y][start.y] = '|'

	// maxpipe now contains the coords of the longest cycle (closed pipe)

	total := 0
	for y, line := range mat {
		in := false // assume we're not inside the walls of the pipe
		prev := 'x'
		for x, curr := range line {
			if !maxpipe[Point{x, y}] { // not on the pipe
				if in {
					total++
				}
				continue
			}
			if curr == 'x' { // skip
				continue
			}

			// 'LJ' and 'F7' equate to || - meaning I am either going outside or inside the walls
			if prev == 'L' && curr == 'J' {
				in = !in
				prev = 'x'
			}
			if prev == 'F' && curr == '7' {
				in = !in
				prev = 'x'
			}
			if curr == '|' || curr == 'F' || curr == 'L' { // either stepped outside or scan for LJ/F7
				in = !in
				prev = curr
			}
		}
	}
	return total
}

func outOfBounds(p Point) bool {
	return p.x < 0 || p.y < 0 || p.x >= len(mat[0]) || p.y >= len(mat)
}

func dfs(p Point, step int) int {
	if vis[p] || outOfBounds(p) || mat[p.y][p.x] == '.' {
		return step
	}
	step++
	vis[p] = true
	pipe[p] = true

	switch mat[p.y][p.x] {
	case '|':
		down := dfs(Point{p.x, p.y + 1}, step)
		up := dfs(Point{p.x, p.y - 1}, step)
		return mathutil.Max(down, up)
	case '-':
		right := dfs(Point{p.x + 1, p.y}, step)
		left := dfs(Point{p.x - 1, p.y}, step)
		return mathutil.Max(right, left)
	case 'F':
		right := dfs(Point{p.x + 1, p.y}, step)
		down := dfs(Point{p.x, p.y + 1}, step)
		return mathutil.Max(right, down)
	case '7':
		left := dfs(Point{p.x - 1, p.y}, step)
		down := dfs(Point{p.x, p.y + 1}, step)
		return mathutil.Max(left, down)
	case 'J':
		left := dfs(Point{p.x - 1, p.y}, step)
		up := dfs(Point{p.x, p.y - 1}, step)
		return mathutil.Max(left, up)
	case 'L':
		right := dfs(Point{p.x + 1, p.y}, step)
		up := dfs(Point{p.x, p.y - 1}, step)
		return mathutil.Max(right, up)
	case 'S':
		// nothing to do
		break
	}

	return step
}
