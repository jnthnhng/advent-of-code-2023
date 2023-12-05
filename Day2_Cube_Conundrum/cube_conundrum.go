package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//Algorithm:
//
//	split string by: game number, data
//	split data, by game sets
//	iterate through each game in the set, and

func main() {

	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)

	//gameDataMap := map[string][]string{}
	maxNumberOfCubes := map[string]string{
		"red":   "12",
		"green": "13",
		"blue":  "14",
	}

	var possibleGames []string

	for s.Scan() {
		line := s.Text()

		// Get game number
		splitGameAndData := strings.Split(line, ":")
		gameHeader := splitGameAndData[0]
		gameNumber := strings.Split(gameHeader, " ")[1]

		// Game Data
		gameData := strings.TrimSpace(splitGameAndData[1])
		gameSets := strings.Split(gameData, ";")

		isPossible := true

		for _, set := range gameSets {
			trimSpaces := strings.TrimSpace(set)
			splitByColors := strings.Split(trimSpaces, ",")

			for _, numColor := range splitByColors {
				trimSpacesFromColor := strings.TrimSpace(numColor)
				splitAmountFromColor := strings.Split(trimSpacesFromColor, " ")

				amount, _ := strconv.Atoi(splitAmountFromColor[0])
				color := splitAmountFromColor[1]
				maxCubes, _ := strconv.Atoi(maxNumberOfCubes[color])

				if amount > maxCubes {
					isPossible = false
				}

			}
		}
		if isPossible {
			possibleGames = append(possibleGames, gameNumber)
		}
	}
	totalSum := sumArrayValues(possibleGames)
	fmt.Println(totalSum)
}

func sumArrayValues(array []string) string {
	total := 0

	for _, value := range array {
		num, _ := strconv.Atoi(value)
		total += num
	}
	return strconv.Itoa(total)
}
