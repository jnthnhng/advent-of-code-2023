package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//Algorithm:

func main() {

	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)

	for s.Scan() {
		line := s.Text()
		fmt.Println(line)
	}
}
