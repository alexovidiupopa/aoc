package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Program struct {
	a, b, c      int
	instructions []int
}

func parseInput(filename string) (*Program, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var program Program
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()

	line = scanner.Text()
	fmt.Sscanf(line, "Register A: %d", &program.a)
	scanner.Scan()
	line = scanner.Text()
	fmt.Sscanf(line, "Register B: %d", &program.b)
	scanner.Scan()
	line = scanner.Text()
	fmt.Sscanf(line, "Register C: %d", &program.c)
	scanner.Scan()
	line = scanner.Text()
	var l string
	fmt.Sscanf(line, "Program: %s", &l)
	instructions := strings.Split(l, ",")
	for _, instruction := range instructions {
		instr, _ := strconv.Atoi(instruction)
		program.instructions = append(program.instructions, instr)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &program, nil
}

func solve(program Program) []int {
	var ans []int
	for i := 0; i < len(program.instructions); i += 2 {
		opcode, operand := program.instructions[i], program.instructions[i+1]
		val := operand

		switch operand {
		case 4:
			val = program.a
		case 5:
			val = program.b
		case 6:
			val = program.c
		}

		switch opcode {
		case 0:
			program.a >>= val
		case 1:
			program.b ^= operand
		case 2:
			program.b = val % 8
		case 3:
			if program.a != 0 {
				i = operand - 2
			}
		case 4:
			program.b ^= program.c
		case 5:
			ans = append(ans, val%8)
		case 6:
			program.b = program.a >> val
		case 7:
			program.c = program.a >> val
		}
	}
	return ans
}
func main() {
	program, err := parseInput("/aoc/17.txt")

	if err != nil {
		fmt.Println(err)
	}

	ans := solve(*program)
	a := 0
	for pos := len(program.instructions) - 1; pos >= 0; pos-- {
		a <<= 3 // shift left by 3 bits for each position
		program.a = a

		for !slices.Equal(solve(*program), program.instructions[pos:]) {
			a++
			program.a = a
		}
	}
	fmt.Println(ans)
	fmt.Println(a)
}
