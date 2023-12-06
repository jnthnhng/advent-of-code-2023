package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	if s.Err() != nil {
		log.Fatal(s.Err())
	}

	var data []string

	// Read data from input file and combine into an array of arrays
	for s.Scan() {
		line := s.Text()
		data = append(data, line)
	}
	//fmt.Println(data)

	for r, line := range data {
		for c, _ := range line {
			//fmt.Println(r, c)
			neighbors := findNeighbors(r, c, len(data), len(data[0]))
			fmt.Println(neighbors)
		}
	}

}

type pair struct {
	r int
	c int
}

func findNeighbors(r, c, rowLen, colLen int) []pair {
	deltaRows := []int{-1, 0, 1, 0}
	deltaCols := []int{0, 1, 0, -1}

	var neighbors []pair
	for i := range deltaRows {
		nr := deltaRows[i] + r
		nc := deltaCols[i] + c

		//fmt.Println(nr, nc)
		if (0 <= nr && nr < rowLen) && (0 <= nc && nc < colLen) {
			neighbor := pair{r: nr, c: nc}
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}
