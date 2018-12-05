package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewScanner(file)

	two := 0
	three := 0

	for reader.Scan() {
		m := make(map[rune]int)

		for _, char := range reader.Text() {
			_, ok := m[char]

			if ok {
				m[char] += 1
			} else {
				m[char] = 1
			}
		}

		inc2 := false
		inc3 := false

		for _, cnt := range m {
			if cnt == 2 {
				inc2 = true
			}
            if cnt == 3 {
            	inc3 = true
			}
		}

		if inc2 {
			two += 1
		}
		if inc3 {
			three += 1
		}
	}

	println(two * three)
}