package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("/home/alex/aoc/inputs/2-2015.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	ribbon := 0

	for scanner.Scan() {

		s := scanner.Text()

		abc := strings.Split(s, "x")

		a, _ := strconv.Atoi(abc[0])
		b, _ := strconv.Atoi(abc[1])
		c, _ := strconv.Atoi(abc[2])

		loc := 0
		min := math.Min(float64(a*c), float64(a*b))
		min = math.Min(min, float64(b*c))

		loc = 2*a*b + 2*a*c + 2*b*c + int(min)
		total += loc

		rib := 0
		min = math.Min(float64(2*a+2*c), float64(2*a+2*b))
		min = math.Min(min, float64(2*b+2*c))
		rib = a*b*c + int(min)
		ribbon += rib
	}

	fmt.Println(total)
	fmt.Println(ribbon)
}
