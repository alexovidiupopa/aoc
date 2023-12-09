package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("/home/alex/aoc/inputs/4-2015.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		s := scanner.Bytes()

		n := uint64(1)

		hash := getMD5Hash(strconv.AppendUint(s, n, 10))

		for {
			if strings.HasPrefix(hash, "000000") {
				break
			}
			n++
			hash = getMD5Hash(strconv.AppendUint(s, n, 10))
		}
		fmt.Println(n)

	}
}

func getMD5Hash(text []byte) string {
	hash := md5.Sum(text)
	return hex.EncodeToString(hash[:])
}
