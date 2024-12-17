package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

type Node struct {
	point Point
	dir   int
}

type State struct {
	node Node
	cost int
	idx  int
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].idx, pq[j].idx = i, j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*State)
	item.idx = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.idx = -1
	*pq = old[0 : n-1]
	return item
}

var dirs = []Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func parse() [][]string {
	var grid [][]string
	file, _ := os.Open("/aoc/16.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		temp := strings.Split(line, "")
		grid = append(grid, temp)
	}
	return grid
}

func dijkstra(grid [][]string, start, end Point) (int, [][]bool) {
	vis := map[Node]bool{}
	prev := map[Node]*State{}
	pq := PriorityQueue{&State{Node{start, 1}, 0, 0}}
	heap.Init(&pq)

	for len(pq) > 0 {
		state := heap.Pop(&pq).(*State)
		if state.node.point == end {
			// Backtrack to find the path
			paths := make([][]bool, len(grid))
			for i := range paths {
				paths[i] = make([]bool, len(grid[0]))
			}
			for s := state; s != nil; s = prev[s.node] {
				paths[s.node.point.x][s.node.point.y] = true
			}
			return state.cost, paths
		}
		if vis[state.node] {
			continue
		}
		vis[state.node] = true
		// Move forward in the current direction
		dir := dirs[state.node.dir]
		newPoint := Point{state.node.point.x + dir.x, state.node.point.y + dir.y}
		if newPoint.x >= 0 && newPoint.x < len(grid) && newPoint.y >= 0 && newPoint.y < len(grid[0]) && grid[newPoint.x][newPoint.y] != "#" {
			newState := &State{Node{newPoint, state.node.dir}, state.cost + 1, 0}
			if !vis[newState.node] || newState.cost < prev[newState.node].cost {
				heap.Push(&pq, newState)
				prev[newState.node] = state
			}
		}

		// Rotate clockwise
		newDir := (state.node.dir + 1) % 4
		newState := &State{Node{state.node.point, newDir}, state.cost + 1000, 0}
		if !vis[newState.node] || newState.cost < prev[newState.node].cost {
			heap.Push(&pq, newState)
			prev[newState.node] = state
		}

		// Rotate counterclockwise
		newDir = (state.node.dir + 3) % 4
		newState = &State{Node{state.node.point, newDir}, state.cost + 1000, 0}
		if !vis[newState.node] || newState.cost < prev[newState.node].cost {
			heap.Push(&pq, newState)
			prev[newState.node] = state
		}
	}
	return -1, nil
}

func main() {
	grid := parse()
	var start, end Point
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "S" {
				start = Point{i, j}
			}
			if grid[i][j] == "E" {
				end = Point{i, j}
			}
		}
	}
	ans, paths := dijkstra(grid, start, end)
	onPath := 0
	for i := 0; i < len(paths); i++ {
		for j := 0; j < len(paths[0]); j++ {
			if paths[i][j] {
				onPath++
			}
		}
	}
	fmt.Println(ans, onPath)
}
