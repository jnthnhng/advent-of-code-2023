package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode/utf8"
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

	s := bufio.NewScanner(f)
	//calibrationValuesSum := getCalibrationValuesSum(s)
	//fmt.Println("Part One: ", calibrationValuesSum)

	digitsSum := parseNumberInput(s)
	fmt.Println("Part Two:", digitsSum)

}

func getCalibrationValuesSum(s *bufio.Scanner) int {
	var firstInt int32
	var secondInt uint8
	var totalSum int
	for s.Scan() {
		line := s.Text()

		for index, value := range line {
			if line[index] >= '0' && line[index] <= '9' {
				firstInt = value
				break
			}
		}

		last := len(line) - 1
		for index := range line {
			value := line[last-index]
			if value >= '0' && value <= '9' {
				secondInt = value
				break
			}
		}
		totalSum += addValues(firstInt, int32(secondInt))
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return totalSum
}

func parseNumberInput(s *bufio.Scanner) int {

	digits := map[string]int32{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	total := 0

	for s.Scan() {
		line := s.Text()
		//fmt.Println(line)

		length := utf8.RuneCountInString(line)
		// First Value
		ptr1, ptr2 := 0, 2
		breakPoint := ptr2 + 2
		firstValueFound := false
		var firstValue int32
		//fmt.Println(ptr1, ptr2, breakPoint, length, "\n")

		if length == 1 {
			firstValueFound = true
			firstValue = int32(line[ptr1])
		}
		for ptr2 <= length+2 && firstValueFound == false {
			var slice string
			if ptr2 < length {
				slice = line[ptr1 : ptr2+1]
			} else {
				slice = line[ptr1:]
			}

			//fmt.Println(ptr1, ptr2, breakPoint, "\n", slice)

			for i := ptr1; i < ptr2; i++ {
				if line[i] >= '0' && line[i] <= '9' {
					firstValue = int32(line[i])
					firstValueFound = true
					break
				}
			}

			if firstValue == 0 {
				if value, ok := digits[slice]; ok {
					firstValue = value
					firstValueFound = true
					break
				}
			}

			if ptr2 >= breakPoint {
				ptr1++
				ptr2 = ptr1 + 2
				breakPoint = ptr2 + 2
			} else {
				ptr2++
			}
		}

		// Second Value
		last := length
		ptr1, ptr2 = last, last-3
		breakPoint = ptr2 - 2
		var secondValue int32
		secondValueFound := false
		//fmt.Println(ptr1, ptr2, breakPoint, length, "\n")
		for ptr1 >= -2 && secondValueFound == false && length > 0 {
			var slice string
			if ptr1 > 0 && ptr1 == length && ptr2 > 0 {
				slice = line[ptr2:]
			} else if ptr2 <= 0 {
				slice = line[:ptr1]
			} else {
				slice = line[ptr2 : ptr1+1]
			}

			//fmt.Println(slice)
			for i := ptr1 - 1; i > ptr2; i-- {
				if line[i] >= '0' && line[i] <= '9' {
					secondValue = int32(line[i])
					secondValueFound = true
					break
				}
			}
			//fmt.Println(ptr1, ptr2, breakPoint, slice, len(slice))
			if value, ok := digits[slice]; ok {
				secondValue = value

				secondValueFound = true
				break
			}

			if ptr2 <= breakPoint {
				ptr1--
				ptr2 = ptr1 - 2
				breakPoint = ptr2 - 2
			} else if ptr2 == breakPoint && ptr1 >= 0 {
				ptr1--
			} else {
				ptr2--
			}
		}
		total += addValues(firstValue, secondValue)
	}
	return total
}

func addValues(first int32, second int32) int {

	firstValue, _ := strconv.Atoi(string(first))
	secondValue, _ := strconv.Atoi(string(second))
	fmt.Println(firstValue, secondValue)

	totalSum := 0
	totalSum += (firstValue * 10) + secondValue
	return totalSum
}
