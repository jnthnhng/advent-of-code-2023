package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Open file
	f, err := os.Open(os.Args[1])

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
	var calibrationValuesSum int = getCalibrationValuesSum(f, scanner)
	fmt.Println(calibrationValuesSum)

}

func getCalibrationValuesSum(f *os.File, s *bufio.Scanner) int {

	totalSum := 0
	for s.Scan() {
		line := s.Text()

		var firstInt int32
		for index, value := range line {
			if line[index] >= '0' && line[index] <= '9' {
				firstInt = value
				break
			}
		}

		var secondInt uint8
		last := len(line) - 1
		for index := range line {
			value := line[last-index]
			if value >= '0' && value <= '9' {
				secondInt = value
				break
			}
		}
		firstValue, _ := strconv.Atoi(string(firstInt))
		secondValue, _ := strconv.Atoi(string(secondInt))
		totalSum += (firstValue * 10) + secondValue
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	return totalSum
}
