package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Machine struct {
	ax, ay, bx, by, px, py int
}

func parseInput(filename string, extra int) ([]Machine, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var machines []Machine
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		var machine Machine
		fmt.Sscanf(line, "Button A: X+%d, Y+%d", &machine.ax, &machine.ay)
		scanner.Scan()
		line = scanner.Text()
		fmt.Sscanf(line, "Button B: X+%d, Y+%d", &machine.bx, &machine.by)
		scanner.Scan()
		line = scanner.Text()
		fmt.Sscanf(line, "Prize: X=%d, Y=%d", &machine.px, &machine.py)
		machine.px += extra
		machine.py += extra
		machines = append(machines, machine)
		scanner.Scan()
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return machines, nil
}

func main() {
	machines1, err := parseInput("/Users/popaalex/dev/projects/work/aoc/13.txt", 0)

	if err != nil {
		log.Fatal(err)
	}

	totalTokens := solveEfficient(machines1)
	fmt.Println(totalTokens)

	machines2, err := parseInput("/Users/popaalex/dev/projects/work/aoc/13.txt", 10000000000000)

	if err != nil {
		log.Fatal(err)
	}

	totalTokens2 := solveEfficient(machines2)
	fmt.Println(totalTokens2)
}

func findMinTokens(machine Machine) (int, bool) {
	D := machine.ax*machine.by - machine.ay*machine.bx  // determinant of the main system
	Dx := machine.px*machine.by - machine.py*machine.bx // determinant when we replace the first column with the results
	Dy := machine.ax*machine.py - machine.ay*machine.px // determinant when we replace the second column with the results
	if D != 0 && Dx%D == 0 && Dy%D == 0 {
		return (Dx/D)*3 + (Dy/D)*1, true
	}

	return 0, false
}

func solveEfficient(machines []Machine) int {
	totalTokens := 0

	for _, machine := range machines {
		tokens, found := findMinTokens(machine)
		if found {
			totalTokens += tokens
		}
	}
	return totalTokens
}
