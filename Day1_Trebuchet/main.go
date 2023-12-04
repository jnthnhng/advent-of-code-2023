package main

import (
	"bufio"
	"log"
	"os"
)

/*
*
Pseudo:

	Read data from file
	For each line in the file
		Use two pointers:
			one at the beginning, iterating forward
			one at the end, iterating backwards.
		Stop at the first integer
		Combine the integers and add to the array of results
	Creates a result array

	Add up all the number and return the result
*/
func main() {
	// Open file
	f, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	// defer file closing
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		var firstInt int32
		for index, value := range line {
			if line[index] >= '0' && line[index] <= '9' {
				firstInt = value
				break
			}
		}

		var secondInt int32
		last := len(line) - 1
		for index := range line {
			value := line[last-index]
			if value >= '0' && value <= '9' {
				secondInt = int32(value)
				break
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
