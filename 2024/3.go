package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("D:\\code\\aoc\\inputs\\data3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	strs := ""
	for scanner.Scan() {
		s := scanner.Text()
		strs = strs + s
	}

	solve1(strs)
	solve2(strs)

}

func solve1(strs string) {

	r, _ := regexp.Compile("mul\\([0-9]+\\,[0-9]+\\)")
	match := r.FindAllString(strs, -1)
	sum := 0
	for _, s := range match {
		l := s[4:]
		l = l[:len(l)-1]
		num1, _ := strconv.Atoi(strings.Split(l, ",")[0])
		num2, _ := strconv.Atoi(strings.Split(l, ",")[1])
		sum = sum + num1*num2
	}

	fmt.Println(sum)
}

func solve2(strs string) {

	r, _ := regexp.Compile("mul\\([0-9]+\\,[0-9]+\\)|do\\(\\)|don't\\(\\)")
	match := r.FindAllString(strs, -1)
	sum := 0
	do := true
	for _, s := range match {
		if s == "do()" {
			fmt.Println("here")
			do = true
		} else if s == "don't()" {
			do = false
		} else if do {
			l := s[4:]
			l = l[:len(l)-1]
			num1, _ := strconv.Atoi(strings.Split(l, ",")[0])
			num2, _ := strconv.Atoi(strings.Split(l, ",")[1])
			sum = sum + num1*num2
		}
	}

	fmt.Println(sum)
}
