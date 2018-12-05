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

	scanner := bufio.NewScanner(file)
	var lines []string

	// Populate list
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	s1, s2 := findHamming1(lines)

	// Remove common letter
	res := ""
	for i:=0; i<len(s1); i++ {
		if s1[i] != s2[i] {
          res = s1[:i] + s2[i+1:]
          break
		}
	}
	println(res)
}

func findHamming1(strings []string) (string1 string, string2 string){
    for pivotStringI := 0; pivotStringI<len(strings); pivotStringI++ {
       for checkStringI := 1; checkStringI<len(strings); checkStringI++ {
			hamming := 0
			for i, _ := range strings[pivotStringI] { // iterate through len of string
			  char1 := strings[pivotStringI][i]
			  char2 := strings[checkStringI][i]

              if char1 != char2 {
              	hamming += 1
			  }
			  if hamming > 1 {
			  	break
			  }
			}
			if hamming == 1 {
				foundString1 := strings[pivotStringI]
				foundString2 := strings[checkStringI]
				return foundString1, foundString2
			}
		}
	}
    return "", ""
}