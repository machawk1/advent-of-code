package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	frequency := 0
	m := make(map[int]bool)
	m[0] = true
	revisitedFrequency := 0
	revisitedBool := false

	for ; !revisitedBool;  {
		revisitedBool, revisitedFrequency = iterateLinesUntilFound(os.Args[1], frequency, m)
		if !revisitedBool {
			frequency = revisitedFrequency
		}
	}
	println(revisitedFrequency)
}

func iterateLinesUntilFound(fileName string, frequency int, m map[int]bool) (found bool, revisitedFrequency int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var line string
	var readInt int
	for {
		line, err = reader.ReadString('\n')

		// End of file, repeat again until frequency match is found
		if err != nil {
			if err == io.EOF {
				return false, frequency
			}
		}

		readInt, err = strconv.Atoi(line[:len(line)-1])

		frequency += readInt

		if m[frequency] == true {
			return true, frequency
		}

		// Add an entry to indicate "we've seen this frequecy before"
		// There is probably a better way/structure to do this -- learning experience
		m[frequency] = true
	}
}