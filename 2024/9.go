package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(solve9P1())
	fmt.Println(solve9P2())
}

func parse() filesystem {
	files := ""
	file, _ := os.Open("D:\\code\\aoc\\inputs\\data9.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		files += scanner.Text()
	}

	fs := filesystem{blocks: []block{}}
	alloc := true
	id := 0
	for i := 0; i < len(files); i++ {
		n := int(files[i] - '0')
		for j := 0; j < n; j++ {
			b := block{i: len(fs.blocks)}
			if alloc {
				b.id = id
			} else {
				b.id = -1
			}

			fs.blocks = append(fs.blocks, b)
		}

		alloc = !alloc
		if i%2 == 0 {
			id++
		}
	}

	return fs
}

type block struct {
	i  int
	id int
}

type filesystem struct {
	blocks []block
}

func defragment(fs filesystem) filesystem {
	left := 0
	for right := len(fs.blocks) - 1; right >= 0 && left < right-1; right-- {
		if fs.blocks[right].id != -1 {
			for left < len(fs.blocks) && fs.blocks[left].id != -1 {
				left++
			}
			if left < len(fs.blocks) {
				fs.blocks[left].id = fs.blocks[right].id
				fs.blocks[right].id = -1
			}
		}
	}
	return fs
}

func defragment2(fs filesystem) filesystem {
	right := len(fs.blocks) - 1
	for right >= 0 {
		nextBlockIndex, nextBlockSize := nextBlock(fs, right)
		if nextBlockSize == 0 {
			break
		}
		if nextBlockIndex != -1 {
			nextFreeChunkStart := nextFreeChunk(fs, nextBlockSize)
			if nextFreeChunkStart != -1 && nextFreeChunkStart < nextBlockIndex {
				n := 0
				for i := nextFreeChunkStart; i < nextFreeChunkStart+nextBlockSize; i++ {
					fs.blocks[i].id = fs.blocks[nextBlockIndex+n].id
					fs.blocks[nextBlockIndex+n].id = -1
					n++
				}
			}
		}
		right = nextBlockIndex - 1
	}
	return fs
}

func nextFreeChunk(fs filesystem, size int) int {
	start, end := -1, -1
	for i := 0; i < len(fs.blocks); i++ {
		if fs.blocks[i].id == -1 && start == -1 {
			start = i
			continue
		}
		if fs.blocks[i].id != -1 && start != -1 {
			end = i
			if end-start >= size {
				break
			} else {
				start = -1
			}
		}
	}
	if end-start >= size {
		return start
	}
	return -1
}

func nextBlock(fs filesystem, right int) (int, int) {
	idx, size := -1, -1
	start, end := -1, -1
	id := -1
	i := right
	for start == -1 || end == -1 {
		cur := i
		i--
		if i < 0 {
			break
		}
		boundary := fs.blocks[cur].id == -1 || (id != -1 && fs.blocks[cur].id != id)
		if boundary {
			if idx != -1 {
				start = cur
				break
			}
			continue
		}
		if idx == -1 {
			id = fs.blocks[cur].id
			idx = cur
			end = cur
		}
	}

	idx = start + 1
	size = end - start
	return idx, size
}

func solve9P1() int {
	fs := parse()
	return checksum(defragment(fs))
}

func solve9P2() int {
	fs := parse()
	return checksum(defragment2(fs))
}

func checksum(fs filesystem) int {
	s := 0
	for i, b := range fs.blocks {
		if b.id != -1 {
			s += b.id * i
		}
	}
	return s
}
