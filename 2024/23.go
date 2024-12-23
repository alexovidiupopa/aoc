package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	connections := make(map[string]map[string]bool)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			continue
		}
		a, b := parts[0], parts[1]
		if connections[a] == nil {
			connections[a] = make(map[string]bool)
		}
		if connections[b] == nil {
			connections[b] = make(map[string]bool)
		}
		connections[a][b] = true
		connections[b][a] = true
	}

	sets := findTrios(connections)
	fmt.Println(len(filterTrios(sets)))

	largestClique := findLargestClique(connections)
	fmt.Println(generatePassword(largestClique))

}

func findTrios(connections map[string]map[string]bool) [][]string {
	var sets [][]string
	for a := range connections {
		for b := range connections[a] {
			if a >= b {
				continue
			}
			for c := range connections[a] {
				if b >= c || !connections[b][c] {
					continue
				}
				sets = append(sets, []string{a, b, c})
			}
		}
	}
	return sets
}

func findLargestClique(connections map[string]map[string]bool) []string {
	var largestClique []string
	var currentClique []string
	var visited = make(map[string]bool)

	for node := range connections {
		if !visited[node] {
			currentClique = []string{node}
			visited[node] = true
			for neighbor := range connections[node] {
				if isClique(currentClique, neighbor, connections) {
					currentClique = append(currentClique, neighbor)
					visited[neighbor] = true
				}
			}
			if len(currentClique) > len(largestClique) {
				largestClique = currentClique
			}
		}
	}
	return largestClique
}

func isClique(clique []string, node string, connections map[string]map[string]bool) bool {
	for _, member := range clique {
		if !connections[member][node] {
			return false
		}
	}
	return true
}

func generatePassword(clique []string) string {
	sort.Strings(clique)
	return strings.Join(clique, ",")
}

func filterTrios(sets [][]string) [][]string {
	var filteredSets [][]string
	for _, set := range sets {
		if startsWithT(set) {
			filteredSets = append(filteredSets, set)
		}
	}
	return filteredSets
}

func startsWithT(set []string) bool {
	for _, s := range set {
		if strings.HasPrefix(s, "t") {
			return true
		}
	}
	return false
}
