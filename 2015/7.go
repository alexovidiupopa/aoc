// 7
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("/home/alex/aoc/inputs/7-2015.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	m := make(map[string]int)
	instr := make([]string)
	m["1"] = 1
	for scanner.Scan() {

		s := scanner.Text()
		splt := strings.Split(s, " ")
		n := len(splt)
		if n == 3 { // 123 -> x
			x, _ := strconv.Atoi(splt[0])
			m[splt[2]] = x
		}
		if n == 4 { // not x -> a
			m[splt[3]] = ^m[splt[1]] + 65536
		}
		if n == 5 { // and,or,lshift,rshift
			if splt[1] == "AND" {
				a := (m[splt[0]])
				b := (m[splt[2]])
				m[splt[4]] = (a & b)
			}
			if splt[1] == "OR" {
				a := (m[splt[0]])
				b := (m[splt[2]])
				m[splt[4]] = (a | b)
			}
			if splt[1] == "LSHIFT" {
				a := (m[splt[0]])
				b, _ := strconv.Atoi(splt[2])
				m[splt[4]] = (a << b)
			}
			if splt[1] == "RSHIFT" {
				a := (m[splt[0]])
				b, _ := strconv.Atoi(splt[2])
				m[splt[4]] = (a >> b)
			}
		}

	}
	fmt.Println(m)
}
