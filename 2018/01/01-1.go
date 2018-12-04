package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	state := 0
	for scanner.Scan() {
		readInt, _ := strconv.Atoi(scanner.Text())
		state += readInt
	}
	fmt.Println(state)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}